package prost

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"submission-sequencer-collector/pkgs"
	"submission-sequencer-collector/pkgs/redis"

	log "github.com/sirupsen/logrus"
)

func CleanupSubmissionSet(ctx context.Context, dataMarketAddr string) error {
	log.Debugln("Cleaning up old submission set by header key for", dataMarketAddr)
	// Get current epoch
	currentEpochStr, err := redis.Get(ctx, redis.CurrentEpoch(dataMarketAddr))
	if err != nil {
		log.Errorf("Failed to get current epoch for cleanup: %v", err)
		return err
	}

	currentEpoch, err := strconv.ParseUint(currentEpochStr, 10, 64)
	if err != nil {
		log.Errorf("Failed to parse current epoch %s: %v", currentEpochStr, err)
		return err
	}

	// Use SCAN instead of KEYS
	var cursor uint64
	var keys []string
	pattern := fmt.Sprintf("%s.%s.*", pkgs.CollectorKey, strings.ToLower(dataMarketAddr))
	log.Debugf("Scanning for keys with pattern: %s for data market %s", pattern, dataMarketAddr)
	for {
		var batch []string
		var err error
		batch, cursor, err = redis.RedisClient.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			log.Errorf("Failed to scan submission set keys: %v", err)
			return err
		}
		log.Debugf("Found %d submission set by header keys for data market %s", len(batch), dataMarketAddr)
		// extract epoch ID from key
		for _, key := range batch {
			epochID := strings.Split(key, ".")[2]
			epochIDInt, err := strconv.ParseUint(epochID, 10, 64)
			if err != nil {
				log.Errorf("Failed to parse epoch ID %s: %v", epochID, err)
				continue
			}
			if epochIDInt < currentEpoch-10 {
				keys = append(keys, key)
			}
		}
		keys = append(keys, batch...)

		// If cursor is 0, we've completed the scan
		if cursor == 0 {
			break
		}
	}
	log.Debugf("Found total %d eligible for removal: submission set by header keys for data market %s", len(keys), dataMarketAddr)
	// Delete the keys in batch
	if len(keys) > 0 {
		if err := redis.RedisClient.Del(ctx, keys...).Err(); err != nil {
			log.Errorf("Failed to delete keys in batch: %v", err)
		} else {
			log.Debugf("Successfully deleted %d submission set header keys for data market %s", len(keys), dataMarketAddr)
		}
	}
	return nil
}
