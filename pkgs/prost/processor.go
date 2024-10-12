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
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
)

type EpochDetails struct {
	StartBlockNumber int64
	EndBlockNumber   int64
}

type SubmissionDetails struct {
	EpochID    *big.Int
	ProjectMap map[string][]string // ProjectID -> SubmissionKeys
}

func ProcessEvents(block *types.Block) {
	var logs []types.Log
	var err error

	hash := block.Hash()
	filterQuery := ethereum.FilterQuery{
		BlockHash: &hash,
		Addresses: []common.Address{common.HexToAddress(config.SettingsObj.ContractAddress)},
	}

	operation := func() error {
		logs, err = Client.FilterLogs(context.Background(), filterQuery)
		return err
	}

	if err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewConstantBackOff(200*time.Millisecond), 3)); err != nil {
		log.Errorln("Error fetching logs: ", err.Error())
		clients.SendFailureNotification(pkgs.ProcessEvents, fmt.Sprintf("Error fetching logs: %s", err.Error()), time.Now().String(), "High")
		return
	}

	// Process the logs for the current block
	for _, vLog := range logs {
		// Check the event signature and handle the `EpochReleased` event
		switch vLog.Topics[0].Hex() {
		case ContractABI.Events["EpochReleased"].ID.Hex():
			// Parse the `EpochReleased` event from the log
			releasedEvent, err := Instance.ParseEpochReleased(vLog)
			if err != nil {
				clients.SendFailureNotification("Epoch release parse error: ", err.Error(), time.Now().String(), "High")
				log.Errorln("Error unpacking epoch released event: ", err)
				continue
			}

			// Ensure the DataMarketAddress in the event matches the configured DataMarketAddress
			if releasedEvent.DataMarketAddress.Hex() == config.SettingsObj.DataMarketAddress {
				log.Debugf("Epoch Released at block %d: %s\n", block.Header().Number, releasedEvent.EpochId.String())

				// Fetch the current epoch ID from Redis
				currentEpochID, err := redis.Get(context.Background(), pkgs.CurrentEpoch)
				if err != nil {
					log.Debugln("Failed to fetch the current epoch from Redis: ", err.Error())
					continue
				}

				// If the new epoch is greater than the current epoch, update Redis
				if currentEpochID < releasedEvent.EpochId.String() {
					newEpochID := releasedEvent.EpochId

					// Update Redis with the new current epoch ID
					if err := redis.Set(context.Background(), pkgs.CurrentEpoch, newEpochID.String(), 0); err != nil {
						log.Errorf("Failed to update current epoch in Redis: %s", err)
						continue
					}

					// Calculate submission limit based on the current block number
					submissionLimit := UpdateSubmissionLimit(new(big.Int).Set(block.Number()))

					// Determine the end block number by adding the submission limit to the current block number
					endBlockNum := new(big.Int).Add(block.Number(), submissionLimit)

					// Cache the submission limit marker for the new epoch
					epochEndBlockMarkers[newEpochID] = EpochDetails{
						StartBlockNumber: block.Number().Int64(),
						EndBlockNumber:   endBlockNum.Int64(),
					}
				}
			}
		}
	}
}

func checkAndTriggerBatchPreparation(currentBlock *types.Block) {
	// Get the current block number
	currentBlockNum := currentBlock.Number().Int64()

	// Iterate through the epoch-end block markers
	for epochID, epochDetails := range epochEndBlockMarkers {
		// Check if the current block number has reached the end block number for an epoch
		if currentBlockNum == epochDetails.EndBlockNumber {
			log.Infof("Triggering batch preparation for epoch %s", epochID.String())

			// Trigger the batch preparation logic
			go triggerBatchPreparation(epochID, epochDetails.StartBlockNumber, currentBlockNum)

			// Remove the marker for this epoch once processed
			delete(epochEndBlockMarkers, epochID)
		}
	}
}

func triggerBatchPreparation(epochID *big.Int, startBlockNum, endBlockNum int64) {
	// Initialize a slice for storing the headers
	headers := make([]string, 0)

	// Iterate over the block numbers and fetch the headers
	for blockNum := startBlockNum; blockNum <= endBlockNum; blockNum++ {
		// Fetch the block hash
		blockHash := blockNumberToHash[blockNum]

		// Add the block header to the list
		headers = append(headers, blockHash.Hex())
	}

	log.Debugf("Collected headers for epoch %s: %v", epochID.String(), headers)

	// Fetch valid submission keys for the epoch
	submissionkeys, err := getValidSubmissionKeys(epochID.Uint64(), headers)
	if err != nil {
		log.Errorln("Failed to fetch submission keys: ", submissionkeys)
	}

	// Construct the project map [ProjectID -> SubmissionKeys]
	projectMap := constructProjectMap(submissionkeys)

	// Create an instance of submission details
	submissionDetails := SubmissionDetails{
		EpochID:    epochID,
		ProjectMap: projectMap,
	}

	// Serialize the struct to JSON
	jsonData, err := json.Marshal(submissionDetails)
	if err != nil {
		log.Fatalf("Error serializing submission details: %s", err)
	}

	// Push the serialized data to Redis
	err = redis.RedisClient.LPush(context.Background(), "batchQueue", jsonData).Err()
	if err != nil {
		log.Fatalf("Error pushing data to Redis: %s", err)
	}
}

func getValidSubmissionKeys(epochID uint64, headers []string) ([]string, error) {
	// Initialize an empty slice to store valid submission keys
	submissionKeys := make([]string, 0)

	// Iterate through the list of headers
	for _, header := range headers {
		keys := redis.RedisClient.SMembers(context.Background(), redis.SubmissionSetByHeaderKey(epochID, header)).Val()
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
		if len(parts) != 3 {
			log.Errorln("Improper key stored in redis: ", submissionKey)
			clients.SendFailureNotification(pkgs.ConstructProjectMap, fmt.Sprintf("Improper key stored in redis: %s", submissionKey), time.Now().String(), "High")
			continue // skip malformed entries
		}

		projectID := parts[1]
		projectMap[projectID] = append(projectMap[projectID], submissionKey)
	}

	return projectMap
}
