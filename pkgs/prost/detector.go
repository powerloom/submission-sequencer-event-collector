package prost

import (
	"context"
	"fmt"
	"math/big"
	"runtime"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs/redis"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// Context timeout constants
const (
	// Operation-specific timeouts
	blockFetchTimeout     = 5 * time.Second
	blockProcessTimeout   = 30 * time.Second
	eventProcessTimeout   = 30 * time.Second
	batchPrepareTimeout   = 60 * time.Second
	redisOperationTimeout = 5 * time.Second
)

var (
	lastProcessedBlock int64
	// Worker pool for block processing
	blockWorkerPool chan struct{}
	// Metrics for monitoring
	activeGoroutines sync.WaitGroup
	// Resource cleanup on shutdown
	shutdownOnce sync.Once
	// Rate limiter for block processing
	blockProcessingLimiter = time.NewTicker(100 * time.Millisecond)
)

// Initialize worker pools and resources
func init() {
	// Set worker pool size based on CPU cores but with a reasonable maximum
	maxWorkers := runtime.GOMAXPROCS(0) * 2
	if maxWorkers > 16 {
		maxWorkers = 16 // Cap maximum workers to prevent CPU spikes
	}
	blockWorkerPool = make(chan struct{}, maxWorkers)
}

// Cleanup resources
func cleanup() {
	shutdownOnce.Do(func() {
		blockProcessingLimiter.Stop()
		close(blockWorkerPool)
	})
}

// StartFetchingBlocks continuously fetches blocks and processes events
func StartFetchingBlocks(ctx context.Context) {
	// Parent context is used for shutdown signaling
	defer cleanup()
	log.Info("Submission Event Collector started")

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Info("Shutting down block fetching: context cancelled")
			activeGoroutines.Wait()
			return
		case <-ticker.C:
			select {
			case blockWorkerPool <- struct{}{}:
				activeGoroutines.Add(1)
				processBlock := func() {
					defer func() {
						<-blockWorkerPool
						activeGoroutines.Done()
					}()

					// Create fetch context as child of parent context
					fetchCtx, fetchCancel := context.WithTimeout(ctx, blockFetchTimeout)
					defer fetchCancel()

					latestBlock, err := fetchBlock(fetchCtx, nil)
					if err != nil {
						if err != context.Canceled && err != context.DeadlineExceeded {
							log.Errorf("Error fetching latest block: %s", err)
						}
						return
					}

					// Process block with its own context
					processCtx, processCancel := context.WithTimeout(ctx, blockProcessTimeout)
					defer processCancel()

					if err := processBlock(processCtx, latestBlock); err != nil {
						if err != context.Canceled && err != context.DeadlineExceeded {
							log.Errorf("Error processing block: %s", err)
						}
					}
				}

				// Launch the block processing
				go processBlock()
			default:
				log.Warn("Worker pool is full, skipping block processing")
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

// processBlock processes a single block for batch preparation and event processing.
func processBlock(ctx context.Context, block *types.Block) error {
	if block == nil {
		return fmt.Errorf("received nil block")
	}

	latestBlockNumber := block.Number().Int64()
	if lastProcessedBlock == 0 {
		lastProcessedBlock = latestBlockNumber
	}

	// Use a bounded buffer for block processing
	const maxBlocksPerBatch = 100
	blocksToProcess := latestBlockNumber - lastProcessedBlock
	if blocksToProcess > maxBlocksPerBatch {
		log.Warnf("Large block gap detected (%d blocks). Processing only last %d blocks to prevent memory spikes",
			blocksToProcess, maxBlocksPerBatch)
		lastProcessedBlock = latestBlockNumber - maxBlocksPerBatch
	}

	// Create error group with parent context
	g, groupCtx := errgroup.WithContext(ctx)

	// Process blocks in batches
	for blockNum := lastProcessedBlock + 1; blockNum <= latestBlockNumber; blockNum++ {
		blockNum := blockNum // Create new variable for goroutine

		select {
		case <-groupCtx.Done():
			return groupCtx.Err()
		default:
			g.Go(func() error {
				// Fetch and process individual block
				block, err := fetchBlock(groupCtx, big.NewInt(blockNum))
				if err != nil {
					return fmt.Errorf("failed to fetch block %d: %w", blockNum, err)
				}

				if block == nil {
					return fmt.Errorf("received nil block for number: %d", blockNum)
				}

				// Process events with timeout
				if err := ProcessEvents(groupCtx, block); err != nil {
					return fmt.Errorf("failed to process events for block %d: %w", blockNum, err)
				}

				// Prepare batch
				if err := checkAndTriggerBatchPreparation(groupCtx, block); err != nil {
					return fmt.Errorf("failed to prepare batch for block %d: %w", blockNum, err)
				}

				// Store block hash in Redis with retry
				retryOp := func() error {
					// Use a shorter timeout for Redis operations
					redisCtx, cancel := context.WithTimeout(groupCtx, redisOperationTimeout)
					defer cancel()

					return redis.SetWithExpiration(redisCtx,
						redis.BlockHashByNumber(blockNum),
						block.Hash().Hex(),
						30*time.Minute)
				}

				if err := backoff.Retry(retryOp, backoff.WithContext(backoff.NewExponentialBackOff(), groupCtx)); err != nil {
					return fmt.Errorf("failed to store block hash %d in Redis: %w", blockNum, err)
				}

				// Update last processed block only if all operations succeeded
				lastProcessedBlock = blockNum
				return nil
			})
		}
	}

	// Wait for all block processing to complete
	if err := g.Wait(); err != nil {
		return fmt.Errorf("block processing failed: %w", err)
	}

	// Add delay to prevent excessive CPU usage while staying responsive
	if latestBlockNumber-lastProcessedBlock > 10 {
		time.Sleep(time.Duration(config.SettingsObj.BlockTime*100) * time.Millisecond)
	} else {
		time.Sleep(time.Duration(config.SettingsObj.BlockTime*500) * time.Millisecond)
	}

	return nil
}
