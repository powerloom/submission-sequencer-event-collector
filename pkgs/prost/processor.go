package prost

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
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

type EpochMarkerDetails struct {
	EpochReleaseBlockNumber    int64
	SubmissionLimitBlockNumber int64
}

type SubmissionDetails struct {
	DataMarketAddress string
	EpochID           *big.Int
	BatchID           int
	Batch             map[string][]string // ProjectID -> SubmissionKeys
}

// ProcessEvents processes logs for the given block and handles the EpochReleased event
func ProcessEvents(block *types.Block) {
	var logs []types.Log
	var err error

	hash := block.Hash()

	// Create a filter query to fetch logs for the block
	filterQuery := ethereum.FilterQuery{
		BlockHash: &hash,
		Addresses: []common.Address{common.HexToAddress(config.SettingsObj.ContractAddress)},
	}

	operation := func() error {
		logs, err = Client.FilterLogs(context.Background(), filterQuery)
		return err
	}

	if err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewConstantBackOff(200*time.Millisecond), 3)); err != nil {
		errorMsg := fmt.Sprintf("Error fetching logs for block number %d: %s", block.Number().Int64(), err.Error())
		log.Error(errorMsg)
		clients.SendFailureNotification(pkgs.ProcessEvents, errorMsg, time.Now().String(), "High")
		return
	}

	log.Infof("Processing %d logs for block number %d", len(logs), block.Number().Int64())

	// Process the logs for the current block
	for _, vLog := range logs {
		// Check the event signature and handle the `EpochReleased` event
		switch vLog.Topics[0].Hex() {
		case ContractABI.Events["EpochReleased"].ID.Hex():
			log.Debugf("EpochReleased event detected in block %d", block.Number().Int64())

			// Parse the `EpochReleased` event from the log
			releasedEvent, err := Instance.ParseEpochReleased(vLog)
			if err != nil {
				errorMsg := fmt.Sprintf("Epoch release parse error for block %d: %s", block.Number().Int64(), err.Error())
				clients.SendFailureNotification(pkgs.ProcessEvents, errorMsg, time.Now().String(), "High")
				log.Error(errorMsg)
				continue
			}

			// Check if the DataMarketAddress in the event matches any address in the DataMarketAddress array
			if isValidDataMarketAddress(releasedEvent.DataMarketAddress.Hex()) {
				// Extract the epoch ID and the data market address from the event
				newEpochID := releasedEvent.EpochId
				dataMarketAddress := releasedEvent.DataMarketAddress.Hex()

				log.Debugf("Detected epoch released event at block %d for data market %s with epochID %s", block.Header().Number, dataMarketAddress, newEpochID.String())

				// Calculate the submission limit block based on the epoch release block number (current block number)
				submissionLimitBlockNumber, err := calculateSubmissionLimitBlock(dataMarketAddress, new(big.Int).Set(block.Number()))
				if err != nil {
					log.Errorf("Failed to calculate submission limit block number for data market %s, epoch ID %s, block %d: %s", dataMarketAddress, newEpochID.String(), block.Number().Int64(), err.Error())
					continue
				}

				log.Infof("Calculated Submission Limit Block Number for epochID %s, data market %s: %d", newEpochID.String(), dataMarketAddress, submissionLimitBlockNumber.Int64())

				// Prepare the epoch marker details
				epochMarkerDetails := EpochMarkerDetails{
					EpochReleaseBlockNumber:    block.Number().Int64(),
					SubmissionLimitBlockNumber: submissionLimitBlockNumber.Int64(),
				}

				epochMarkerDetailsJSON, err := json.Marshal(epochMarkerDetails)
				if err != nil {
					errorMsg := fmt.Sprintf("Failed to marshal epoch marker details for data market %s, epochID %s: %s", dataMarketAddress, newEpochID.String(), err.Error())
					clients.SendFailureNotification(pkgs.ProcessEvents, errorMsg, time.Now().String(), "High")
					log.Error(errorMsg)
					continue
				}

				// Store the details associated with the new epoch in Redis
				if err := redis.StoreEpochDetails(context.Background(), dataMarketAddress, newEpochID.String(), string(epochMarkerDetailsJSON)); err != nil {
					errorMsg := fmt.Sprintf("Failed to store epoch marker details for epochID %s, data market %s in Redis: %s", newEpochID.String(), dataMarketAddress, err.Error())
					clients.SendFailureNotification(pkgs.ProcessEvents, errorMsg, time.Now().String(), "High")
					log.Error(errorMsg)
				}

				log.Infof("✅ Successfully stored epoch marker details for epochID %s, data market %s in Redis", newEpochID.String(), dataMarketAddress)
			}
		}
	}
}

func checkAndTriggerBatchPreparation(currentBlock *types.Block) {
	// Get the current block number
	currentBlockNum := currentBlock.Number().Int64()
	log.Infof("🔍 Starting batch preparation check for block number: %d", currentBlockNum)

	var wg sync.WaitGroup

	// Fetch and process epoch markers set concurrently for each data market address
	for _, dataMarketAddress := range config.SettingsObj.DataMarketAddresses {
		wg.Add(1)

		go func(dataMarketAddress string) {
			defer wg.Done()

			log.Infof("Processing started for data market %s at block number: %d", dataMarketAddress, currentBlockNum)

			// Fetch all the epoch marker keys from Redis for this data market address
			epochMarkerKeys, err := redis.RedisClient.SMembers(context.Background(), redis.EpochMarkerSet(dataMarketAddress)).Result()
			if err != nil {
				log.Errorf("Failed to fetch epoch markers from Redis for data market %s: %s", dataMarketAddress, err)
				return
			}

			log.Infof("Fetched %d epoch marker keys for data market %s: %v", len(epochMarkerKeys), dataMarketAddress, epochMarkerKeys)

			// Process each epoch marker key for this data market address
			for _, epochMarkerKey := range epochMarkerKeys {
				// Retrieve the epoch marker details from Redis
				epochMarkerDetailsJSON, err := redis.RedisClient.Get(context.Background(), redis.EpochMarkerDetails(dataMarketAddress, epochMarkerKey)).Result()
				if err != nil {
					log.Errorf("Failed to fetch epoch marker details from Redis for key %s: %s", epochMarkerKey, err)
					continue
				}

				var epochMarkerDetails EpochMarkerDetails
				if err := json.Unmarshal([]byte(epochMarkerDetailsJSON), &epochMarkerDetails); err != nil {
					log.Errorf("Failed to unmarshal epoch marker details for key %s: %s", epochMarkerKey, err)
					continue
				}

				log.Debugf("📊 Epoch marker details for epoch marker key %s: %+v", epochMarkerKey, epochMarkerDetails)

				// Check if the current block number matches the submission limit block number for this epoch
				if currentBlockNum == epochMarkerDetails.SubmissionLimitBlockNumber {
					log.Infof("🔄 Initiating batch preparation for epoch %s, data market %s at submission limit block number: %d", epochMarkerKey, dataMarketAddress, currentBlockNum)

					// Convert the epoch ID string to big.Int for further processing
					epochID, ok := big.NewInt(0).SetString(epochMarkerKey, 10)
					if !ok {
						log.Errorf("Failed to convert epochID %s to big.Int for data market %s", epochMarkerKey, dataMarketAddress)
						continue
					}

					// Trigger batch preparation logic for the current epoch
					go triggerBatchPreparation(dataMarketAddress, epochID, epochMarkerDetails.EpochReleaseBlockNumber, currentBlockNum)
				}
			}
			log.Infof("Completed processing for data market %s at block number: %d", dataMarketAddress, currentBlockNum)
		}(dataMarketAddress) // Pass data market address to avoid closure issues
	}

	// Wait for all data market goroutines to finish
	wg.Wait()
}

func triggerBatchPreparation(dataMarketAddress string, epochID *big.Int, startBlockNum, endBlockNum int64) {
	// Initialize a slice to store block headers (block hashes)
	headers := make([]string, 0)

	// Calculate the total number of blocks in the range
	blockCount := endBlockNum - startBlockNum + 1
	log.Infof("🚀 Starting batch preparation for epoch %s, data market %s, processing %d blocks from %d to %d", epochID.String(), dataMarketAddress, blockCount, startBlockNum, endBlockNum)

	// Iterate through the block numbers and fetch the block headers (hashes)
	for blockNum := startBlockNum; blockNum <= endBlockNum; blockNum++ {
		// Generate the Redis key for the current block number
		blockKey := redis.BlockHashByNumber(blockNum)

		// Fetch the block hash from Redis using the generated key
		blockHashValue, err := redis.Get(context.Background(), blockKey)
		if err != nil {
			log.Errorf("Failed to fetch block hash for block %d: %s", blockNum, err.Error())
			continue // Skip this block and move to the next
		}

		// Convert the block hash from string to common.Hash type
		blockHash := common.HexToHash(blockHashValue)

		// Add the block hash to the headers slice
		headers = append(headers, blockHash.Hex())
	}

	log.Infof("📦 Collected %d headers for epoch %s in data market %s", len(headers), epochID.String(), dataMarketAddress)

	// Fetch valid submission keys for the epoch
	submissionKeys, err := getValidSubmissionKeys(context.Background(), epochID.Uint64(), headers, dataMarketAddress)
	if err != nil {
		log.Errorf("Failed to fetch submission keys for epoch %s in data market %s: %s", epochID.String(), dataMarketAddress, err.Error())
	}

	log.Infof("🔑 Retrieved %d valid submission keys for epoch %s in data market %s", len(submissionKeys), epochID.String(), dataMarketAddress)

	// Update total submission count for the specified data market address
	if err := UpdateSlotSubmissionCount(context.Background(), epochID, dataMarketAddress, submissionKeys); err != nil {
		log.Errorf("Failed to update slot submission counts for epoch %s in data market %s: %s", epochID, dataMarketAddress, err.Error())
	}

	// Construct the project map [ProjectID -> SubmissionKeys]
	projectMap := constructProjectMap(submissionKeys)
	log.Infof("📊 Project map created with %d projects for epoch %s in data market %s", len(projectMap), epochID.String(), dataMarketAddress)

	// Arrange the projectMap into batches of submission keys
	batches := arrangeSubmissionKeysInBatches(projectMap)
	log.Infof("🔄 Arranged %d batches of submission keys for epoch %s in data market %s", len(batches), epochID.String(), dataMarketAddress)

	// Send the size of the batches to the external tx relayer service
	if err = sendBatchSizeToRelayer(dataMarketAddress, epochID, len(batches)); err != nil {
		log.Errorf("Failed to send submission batch size for epoch %s in data market %s: %v", epochID, dataMarketAddress, err)
	}

	log.Infof("📨 Batch size %d sent successfully for epoch %s in data market %s", len(batches), epochID.String(), dataMarketAddress)

	// Iterate over all the batches, construct submission details for each batch,
	// serialize them to JSON, and push the serialized data to Redis for further processing
	for i, batch := range batches {
		// Create an instance of submission details
		submissionDetails := SubmissionDetails{
			EpochID:           epochID,
			BatchID:           i + 1,
			Batch:             batch,
			DataMarketAddress: dataMarketAddress,
		}

		// Serialize the struct to JSON
		jsonData, err := json.Marshal(submissionDetails)
		if err != nil {
			log.Fatalf("Serialization failed for submission details of batch %d, epoch %s in data market %s: %v", i+1, epochID.String(), dataMarketAddress, err)
		}

		// Push the serialized data to Redis
		err = redis.RedisClient.LPush(context.Background(), "finalizerQueue", jsonData).Err()
		if err != nil {
			log.Fatalf("Error pushing submission details of batch %d to Redis for epoch %s in data market %s: %v", i+1, epochID.String(), dataMarketAddress, err)
		}

		// Serialize the batch details to JSON
		batchJsonData, err := json.Marshal(batch)
		if err != nil {
			log.Fatalf("Serialization failed for batch details of batch %d, epoch %s in data market %s: %v", i+1, epochID.String(), dataMarketAddress, err)
		}

		// Convert the batch ID to a big integer
		batchID := big.NewInt(int64(i + 1))

		// Store the batch details with a key generated from dataMarketAddress, epochID, and batchID
		if err := redis.StoreBatchDetails(context.Background(), dataMarketAddress, epochID.String(), batchID.String(), string(batchJsonData)); err != nil {
			log.Errorf("Failed to store details for batch %d of epoch %s in data market %s: %v", batchID.Int64(), epochID.String(), dataMarketAddress, err)
			continue
		}

		log.Infof("✅ Batch %d successfully pushed to Redis and stored for epoch %s in data market %s", batchID.Int64(), epochID.String(), dataMarketAddress)
	}

	// Remove the epochID and its details from Redis after processing all batches
	if err := redis.RemoveEpochFromRedis(context.Background(), dataMarketAddress, epochID.String()); err != nil {
		log.Errorf("Error removing epoch %s from Redis for data market %s: %v", epochID.String(), dataMarketAddress, err)
	}

	log.Infof("🧹 Successfully removed epoch %s data from Redis for data market %s", epochID.String(), dataMarketAddress)
}

func getValidSubmissionKeys(ctx context.Context, epochID uint64, headers []string, dataMarketAddress string) ([]string, error) {
	// Initialize an empty slice to store valid submission keys
	submissionKeys := make([]string, 0)

	// Iterate through the list of headers
	for _, header := range headers {
		keys := redis.RedisClient.SMembers(ctx, redis.SubmissionSetByHeaderKey(dataMarketAddress, epochID, header)).Val()
		if len(keys) > 0 {
			submissionKeys = append(submissionKeys, keys...)
		}
	}

	return submissionKeys, nil
}

func constructProjectMap(submissionKeys []string) map[string][]string {
	// Initialize an empty map to store the projectID and the submission keys
	projectMap := make(map[string][]string)

	for _, submissionKey := range submissionKeys {
		parts := strings.Split(submissionKey, ".")
		if len(parts) != 4 {
			log.Errorln("Improper submission key stored in Redis: ", submissionKey)
			clients.SendFailureNotification(pkgs.ConstructProjectMap, fmt.Sprintf("Improper submission key stored in Redis: %s", submissionKey), time.Now().String(), "High")
			continue // skip malformed entries
		}

		projectID := parts[2]
		projectMap[projectID] = append(projectMap[projectID], submissionKey)
	}

	return projectMap
}

func arrangeSubmissionKeysInBatches(projectMap map[string][]string) []map[string][]string {
	batchSize := config.SettingsObj.BatchSize // Target number of project IDs per batch
	batches := make([]map[string][]string, 0) // Initialize a slice for storing batches
	currentBatch := make(map[string][]string) // Current batch being filled
	projectCount := 0                         // Track the number of project IDs in the current batch

	for projectID, submissionKeys := range projectMap {
		// Add the project to the current batch
		currentBatch[projectID] = submissionKeys
		projectCount++

		// If we've reached the batch size, finalize the current batch
		if projectCount == batchSize {
			// Add the current batch to the list of batches and reset for a new batch
			batches = append(batches, currentBatch)
			currentBatch = make(map[string][]string) // Start a new batch
			projectCount = 0                         // Reset count for the new batch
		}
	}

	// If there are leftover projects that didn't fill a complete batch, add them as well
	if projectCount > 0 {
		batches = append(batches, currentBatch)
	}

	return batches
}

// Calculate and update total submission count for a data market address
func UpdateSlotSubmissionCount(ctx context.Context, epochID *big.Int, dataMarketAddress string, submissionKeys []string) error {
	// Fetch the current day
	currentDay, err := FetchCurrentDay(common.HexToAddress(dataMarketAddress))
	if err != nil {
		log.Errorf("Failed to fetch current day for data market %s: %v", dataMarketAddress, err)
		return err
	}

	// Send eligible nodes count and slotIDs alert every 5 epochs
	if config.SettingsObj.PeriodicEligibleCountAlerts {
		if epochID.Uint64()%5 == 0 {
			// Fetch the slotIDs whose eligible submissions are recorded for the current day
			eligibleSlotSubmissionsByDayKeys := redis.EligibleSlotSubmissionsByDayKey(dataMarketAddress, currentDay.String())
			eligibleSlotIDs := redis.GetSetKeys(context.Background(), eligibleSlotSubmissionsByDayKeys)

			// Construct the message with eligible node count and slot IDs
			alertMsg := fmt.Sprintf("🔔 Epoch %d: Eligible node count for data market %s on day %s: %d\nSlot IDs: %v", epochID, dataMarketAddress, currentDay.String(), len(eligibleSlotIDs), eligibleSlotIDs)

			// Send the alert
			clients.SendFailureNotification(pkgs.SendEligibleNodesCount, alertMsg, time.Now().String(), "High")
			log.Infof(alertMsg)
		}
	}

	// Handle day transitions and trigger updateRewards
	if err := handleDayTransition(dataMarketAddress, currentDay); err != nil {
		log.Errorf("Error handling day transition for data market %s: %v", dataMarketAddress, err)
		return err
	}

	// Fetch day size for the specified data market address from Redis
	daySize, err := redis.GetDaySize(ctx, dataMarketAddress)
	if err != nil {
		log.Errorf("Failed to fetch day size for data market %s: %v", dataMarketAddress, err)
		return err
	}

	// Calculate expiration time
	expirationTime := getExpirationTime(epochID.Int64(), daySize.Int64())

	// Set the current day in Redis with the calculated expiration duration
	if err := redis.SetWithExpiration(context.Background(), redis.GetCurrentDayKey(dataMarketAddress), currentDay.String(), time.Until(expirationTime)); err != nil {
		return fmt.Errorf("failed to cache day value for data market %s in Redis: %v", dataMarketAddress, err)
	}

	// Set the last known day in Redis (for detecting day transition)
	if err := redis.Set(context.Background(), redis.LastKnownDay(dataMarketAddress), currentDay.String()); err != nil {
		return fmt.Errorf("failed to cache last known day value for data market %s in Redis: %v", dataMarketAddress, err)
	}

	// Increment the slot submissions count for a data market in Redis
	for _, submissionKey := range submissionKeys {
		parts := strings.Split(submissionKey, ".")
		slotID := parts[3]

		count, err := redis.Incr(context.Background(), redis.SlotSubmissionKey(dataMarketAddress, slotID, currentDay.String()))
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to increment submission count for slotID %s, epoch %s in data market %s: %s", slotID, epochID, dataMarketAddress, err.Error())
			clients.SendFailureNotification(pkgs.UpdateSlotSubmissionCount, errorMsg, time.Now().String(), "High")
			log.Error(errorMsg)
			return err
		}

		log.Infof("📈 Slot submission count updated successfully for slotID %s, epoch %s in data market %s: %d", slotID, epochID, dataMarketAddress, count)
	}

	return nil
}

func handleDayTransition(dataMarketAddress string, currentDay *big.Int) error {
	// Fetch the last known day value from Redis
	lastKnownDay, err := redis.Get(context.Background(), redis.LastKnownDay(dataMarketAddress))
	if err != nil {
		log.Errorf("Error fetching last known day value for data market %s from Redis: %v", dataMarketAddress, err)
		return err
	}

	// Check for day transition
	if lastKnownDay != "" && lastKnownDay != currentDay.String() {
		log.Infof("Day transition detected for data market %s: %s -> %s", dataMarketAddress, lastKnownDay, currentDay.String())

		// Fetch the eligible nodes count for the last known day (prev day)
		eligibleNodesCount, err := calculateEligibleNodes(dataMarketAddress, lastKnownDay)
		if err != nil {
			log.Errorf("Error calculating eligible nodes for data market %s on day %s: %v", dataMarketAddress, lastKnownDay, err)
			return err
		}

		log.Infof("✅ Successfully fetched eligible nodes count for data market %s on day %s: %d", dataMarketAddress, lastKnownDay, eligibleNodesCount)

		// Send the update rewards request to the external tx relayer service
		if err = sendUpdateRewardsToRelayer(dataMarketAddress, nil, nil, lastKnownDay, eligibleNodesCount); err != nil {
			errorMsg := fmt.Sprintf("🚨 Failed to send rewards update for data market %s on day %s: %v", dataMarketAddress, lastKnownDay, err)
			clients.SendFailureNotification(pkgs.SendUpdateRewardsToRelayer, errorMsg, time.Now().String(), "High")
			log.Errorf(errorMsg)
			return err
		}
	}

	return nil
}

// calculateEligibleNodes calculates the number of eligible nodes for a given data market and day
func calculateEligibleNodes(dataMarketAddress, day string) (int, error) {
	// Build the Redis key to fetch eligible slot submissions for the specified day
	eligibleSubmissionsSetKey := redis.EligibleSlotSubmissionsByDayKey(dataMarketAddress, day)

	// Retrieve the cardinality (count) of the set associated with the Redis key
	eligibleNodesCount, err := redis.GetSetCardinality(context.Background(), eligibleSubmissionsSetKey)
	if err != nil {
		log.Errorf("Error fetching eligible nodes for data market %s on day %s: %v", dataMarketAddress, day, err)
		return 0, fmt.Errorf("failed to fetch eligible nodes for data market %s on day %s: %w", dataMarketAddress, day, err)
	}

	return eligibleNodesCount, nil
}

func sendBatchSizeToRelayer(dataMarketAddress string, epochID *big.Int, batchSize int) error {
	// Define the operation that will be retried
	operation := func() error {
		// Attempt to submit the batch size
		err := clients.SendSubmissionBatchSize(dataMarketAddress, epochID, batchSize)
		if err != nil {
			log.Errorf("Error sending submission batch size for epoch %s, data market %s: %v", epochID, dataMarketAddress, err)
			return err // Return error to trigger retry
		}

		log.Infof("Successfully submitted batch size for epoch %s, data market %s", epochID, dataMarketAddress)
		return nil // Successful submission, no need for further retries
	}

	// Customize the backoff configuration
	backoffConfig := backoff.NewExponentialBackOff()
	backoffConfig.InitialInterval = 1 * time.Second // Start with a 1-second delay
	backoffConfig.Multiplier = 1.5                  // Increase interval by 1.5x after each retry
	backoffConfig.MaxInterval = 4 * time.Second     // Set max interval between retries
	backoffConfig.MaxElapsedTime = 10 * time.Second // Retry for a maximum of 10 seconds

	// Limit retries to 3 times within 10 seconds
	if err := backoff.Retry(operation, backoff.WithMaxRetries(backoffConfig, 3)); err != nil {
		log.Errorf("Failed to submit batch size for epoch %s, data market %s after multiple retries: %v", epochID, dataMarketAddress, err)
		return err
	}

	return nil
}

func sendUpdateRewardsToRelayer(dataMarketAddress string, slotIDs, submissionsList []*big.Int, day string, eligibleNodes int) error {
	// Define the operation that will be retried
	operation := func() error {
		// Attempt to send the updateRewards request
		err := clients.SendUpdateRewardsRequest(dataMarketAddress, slotIDs, submissionsList, day, eligibleNodes)
		if err != nil {
			log.Errorf("Error sending final updateRewards request for data market %s on day %s: %v. Retrying...", dataMarketAddress, day, err)
			return err // Return error to trigger retry
		}

		log.Infof("📤 Successfully sent final updateRewards request for data market %s on day %s", dataMarketAddress, day)
		return nil // Successful submission, no need for further retries
	}

	// Customize the backoff configuration
	backoffConfig := backoff.NewExponentialBackOff()
	backoffConfig.InitialInterval = 1 * time.Second // Start with a 1-second delay
	backoffConfig.Multiplier = 1.5                  // Increase interval by 1.5x after each retry
	backoffConfig.MaxInterval = 4 * time.Second     // Set max interval between retries
	backoffConfig.MaxElapsedTime = 10 * time.Second // Retry for a maximum of 10 seconds

	// Limit retries to a maximum of 3 attempts within 10 seconds
	if err := backoff.Retry(operation, backoff.WithMaxRetries(backoffConfig, 3)); err != nil {
		log.Errorf("Failed to send final updateRewards request after retries for data market %s on day %s: %v", dataMarketAddress, day, err)
		return err
	}

	return nil
}
