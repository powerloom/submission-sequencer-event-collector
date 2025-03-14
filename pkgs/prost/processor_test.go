package prost

import (
	"context"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs/clients"
	"submission-sequencer-collector/pkgs/redis"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
)

func setup(t *testing.T) *miniredis.Miniredis {
	// Initialize a miniredis instance
	mockRedis, err := miniredis.Run()
	if err != nil {
		t.Fatalf("Failed to start miniredis: %v", err)
	}

	// Configure the application to use miniredis
	config.SettingsObj = &config.Settings{
		RedisHost: mockRedis.Host(),
		RedisPort: mockRedis.Port(),
		RedisDB:   "0",
	}

	clients.InitializeReportingClient(config.SettingsObj.SlackReportingUrl, time.Duration(config.SettingsObj.HttpTimeout)*time.Second)

	// Initialize the Redis client
	redis.RedisClient = redis.NewRedisClient()

	// Use t.Cleanup to ensure the miniredis instance is closed after the test
	t.Cleanup(func() {
		redis.RedisClient.FlushAll(context.Background())
		mockRedis.Close()
	})

	return mockRedis
}
