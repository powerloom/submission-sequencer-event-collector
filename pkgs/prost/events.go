package prost

import (
	"context"
	"encoding/json"
	"fmt"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs"
	"submission-sequencer-collector/pkgs/clients"
	"submission-sequencer-collector/pkgs/redis"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
)

// ProcessEvents processes all events from a given block with proper context management
func ProcessEvents(ctx context.Context, block *types.Block) error {
	if block == nil {
		return fmt.Errorf("received nil block")
	}

	hash := block.Hash()

	// Create a filter query to fetch logs for the block
	filterQuery := ethereum.FilterQuery{
		BlockHash: &hash,
		Addresses: []common.Address{common.HexToAddress(config.SettingsObj.ContractAddress)},
	}

	var logs []types.Log
	var err error

	operation := func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			logs, err = Client.FilterLogs(ctx, filterQuery)
			return err
		}
	}

	if err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewConstantBackOff(200*time.Millisecond), 3)); err != nil {
		errorMsg := fmt.Sprintf("Error fetching logs for block number %d: %s", block.Number().Int64(), err.Error())
		clients.SendFailureNotification(pkgs.ProcessEvents, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return fmt.Errorf("failed to fetch logs: %w", err)
	}

	log.Infof("Processing %d logs for block number %d", len(logs), block.Number().Int64())

	// Create a semaphore to limit concurrent event processing
	// Limit to 2 concurrent events, we are processing only EpochReleased and SnapshotBatchSubmitted events
	eventSemaphore := make(chan struct{}, 2)
	var wg sync.WaitGroup
	errChan := make(chan error, len(logs)) // Buffered channel for errors

	// Create a single event processing context with timeout
	eventCtx, eventCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer eventCancel()

	for _, vLog := range logs {
		vLog := vLog
		wg.Add(1)

		// Acquire semaphore with timeout
		select {
		case eventSemaphore <- struct{}{}:
			// Got permission to proceed
		case <-ctx.Done():
			wg.Done()
			return ctx.Err()
		case <-time.After(5 * time.Second):
			log.Warnf("⚠️ Timeout waiting for event semaphore for block %d", block.Number().Int64())
			wg.Done()
			continue
		}

		go func() {
			defer wg.Done()
			defer func() { <-eventSemaphore }() // Release semaphore when done

			// Check the event signature and handle the events
			switch vLog.Topics[0].Hex() {
			case ContractABI.Events["EpochReleased"].ID.Hex():
				if err := handleEpochReleasedEvent(eventCtx, block, vLog); err != nil {
					select {
					case errChan <- fmt.Errorf("epoch released event: %w", err):
					case <-eventCtx.Done():
					}
				}
			case ContractABI.Events["SnapshotBatchSubmitted"].ID.Hex():
				if err := handleSnapshotBatchSubmittedEvent(eventCtx, block, vLog); err != nil {
					select {
					case errChan <- fmt.Errorf("snapshot batch event: %w", err):
					case <-eventCtx.Done():
					}
				}
			}
		}()
	}

	// Wait for all events to complete or context to be cancelled
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
		close(errChan)
	}()

	select {
	case <-done:
		// Collect errors
		var errs []error
		for err := range errChan {
			errs = append(errs, err)
		}

		if len(errs) > 0 {
			return fmt.Errorf("multiple event processing errors: %v", errs)
		}
	case <-ctx.Done():
		return ctx.Err()
	case <-eventCtx.Done():
		return fmt.Errorf("event processing timeout after 30 seconds")
	}

	return nil
}

// handleEpochReleasedEvent processes epoch released events
func handleEpochReleasedEvent(ctx context.Context, block *types.Block, vLog types.Log) error {
	log.Debugf("EpochReleased event detected in block %d", block.Number().Int64())

	// Parse the `EpochReleased` event from the log
	releasedEvent, err := Instance.ParseEpochReleased(vLog)
	if err != nil {
		errorMsg := fmt.Sprintf("Epoch release parse error for block %d: %s", block.Number().Int64(), err.Error())
		clients.SendFailureNotification(pkgs.ProcessEvents, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return fmt.Errorf("failed to parse epoch released event: %w", err)
	}

	// Check if the DataMarketAddress in the event matches any address in the DataMarketAddress array
	if isValidDataMarketAddress(releasedEvent.DataMarketAddress.Hex()) {
		// Extract the epoch ID and the data market address from the event
		newEpochID := releasedEvent.EpochId
		dataMarketAddress := releasedEvent.DataMarketAddress.Hex()

		log.Debugf("Detected epoch released event at block %d for data market %s with epochID %s", block.Header().Number, dataMarketAddress, newEpochID.String())

		// Get submission window duration from contract or config
		submissionLimitTimeDuration, err := getSubmissionLimitTimeDuration(ctx, dataMarketAddress)
		if err != nil {
			log.Errorf("Failed to get submission limit time duration for data market %s, epoch ID %s, block %d: %s", dataMarketAddress, newEpochID.String(), block.Number().Int64(), err.Error())
			return fmt.Errorf("failed to get submission limit time duration: %w", err)
		}

		log.Infof("⏲️ Beginning submission window for epochID %s, data market %s, duration: %f seconds", newEpochID.String(), dataMarketAddress, submissionLimitTimeDuration.Seconds())

		// Create a background context for StartSubmissionWindow since it manages its own lifecycle
		if err := windowManager.StartSubmissionWindow(context.Background(), dataMarketAddress, newEpochID, submissionLimitTimeDuration, block.Number().Int64()); err != nil {
			log.Errorf("Failed to start submission window: %v", err)
			// Don't return error, continue processing other events
		}

		// Send updateRewards to relayer when current epoch is a multiple of epoch interval (config param)
		if newEpochID.Int64()%config.SettingsObj.RewardsUpdateEpochInterval == 0 {
			// Create a new context with timeout for reward updates
			rewardCtx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
			defer cancel()

			// Send periodic updateRewards to relayer throughout the day
			if err := sendRewardUpdates(rewardCtx, dataMarketAddress, newEpochID.String()); err != nil {
				errMsg := fmt.Sprintf("Failed to send reward updates for epoch %s within data market %s: %v", newEpochID.String(), dataMarketAddress, err)
				clients.SendFailureNotification(pkgs.SendUpdateRewardsToRelayer, errMsg, time.Now().String(), "High")
				log.Error(errMsg)
				return fmt.Errorf("failed to send reward updates: %w", err)
			}
		}
	}

	return nil
}

// handleSnapshotBatchSubmittedEvent processes snapshot batch submitted events
func handleSnapshotBatchSubmittedEvent(ctx context.Context, block *types.Block, vLog types.Log) error {
	log.Debugf("SnapshotBatchSubmitted event detected in block %d", block.Number().Int64())

	// Parse the `SnapshotBatchSubmitted` event from the log
	releasedEvent, err := Instance.ParseSnapshotBatchSubmitted(vLog)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to parse SnapshotBatchSubmitted event for block %d: %s", block.Number().Int64(), err.Error())
		clients.SendFailureNotification(pkgs.ProcessEvents, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return fmt.Errorf("failed to parse snapshot batch submitted event: %w", err)
	}

	// Check if the DataMarketAddress in the event matches any address in the DataMarketAddress array
	if isValidDataMarketAddress(releasedEvent.DataMarketAddress.Hex()) {
		// Extract the epoch ID and the data market address from the event
		epochID := releasedEvent.EpochId
		dataMarketAddress := releasedEvent.DataMarketAddress.Hex()

		// Create an instance of batch details
		batch := BatchDetails{
			EpochID:           epochID,
			DataMarketAddress: dataMarketAddress,
			BatchCID:          releasedEvent.BatchCid,
		}

		// Serialize the struct to JSON
		jsonData, err := json.Marshal(batch)
		if err != nil {
			errMsg := fmt.Sprintf("Serialization failed for batch details of epoch %s, data market %s: %v", epochID.String(), dataMarketAddress, err)
			clients.SendFailureNotification(pkgs.ProcessEvents, err.Error(), time.Now().String(), "High")
			log.Error(errMsg)
			return fmt.Errorf("failed to marshal batch details: %w", err)
		}

		if config.SettingsObj.AttestorQueuePushEnabled {
			// Push the serialized data to Redis
			if err = redis.LPush(ctx, "attestorQueue", jsonData).Err(); err != nil {
				errMsg := fmt.Sprintf("Error pushing batch details to attestor queue in Redis for epoch %s, data market %s: %v",
					epochID.String(), dataMarketAddress, err)
				clients.SendFailureNotification(pkgs.ProcessEvents, errMsg, time.Now().String(), "High")
				log.Error(errMsg)
				return fmt.Errorf("failed to push batch details to attestor queue: %w", err)
			}
		}

		log.Infof("✅ Batch details successfully pushed to Redis for epoch %s in data market %s", epochID.String(), dataMarketAddress)
	}

	return nil
}
