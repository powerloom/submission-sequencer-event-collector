package prost

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs"
	"submission-sequencer-collector/pkgs/clients"
	"submission-sequencer-collector/pkgs/redis"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// getValidSubmissionKeys retrieves valid submission keys for the given epoch and headers
func getValidSubmissionKeys(ctx context.Context, epochID uint64, headers []string, dataMarketAddress string) ([]string, error) {
	// Store valid submission keys
	var submissionKeys []string

	// Process each header to get its submission keys
	for _, header := range headers {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			// Get all submission keys for this header from Redis
			keys := redis.RedisClient.SMembers(ctx, redis.SubmissionSetByHeaderKey(dataMarketAddress, epochID, header)).Val()
			if len(keys) > 0 {
				submissionKeys = append(submissionKeys, keys...)
			}
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
			errMsg := fmt.Sprintf("Improper submission key stored in Redis: %s", submissionKey)
			clients.SendFailureNotification(pkgs.ConstructProjectMap, errMsg, time.Now().String(), "High")
			log.Error(errMsg)
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

// fetchEligibleSlotIDs returns the slot IDs and their count for a given data market and day.
// SlotIDs with eligible submission counts greater than equal to daily snapshot quota are stored.
func fetchEligibleSlotIDs(ctx context.Context, dataMarketAddress, day string) (int, []string) {
	// Build the Redis key to fetch the slotIDs for the specified day
	eligibleNodesSetKey := redis.EligibleNodesByDayKey(dataMarketAddress, day)

	// Retrieve the slot IDs stored in the set associated with the Redis key
	slotIDs := redis.GetSetKeys(ctx, eligibleNodesSetKey)

	// Return the slot IDs and their count
	return len(slotIDs), slotIDs
}

// startPeriodicCleanupRoutine runs a periodic cleanup every 10 minutes
func startPeriodicCleanupRoutine(cleanupCtx context.Context) {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-cleanupCtx.Done():
			log.Info("⏹️ Periodic cleanup routine stopped")
			return
		case <-ticker.C:
			// Create timeout context for block fetch
			fetchCtx, cancel := context.WithTimeout(cleanupCtx, 30*time.Second)
			// Fetch the latest block
			defer cancel()
			currentBlock, err := fetchBlock(fetchCtx, nil)

			if err != nil {
				log.Errorf("Failed to fetch the latest block during cleanup routine: %s", err)
				continue
			}

			currentBlockNum := currentBlock.Number().Int64()
			log.Infof("Starting periodic cleanup for stale epoch markers at block number: %d", currentBlockNum)

			// Trigger the cleanup process
			startPeriodicCleanup(currentBlockNum)
		}
	}
}

// startPeriodicCleanup cleans up stale epoch markers
func startPeriodicCleanup(currentBlockNum int64) {
	log.Infof("🧹 Starting periodic cleanup for stale epoch markers at block number: %d", currentBlockNum)
	var wg sync.WaitGroup

	// Cleanup for each data market in parallel
	for _, dataMarketAddress := range config.SettingsObj.DataMarketAddresses {
		wg.Add(1)

		go func(dataMarketAddress string) {
			defer wg.Done()

			// Create new context with appropriate timeout
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			log.Infof("🏁 Starting cleanup for stale epoch markers for data market %s at block number: %d",
				dataMarketAddress, currentBlockNum)

			epochMarkerKeys, err := redis.RedisClient.SMembers(ctx, redis.EpochMarkerSet(dataMarketAddress)).Result()
			if err != nil {
				log.Errorf("Failed to fetch epoch markers for data market %s during cleanup: %s", dataMarketAddress, err)
				return
			}

			// Process each epoch marker key for this data market address
			for _, epochMarkerKey := range epochMarkerKeys {
				// Retrieve the epoch marker details from Redis
				epochMarkerDetailsJSON, err := redis.RedisClient.Get(ctx, redis.EpochMarkerDetails(dataMarketAddress, epochMarkerKey)).Result()
				if err != nil {
					log.Errorf("Failed to fetch epoch marker details for key %s during cleanup: %s", epochMarkerKey, err)
					continue
				}

				var epochMarkerDetails EpochMarkerDetails
				if err := json.Unmarshal([]byte(epochMarkerDetailsJSON), &epochMarkerDetails); err != nil {
					log.Errorf("Failed to unmarshal epoch marker details for key %s during cleanup: %s", epochMarkerKey, err)
					continue
				}

				// Remove stale epoch markers if the submission limit block has passed
				if currentBlockNum > epochMarkerDetails.SubmissionLimitBlockNumber {
					log.Infof("🗑️ Removing stale epoch marker key %s for data market %s", epochMarkerKey, dataMarketAddress)

					// Remove the epochID and its details from Redis
					if err := redis.RemoveEpochFromRedis(ctx, dataMarketAddress, epochMarkerKey); err != nil {
						log.Errorf("Failed to remove epoch %s from Redis for data market %s during cleanup: %v", epochMarkerKey, dataMarketAddress, err)
						continue
					}

					log.Infof("✅ Successfully removed stale epoch marker key %s for data market %s", epochMarkerKey, dataMarketAddress)
				}
			}
		}(dataMarketAddress)
	}

	// Wait for all data market goroutines to finish
	wg.Wait()

	log.Infof("🧹 Completed cleanup for all stale epoch markers")
}
