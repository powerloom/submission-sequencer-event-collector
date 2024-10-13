package prost

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"submission-sequencer-collector/pkgs"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

// Function to fetch epoch markers from Redis
func FetchEpochMarkers(redisClient *redis.Client) ([]string, error) {
	// Fetch all epoch marker keys from Redis
	redisKeys, err := redisClient.Keys(context.Background(), fmt.Sprintf("%s.*", pkgs.EpochMarkerKey)).Result()
	if err != nil {
		return nil, err
	}

	return redisKeys, nil
}

// Test case for FetchEpochMarkers function
func TestFetchEpochMarkers(t *testing.T) {
	// Create a mini Redis server for testing
	mockRedis, err := miniredis.Run()
	assert.NoError(t, err)
	defer mockRedis.Close()

	// Initialize Redis client with mock server
	redisClient := redis.NewClient(&redis.Options{
		Addr: mockRedis.Addr(),
	})

	// Set up the expected Redis keys in the mini Redis server
	expectedKeys := make([]string, 5)
	for count := 0; count < 5; count++ {
		key := fmt.Sprintf("%s.%s", pkgs.EpochMarkerKey, strconv.Itoa(count))
		mockRedis.Set(key, "some_value") // Set dummy values to the keys
		expectedKeys[count] = key        // Store the expected key
	}

	// Set up the expected Redis keys in the mini Redis server
	for count := 0; count < 5; count++ {
		mockRedis.Set(fmt.Sprintf("%s.%s", pkgs.EpochMarkerKey, strconv.Itoa(count)), "some_value") // Set dummy values to the keys
	}

	// Call the function to fetch epoch markers
	actualKeys, err := FetchEpochMarkers(redisClient)
	assert.NoError(t, err)

	// Extract key by trimming the prefix from the Redis key
	for _, key := range actualKeys {
		updatedKey := strings.TrimPrefix(key, fmt.Sprintf("%s.", pkgs.EpochMarkerKey))
		fmt.Println("Updated key after trimming: ", updatedKey)
	}

	// Check if the function returned the expected keys
	assert.Equal(t, expectedKeys, actualKeys)
}
