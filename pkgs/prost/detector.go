package prost

import (
	"context"
	"math/big"
	"strconv"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs"
	"submission-sequencer-collector/pkgs/redis"
	"time"

	log "github.com/sirupsen/logrus"
)

// StartFetchingBlocks continuously fetches blocks and processes events
func StartFetchingBlocks() {
	log.Println("Submission Event Collector started")

	// Fetch the last processed block number from Redis
	lastProcessedBlockNum, err := redis.Get(context.Background(), pkgs.CurrentBlockNumber)
	if err != nil {
		log.Warnln("Failed to fetch the last processed block number, starting from 0: ", err.Error())
		lastProcessedBlockNum = "0" // Default to 0 if unable to retrieve from Redis
	}

	// Parse the last processed block number
	currentBlockNum, err := strconv.ParseInt(lastProcessedBlockNum, 10, 64)
	if err != nil {
		log.Warnln("Could not parse last processed block number, starting from 0: ", err.Error())
		currentBlockNum = 0 // Default to 0 if parsing fails
	}

	for {
		// Fetch the latest block available on the chain
		latestBlock, err := Client.BlockByNumber(context.Background(), nil)
		if err != nil {
			log.Errorf("Error fetching latest block: %s", err)
			time.Sleep(100 * time.Millisecond) // Sleep briefly before retrying to prevent spamming
			continue
		}

		latestBlockNum := latestBlock.Number().Int64()

		// Process all blocks from the currentBlockNum to the latest block
		for blockNum := currentBlockNum + 1; blockNum <= latestBlockNum; blockNum++ {
			block, err := Client.BlockByNumber(context.Background(), big.NewInt(blockNum))
			if err != nil {
				log.Errorf("Failed to fetch block %d: %s", blockNum, err)
				break // Break the inner loop to retry fetching this block on the next cycle
			}

			if block == nil {
				log.Errorln("Received nil block for number: ", blockNum)
				break
			}

			log.Debugf("Processing block: %d", blockNum)

			// Process events in the block.
			ProcessEvents(block)

			// Update and store the current block number in redis
			currentBlockNum = blockNum
			if err := redis.Set(context.Background(), pkgs.CurrentBlockNumber, strconv.FormatInt(currentBlockNum, 10), 0); err != nil {
				log.Errorf("Failed to update current block number in Redis: %s", err)
			}
		}

		// Sleep for the configured block time before rechecking
		time.Sleep(time.Duration(config.SettingsObj.BlockTime) * time.Millisecond)
	}
}
