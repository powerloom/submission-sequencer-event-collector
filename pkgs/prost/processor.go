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
		switch vLog.Topics[0].Hex() {
		case ContractABI.Events["EpochReleased"].ID.Hex():
			// Parse the `EpochReleased` event
			releasedEvent, err := Instance.ParseEpochReleased(vLog)
			if err != nil {
				clients.SendFailureNotification("Epoch release parse error: ", err.Error(), time.Now().String(), "High")
				log.Errorln("Error unpacking epoch released event: ", err)
				continue
			}

			// Check if the DataMarketAddress matches
			if releasedEvent.DataMarketAddress.Hex() == config.SettingsObj.DataMarketAddress {
				log.Debugf("Epoch Released at block %d: %s\n", block.Header().Number, releasedEvent.EpochId.String())

				// Fetch the current epoch from the Redis
				currentEpochID, err := redis.Get(context.Background(), pkgs.CurrentEpoch)
				if err != nil {
					log.Debugln("Failed to fetch the current epoch from Redis: ", err.Error())
					continue
				}

				// Update Redis if a new epoch has been released
				if currentEpochID < releasedEvent.EpochId.String() {
					newEpochID := releasedEvent.EpochId.String()
					if err := redis.Set(context.Background(), pkgs.CurrentEpoch, newEpochID, 0); err != nil {
						log.Errorf("Failed to update current epoch in Redis: %s", err)
						continue
					}

					// Calculate submission limit based on the current block number
					submissionLimit := UpdateSubmissionLimit(new(big.Int).Set(block.Number()))

					// Wait for blocks until (block number + submission limit)
					go waitForBlocks(releasedEvent.EpochId, submissionLimit, block)
				}
			}
		}
	}
}

func waitForBlocks(epochID, submissionLimit *big.Int, startBlock *types.Block) {
	// Calculate the end block number
	endBlockNum := new(big.Int).Add(startBlock.Number(), submissionLimit)

	// Initialize with the starting block header
	headers := []string{startBlock.Header().Hash().Hex()}

	log.Infof("Waiting for blocks from %d to %d for epoch %s", startBlock.Number().Uint64(), endBlockNum.Uint64(), epochID.String())

	// Wait until the current block number reaches the end block number
	for {
		currentBlock, err := Client.BlockByNumber(context.Background(), nil) // Fetch the latest block to get the current block number
		if err != nil {
			log.Errorf("Error fetching latest block: %s", err)
			time.Sleep(100 * time.Millisecond) // Sleep briefly before retrying to prevent spamming
			continue
		}

		if currentBlock.Number().Cmp(endBlockNum) >= 0 {
			break // Exit loop when we reach or exceed the end block number
		}

		// Add current block header to the list
		headers = append(headers, currentBlock.Header().Hash().Hex())

		// Sleep briefly before checking again.
		time.Sleep(time.Duration(config.SettingsObj.BlockTime*500) * time.Millisecond)
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
