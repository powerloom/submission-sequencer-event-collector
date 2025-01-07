package prost

import (
	"context"
	"fmt"
	"strings"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs"
	"submission-sequencer-collector/pkgs/clients"
	"submission-sequencer-collector/pkgs/redis"
	"time"

	log "github.com/sirupsen/logrus"
)

func getValidSubmissionKeys(ctx context.Context, epochID uint64, headers []string, dataMarketAddress string) ([]string, error) {
	// Initialize an empty slice to store valid submission keys
	submissionKeys := make([]string, 0)

	// Iterate through the list of headers
	for _, header := range headers {
		keys := redis.RedisClient.SMembers(ctx, redis.SubmissionSetByHeaderKey(dataMarketAddress, epochID, header)).Val()
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
func fetchEligibleSlotIDs(dataMarketAddress, day string) (int, []string) {
	// Build the Redis key to fetch the slotIDs for the specified day
	eligibleNodesSetKey := redis.EligibleNodesByDayKey(dataMarketAddress, day)

	// Retrieve the slot IDs stored in the set associated with the Redis key
	slotIDs := redis.GetSetKeys(context.Background(), eligibleNodesSetKey)

	// Return the slot IDs and their count
	return len(slotIDs), slotIDs
}
