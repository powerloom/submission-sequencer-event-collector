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
	"golang.org/x/sync/errgroup"
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

	// Create a parent error group for all block processing
	mainGroup, groupCtx := errgroup.WithContext(ctx)

	for {
		select {
		case <-ctx.Done():
			log.Info("Shutting down block fetching: context cancelled")
			mainGroup.Wait() // Wait for all processing to finish on shutdown
			return
		case <-ticker.C:
			// Get latest block number using HeaderByNumber(nil)
			fetchCtx, fetchCancel := context.WithTimeout(groupCtx, blockFetchTimeout)
			header, err := Client.HeaderByNumber(fetchCtx, nil)
			fetchCancel()

			if err != nil {
				if err != context.Canceled && err != context.DeadlineExceeded {
					log.Errorf("Error fetching latest block header: %s", err)
				}
				continue
			}

			currentBlockNum := header.Number.Int64()
			lastProcessedBlock = currentBlockNum - 1
			if currentBlockNum <= lastProcessedBlock {
				continue // Skip if we've already processed this block
			}

			// Process all blocks from last processed to current
			for blockNum := lastProcessedBlock + 1; blockNum <= currentBlockNum; blockNum++ {
				currentNum := blockNum // Capture for closure

				// Start a new goroutine for each block's processing
				mainGroup.Go(func() error {
					// Create block processing context as child of group context
					blockCtx, blockCancel := context.WithTimeout(groupCtx, blockProcessTimeout)
					defer blockCancel()

					// Fetch the block
					block, err := fetchBlock(blockCtx, big.NewInt(currentNum))
					if err != nil {
						log.Errorf("Failed to fetch block %d: %s", currentNum, err)
						return nil // Don't fail the group for fetch errors
					}

					// Create error group for this block's concurrent operations
					blockGroup, blockGroupCtx := errgroup.WithContext(blockCtx)

					// Process events concurrently
					blockGroup.Go(func() error {
						if err := ProcessEvents(blockGroupCtx, block); err != nil {
							if err != context.Canceled && err != context.DeadlineExceeded {
								log.Errorf("Error processing events for block %d: %s", currentNum, err)
							}
							return err
						}
						return nil
					})

					// Prepare batch concurrently
					blockGroup.Go(func() error {
						if err := checkAndTriggerBatchPreparation(blockGroupCtx, block); err != nil {
							if err != context.Canceled && err != context.DeadlineExceeded {
								log.Errorf("Error preparing batch for block %d: %s", currentNum, err)
							}
							return err
						}
						return nil
					})

					// Store block hash in Redis concurrently
					blockGroup.Go(func() error {
						retryOp := func() error {
							redisCtx, cancel := context.WithTimeout(blockGroupCtx, redisOperationTimeout)
							defer cancel()
							return redis.SetWithExpiration(redisCtx,
								redis.BlockHashByNumber(currentNum),
								block.Hash().Hex(),
								30*time.Minute)
						}

						if err := backoff.Retry(retryOp, backoff.WithContext(backoff.NewExponentialBackOff(), blockGroupCtx)); err != nil {
							log.Errorf("Failed to store block hash for block %d: %s", currentNum, err)
							return err
						}
						return nil
					})

					// Wait for this block's operations to complete
					if err := blockGroup.Wait(); err != nil {
						log.Errorf("Error in block %d processing: %s", currentNum, err)
						return err
					}

					// Update last processed block atomically
					if currentNum > lastProcessedBlock {
						lastProcessedBlock = currentNum
					}

					if currentNum < currentBlockNum {
						log.Infof("Catching up: Processed block %d, %d more to go", currentNum, currentBlockNum-currentNum)
					}

					return nil
				})
			}
			// Don't wait here - let blocks process independently
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
