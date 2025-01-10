package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func NewRedisClient() *redis.Client {
	db, err := strconv.Atoi(config.SettingsObj.RedisDB)
	if err != nil {
		log.Fatalf("Incorrect redis db: %s", err.Error())
	}
	return redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.SettingsObj.RedisHost, config.SettingsObj.RedisPort), // Redis server address
		Password:     "",                                                                               // no password set
		DB:           db,
		PoolSize:     1000,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 200 * time.Millisecond,
		DialTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Minute,
	})
}

func AddToSet(ctx context.Context, set string, keys ...string) error {
	if err := RedisClient.SAdd(ctx, set, keys).Err(); err != nil {
		return fmt.Errorf("unable to add to set: %s", err.Error())
	}
	return nil
}

func GetSetKeys(ctx context.Context, set string) []string {
	return RedisClient.SMembers(ctx, set).Val()
}

func RemoveFromSet(ctx context.Context, set, key string) error {
	return RedisClient.SRem(context.Background(), set, key).Err()
}

func Delete(ctx context.Context, set string) error {
	return RedisClient.Del(ctx, set).Err()
}

func Expire(ctx context.Context, key string, expiration time.Duration) error {
	return RedisClient.Expire(ctx, key, expiration).Err()
}

func Get(ctx context.Context, key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		} else {
			return "", err
		}
	}
	return val, nil
}

func Set(ctx context.Context, key, value string) error {
	return RedisClient.Set(ctx, key, value, 0).Err()
}

// Use this when you want to set an expiration
func SetWithExpiration(ctx context.Context, key, value string, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

func Incr(ctx context.Context, key string) (int64, error) {
	result, err := RedisClient.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func GetBooleanValue(ctx context.Context, key string) (bool, error) {
	value, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return value == "true", nil
}

func SetBooleanValue(ctx context.Context, key string, value bool, ttl time.Duration) error {
	stringValue := "false"
	if value {
		stringValue = "true"
	}

	return RedisClient.Set(ctx, key, stringValue, ttl).Err()
}

func GetSetCardinality(ctx context.Context, key string) (int, error) {
	count, err := RedisClient.SCard(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func GetDaySize(ctx context.Context, dataMarketAddress string) (*big.Int, error) {
	// Fetch DAY_SIZE for the given data market address from Redis
	daySizeStr, err := RedisClient.HGet(context.Background(), GetDaySizeTableKey(), dataMarketAddress).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch day size for data market %s: %s", dataMarketAddress, err)
	}

	// Convert the day size from string to *big.Int
	daySize, ok := new(big.Int).SetString(daySizeStr, 10)
	if !ok {
		return nil, fmt.Errorf("invalid day size value for data market %s: %s", dataMarketAddress, daySizeStr)
	}

	return daySize, nil
}

func GetDailySnapshotQuota(ctx context.Context, dataMarketAddress string) (*big.Int, error) {
	// Fetch daily snapshot quota for the given data market address from Redis
	dailySnapshotQuotaStr, err := RedisClient.HGet(context.Background(), GetDailySnapshotQuotaTableKey(), dataMarketAddress).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch daily snapshot quota for data market %s: %s", dataMarketAddress, err)
	}

	// Convert the daily snapshot quota from string to *big.Int
	dailySnapshotQuota, ok := new(big.Int).SetString(dailySnapshotQuotaStr, 10)
	if !ok {
		return nil, fmt.Errorf("invalid daily snapshot quota value for data market %s: %s", dataMarketAddress, dailySnapshotQuotaStr)
	}

	return dailySnapshotQuota, nil
}

// StoreEpochDetails stores the epoch ID in the master set and its associated details in Redis
func StoreEpochDetails(ctx context.Context, dataMarketAddress, epochID, details string) error {
	// Store the epoch ID in the master set
	if err := AddToSet(ctx, EpochMarkerSet(dataMarketAddress), epochID); err != nil {
		return fmt.Errorf("failed to add epoch %s to master set for data market %s: %w", epochID, dataMarketAddress, err)
	}

	// Store the details for the specified epochID related to the given data market in Redis
	if err := Set(ctx, EpochMarkerDetails(dataMarketAddress, epochID), details); err != nil {
		return fmt.Errorf("failed to store epoch marker details for epoch %s, data market %s in Redis: %w", epochID, dataMarketAddress, err)
	}

	return nil
}

// StoreDayTransitionEpochDetails stores the epoch ID in the master set and its associated day transition details in Redis
func StoreDayTransitionEpochDetails(ctx context.Context, dataMarketAddress, epochID, details string) error {
	// Store the epoch ID in the master set
	if err := AddToSet(ctx, DayRolloverEpochMarkerSet(dataMarketAddress), epochID); err != nil {
		return fmt.Errorf("failed to add epoch %s to day roll over master set for data market %s: %w", epochID, dataMarketAddress, err)
	}

	// Store the day transition details for the specified epochID related to the given data market in Redis
	if err := Set(ctx, DayRolloverEpochMarkerDetails(dataMarketAddress, epochID), details); err != nil {
		return fmt.Errorf("failed to store day rollover epoch details for epoch %s, data market %s in Redis: %w", epochID, dataMarketAddress, err)
	}

	return nil
}

func RemoveEpochFromRedis(ctx context.Context, dataMarketAddress, epochID string) error {
	// Remove the epoch marker from the master set
	if err := RedisClient.SRem(ctx, EpochMarkerSet(dataMarketAddress), epochID).Err(); err != nil {
		return fmt.Errorf("failed to delete epoch %s for data market %s from Redis: %w", epochID, dataMarketAddress, err)
	}

	if expireErr := RedisClient.Expire(ctx, EpochMarkerDetails(dataMarketAddress, epochID), 30*time.Minute).Err(); expireErr != nil {
		return fmt.Errorf("failed to set expiry for epoch marker details of epoch %s, data market %s: %w", epochID, dataMarketAddress, expireErr)
	}

	return nil
}

func RemoveDayTransitionEpochFromRedis(ctx context.Context, dataMarketAddress, epochID string) error {
	// Remove the epoch marker from the day transition master set
	if err := RedisClient.SRem(ctx, DayRolloverEpochMarkerSet(dataMarketAddress), epochID).Err(); err != nil {
		return fmt.Errorf("failed to delete epoch %s from day transition master set for data market %s from Redis: %w", epochID, dataMarketAddress, err)
	}

	if expireErr := RedisClient.Expire(ctx, DayRolloverEpochMarkerDetails(dataMarketAddress, epochID), 30*time.Minute).Err(); expireErr != nil {
		return fmt.Errorf("failed to set expiry for day transition epoch marker details of epoch %s, data market %s: %w", epochID, dataMarketAddress, expireErr)
	}

	return nil
}

// StoreBatchDetails saves the details of a batch submission to Redis
func StoreBatchDetails(ctx context.Context, dataMarketAddress, epochID, batchID, details string) error {
	// Attempt to set the batch details in Redis with an expiration time of 5 minutes
	if err := SetWithExpiration(ctx, BatchSubmissionKey(dataMarketAddress, epochID, batchID), details, 5*time.Minute); err != nil {
		return fmt.Errorf("failed to store details for batch %s of epoch %s, data market %s in Redis: %w", batchID, epochID, dataMarketAddress, err)
	}

	return nil
}

func GetEpochsInADay(ctx context.Context, dataMarketAddress string) (*big.Int, error) {
	// Construct the key using ContractStateVariableWithDataMarket pattern
	key := ContractStateVariableWithDataMarket(dataMarketAddress, pkgs.EpochsInADay)

	// Fetch epochs in a day directly as a string value
	epochsInADayStr, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch epochs in a day for data market %s: %s", dataMarketAddress, err)
	}

	// Convert the epochs in a day from string to *big.Int
	epochsInADay, ok := new(big.Int).SetString(epochsInADayStr, 10)
	if !ok {
		return nil, fmt.Errorf("invalid epochs in a day value for data market %s: %s", dataMarketAddress, epochsInADayStr)
	}

	return epochsInADay, nil
}
