package prost

import (
	"context"
	"fmt"
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
		windowManager.Shutdown()
	})
}

// StartFetchingBlocks continuously fetches blocks and processes events
func StartFetchingBlocks(ctx context.Context) {
	defer cleanup()
	log.Info("Submission Event Collector started")

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Create a semaphore to limit concurrent block processing
	blockSemaphore := make(chan struct{}, 3) // Limit to 3 concurrent blocks

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

				// Acquire semaphore with timeout
				select {
				case blockSemaphore <- struct{}{}:
					// Got permission to proceed
				case <-ctx.Done():
					return
				case <-time.After(5 * time.Second):
					log.Warnf("⚠️ Timeout waiting for block semaphore for block %d", currentNum)
					continue
				}

				// Create block processing context
				blockCtx, blockCancel := context.WithTimeout(ctx, blockFetchTimeout)
				defer blockCancel()

				// Fetch the block
				block, err := fetchBlock(blockCtx, big.NewInt(currentNum))
				if err != nil {
					log.Errorf("failed to fetch block %d: %s", currentNum, err)
					<-blockSemaphore // Release semaphore on error
					continue
				}

				// error collection channel
				errChan := make(chan error, 3)

				// Process events and prepare batch concurrently
				var wg sync.WaitGroup

				// Process Events
				wg.Add(1)
				eventProcessCtx, eventProcessCancel := context.WithTimeout(ctx, eventProcessingTimeout)
				go func() {
					defer wg.Done()
					defer eventProcessCancel()
					if err := ProcessEvents(eventProcessCtx, block); err != nil {
						select {
						case errChan <- fmt.Errorf("failed to process events for block %d: %v", currentNum, err):
						case <-eventProcessCtx.Done():
						}
					}
				}()

				// Store block hash in Redis
				wg.Add(1)
				go func() {
					defer wg.Done()
					if err := redis.SetWithExpiration(ctx,
						redis.BlockHashByNumber(currentNum),
						block.Hash().Hex(),
						30*time.Minute); err != nil {
						select {
						case errChan <- fmt.Errorf("failed to store block hash for block %d: %v", currentNum, err):
						case <-ctx.Done():
						}
					}
				}()

				// Wait for all operations to complete
				done := make(chan struct{})
				go func() {
					wg.Wait()
					close(done)
					close(errChan)
					<-blockSemaphore // Release semaphore when done
				}()

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
