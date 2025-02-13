package prost

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"submission-sequencer-collector/pkgs"
	"submission-sequencer-collector/pkgs/redis"
	"time"

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
	log.Debugf("Current epoch cached query result for data market %s: %s", dataMarketAddr, currentEpochStr)
	if currentEpochStr == "" {
		log.Errorf("Current epoch is empty for data market %s", dataMarketAddr)
		return errors.New("current epoch is empty")
	}

	currentEpoch, err := strconv.Atoi(currentEpochStr)
	if err != nil {
		log.Errorf("Failed to parse current epoch %s: %v", currentEpochStr, err)
		return err
	}

	// Use SCAN instead of KEYS
	var cursor uint64
	var keys []string
	pattern := fmt.Sprintf("%s.%s*", pkgs.CollectorKey, strings.ToLower(dataMarketAddr))
	log.Debugf("Scanning for keys with pattern: %s for data market %s", pattern, dataMarketAddr)
	for {
		var batch []string
		var err error
		batch, cursor, err = redis.RedisClient.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			log.Errorf("Failed to scan submission set keys: %v", err)
			return err
		}

		// extract epoch ID from key
		for _, key := range batch {
			epochID := strings.Split(key, ".")[2]
			epochIDInt, err := strconv.Atoi(epochID)
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

func CleanupSubmissionDumpForAllSlots(dataMarketAddr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Debugf("Cleaning up old submission dump for all slots for data market %s", dataMarketAddr)

	// get current epoch
	currentEpochStr, err := redis.Get(ctx, redis.CurrentEpoch(dataMarketAddr))
	if err != nil {
		log.Errorf("Failed to get current epoch for cleanup: %v", err)
		return err
	}

	if currentEpochStr == "" {
		log.Errorf("Current epoch is empty for data market %s", dataMarketAddr)
		return errors.New("current epoch is empty")
	}

	currentEpoch, err := strconv.Atoi(currentEpochStr)
	if err != nil {
		log.Errorf("Failed to parse current epoch %s: %v", currentEpochStr, err)
		return err
	}

	//get node count
	nodeCountStr, err := redis.Get(ctx, redis.TotalNodesCountKey())
	if err != nil {
		log.Errorf("Failed to get node count for cleanup: %v", err)
		return err
	}
	nodeCount, err := strconv.Atoi(nodeCountStr)
	if err != nil {
		log.Errorf("Failed to parse node count %s: %v", nodeCountStr, err)
		return err
	}
	for slotId := 1; slotId <= nodeCount; slotId++ {
		log.Debugf("Cleaning up old submission dump for slot %d for data market %s", slotId, dataMarketAddr)
		keyPattern := fmt.Sprintf("snapshotter:%s:*:%d.slot_submissions", strings.ToLower(dataMarketAddr), slotId)
		var cursor uint64
		var keys []string
		for {
			var batch []string
			var err error
			batch, cursor, err = redis.RedisClient.Scan(ctx, cursor, keyPattern, 100).Result()
			if err != nil {
				log.Errorf("Failed to scan submission dump keys: %v", err)
				return err
			}
			keys = append(keys, batch...)
			if cursor == 0 {
				break
			}
		}
		if len(keys) > 0 {
			// delete epoch IDs less than current epoch - 30
			// each key is a hash table with epoch IDs as keys
			for _, key := range keys {
				pipe := redis.RedisClient.Pipeline()
				for epochToDelete := currentEpoch - 30; epochToDelete <= 1; epochToDelete-- {
					pipe.HDel(ctx, key, strconv.Itoa(epochToDelete))
				}
				_, err := pipe.Exec(ctx)
				if err != nil {
					log.Errorf("Failed to execute pipeline for key %s: %v", key, err)
				} else {
					log.Debugf("Successfully deleted epoch IDs from %d to 1 for key %s", currentEpoch-30, key)
				}
			}
		}
	}

	return nil
}
