package prost

import (
	"context"
	"math/big"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs/redis"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
)

var lastProcessedBlock int64

// StartFetchingBlocks continuously fetches blocks and processes events
func StartFetchingBlocks(ctx context.Context) {
	log.Println("Submission Event Collector started")

	// Create a child context for the cleanup routine
	cleanupCtx, cancel := context.WithCancel(ctx)
	defer cancel() // Will execute when StartFetchingBlocks returns

	// Start the periodic cleanup routine
	go startPeriodicCleanupRoutine(cleanupCtx)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Info("Shutting down block fetching")
			return
		case <-ticker.C:
			latestBlock, err := fetchBlock(ctx, nil)
			if err != nil {
				log.Errorf("Error fetching latest block: %s", err)
				continue
			}

			if err := processBlock(ctx, latestBlock); err != nil {
				log.Errorf("Error processing block: %s", err)
			}
		}
	}
}

// fetchBlock retrieves a block from the client using retry logic
func fetchBlock(ctx context.Context, blockNumber *big.Int) (*types.Block, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

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

func processBlock(ctx context.Context, block *types.Block) error {
	if block == nil {
		log.Errorf("Received nil block")
		return nil
	}

	latestBlockNumber := block.Number().Int64()

	// Check if lastProcessedBlock is set, if not set it to the latest block
	if lastProcessedBlock == 0 {
		lastProcessedBlock = latestBlockNumber
	}

	for blockNum := lastProcessedBlock + 1; blockNum <= latestBlockNumber; blockNum++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			block, err := fetchBlock(ctx, big.NewInt(blockNum))
			if err != nil {
				log.Errorf("Error fetching block %d: %s", blockNum, err)
				continue
			}

			if block == nil {
				log.Errorf("Received nil block for number: %d", blockNum)
				continue
			}

			log.Debugf("Processing block: %d", blockNum)

			// Launch goroutines with their own timeouts
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
				defer cancel()
				 checkAndTriggerBatchPreparation(ctx, block)
			}()

			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
				defer cancel()
				ProcessEvents(ctx, block)
			}()

			// Redis operation with its own timeout
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if err := redis.SetWithExpiration(ctx, redis.BlockHashByNumber(blockNum), block.Hash().Hex(), 30*time.Minute); err != nil {
					log.Errorf("Failed to set block hash for block number %d in Redis: %s", blockNum, err)
				}
			}()

			// Update last processed block
			lastProcessedBlock = blockNum
		}
	}

	// Sleep for approximately half the expected block time to balance load and responsiveness.
	time.Sleep(time.Duration(config.SettingsObj.BlockTime*500) * time.Millisecond)

	return nil
}
