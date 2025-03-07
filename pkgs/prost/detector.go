package prost

import (
	"context"
	"math/big"
	"submission-sequencer-collector/pkgs/redis"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
)

// Context timeout constants
const ()

var (
	lastProcessedBlock int64
	// Resource cleanup on shutdown
	shutdownOnce sync.Once
)

// Cleanup resources
func cleanup() {
	shutdownOnce.Do(func() {
		log.Info("Cleaning up resources")
	})
}

// StartFetchingBlocks continuously fetches blocks and processes events
func StartFetchingBlocks(ctx context.Context) {
	defer cleanup()
	log.Info("Submission Event Collector started")

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	initFetchCtx, initFetchCancel := context.WithTimeout(ctx, blockFetchTimeout)
	defer initFetchCancel()
	header, err := Client.HeaderByNumber(initFetchCtx, nil)
	if err != nil {
		log.Errorf("Error fetching latest block header: %s", err)
		return
	}
	lastProcessedBlock = header.Number.Int64() - 1

	for {
		select {
		case <-ctx.Done():
			log.Info("Shutting down block fetching: context cancelled")
			return
		case <-ticker.C:
			// Get latest block number using HeaderByNumber(nil)
			fetchCtx, fetchCancel := context.WithTimeout(ctx, blockFetchTimeout)
			header, err := Client.HeaderByNumber(fetchCtx, nil)
			fetchCancel()

			if err != nil {
				if err != context.Canceled && err != context.DeadlineExceeded {
					log.Errorf("Error fetching latest block header: %s", err)
				}
				continue
			}

			currentBlockNum := header.Number.Int64()
			if currentBlockNum <= lastProcessedBlock {
				continue // Skip if we've already processed this block
			}

			// Process all blocks from last processed to current
			for blockNum := lastProcessedBlock + 1; blockNum <= currentBlockNum; blockNum++ {
				currentNum := blockNum // Capture for closure

				// Create block processing context
				blockCtx, blockCancel := context.WithTimeout(ctx, blockProcessTimeout)

				// Fetch the block
				block, err := fetchBlock(blockCtx, big.NewInt(currentNum))
				if err != nil {
					log.Errorf("Failed to fetch block %d: %s", currentNum, err)
					blockCancel()
					continue
				}

				// Process events and prepare batch concurrently
				var wg sync.WaitGroup

				// Process Events
				wg.Add(1)
				eventProcessCtx, eventProcessCancel := context.WithTimeout(ctx, eventProcessingTimeout)
				go func() {
					defer wg.Done()
					defer eventProcessCancel()
					if err := ProcessEvents(eventProcessCtx, block); err != nil {
						log.Errorf("Failed to process events for block %d: %v", currentNum, err)
					}
				}()

				// Check and trigger batch preparation
				wg.Add(1)
				batchProcessCtx, batchProcessCancel := context.WithTimeout(ctx, batchProcessingTimeout)
				go func() {
					defer wg.Done()
					defer batchProcessCancel()
					if err := checkAndTriggerBatchPreparation(batchProcessCtx, block); err != nil {
						log.Errorf("Failed to trigger batch preparation for block %d: %v", currentNum, err)
					}
				}()

				// Store block hash in Redis
				wg.Add(1)
				go func() {
					defer wg.Done()
					redisCtx, cancel := context.WithTimeout(ctx, redisOperationTimeout)
					defer cancel()
					if err := redis.SetWithExpiration(redisCtx,
						redis.BlockHashByNumber(currentNum),
						block.Hash().Hex(),
						30*time.Minute); err != nil {
						log.Errorf("Failed to store block hash for block %d: %v", currentNum, err)
					}
				}()

				// Wait for all operations to complete
				wg.Wait()
				blockCancel()

				// Update last processed block
				if currentNum > lastProcessedBlock {
					lastProcessedBlock = currentNum
				}

				if currentNum < currentBlockNum {
					log.Infof("Catching up: Processed block %d, %d more to go", currentNum, currentBlockNum-currentNum)
				}
			}
		}
	}
}

// fetchBlock retrieves a block from the client using retry logic
func fetchBlock(ctx context.Context, blockNumber *big.Int) (*types.Block, error) {
	var block *types.Block
	operation := func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			var err error
			block, err = Client.BlockByNumber(ctx, blockNumber)
			return err
		}
	}

	err := backoff.Retry(operation, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
	if err != nil {
		return nil, err
	}
	return block, nil
}
