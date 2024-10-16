package prost

import (
	"context"
	"fmt"
	"math/big"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs"
	"submission-sequencer-collector/pkgs/clients"
	"submission-sequencer-collector/pkgs/redis"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
)

// StartFetchingBlocks continuously fetches blocks and processes events
func StartFetchingBlocks() {
	log.Println("Submission Event Collector started")

	for {
		// Fetch the latest block available on the chain
		latestBlock, err := Client.BlockByNumber(context.Background(), nil)
		if err != nil {
			log.Errorf("Error fetching latest block: %s", err)
			time.Sleep(100 * time.Millisecond) // Sleep briefly before retrying to prevent spamming
			continue
		}

		// Check if there's a gap between the current block and the latest block on the chain.
		for blockNum := CurrentBlock.Number().Int64() + 1; blockNum <= latestBlock.Number().Int64(); blockNum++ {
			// Retrieve the block from the client using the specified block number
			block, err := fetchBlock(blockNum)
			if err != nil {
				clients.SendFailureNotification(pkgs.StartFetchingBlocks, fmt.Sprintf("Error fetching block %d: %s", blockNum, err.Error()), time.Now().String(), "High")
				log.Errorf("Error fetching block %d: %s", blockNum, err.Error())
				continue
			}

			if block == nil {
				log.Errorln("Received nil block for number: ", blockNum)
				break
			}

			log.Debugf("Processing block: %d", blockNum)

			// Check and trigger batch preparation if submission limit is reached for any epoch
			go checkAndTriggerBatchPreparation(block)

			// Process the events in the block
			go ProcessEvents(block)

			// Add block number and its hash to Redis
			if err = redis.Set(context.Background(), redis.BlockHashByNumber(blockNum), block.Hash().Hex(), 0); err != nil {
				log.Errorf("Failed to set block hash for block number %d in Redis: %s", blockNum, err)
			}

			// Update current block and store it in Redis
			updateCurrentBlock(context.Background(), block)
		}

		// Sleep for the configured block time before rechecking
		time.Sleep(time.Duration(config.SettingsObj.BlockTime) * time.Millisecond)
	}
}

// fetchBlock retrieves the block from the client using the specified block number
func fetchBlock(blockNum int64) (*types.Block, error) {
	var block *types.Block

	// Define the operation to fetch the block
	operation := func() error {
		var err error
		block, err = Client.BlockByNumber(context.Background(), big.NewInt(blockNum))
		if err != nil {
			log.Errorf("Failed to fetch block %d: %s", blockNum, err)
			return err // Return the error to trigger a retry
		}

		return nil // Block successfully fetched, return nil to stop retries
	}

	// Retry fetching the block with backoff strategy
	if err := backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewConstantBackOff(200*time.Millisecond), 3)); err != nil {
		log.Errorln("Error fetching block after retries: ", err.Error())
		return nil, err
	}

	return block, nil // Return the successfully fetched block
}

// updateCurrentBlock updates the global current block variable and stores it in Redis
func updateCurrentBlock(ctx context.Context, block *types.Block) {
	CurrentBlock = block
	if err := redis.Set(ctx, pkgs.CurrentBlockNumberKey, CurrentBlock.Number().String(), 0); err != nil {
		log.Errorf("Failed to update current block number %d in Redis: %s", CurrentBlock.Number().Int64(), err)
	}
}
