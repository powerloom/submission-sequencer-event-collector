package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var SettingsObj *Settings

type DataMarketMigrationEntry struct {
	OldMarketAddress common.Address
	NewMarketAddress common.Address
}

type Settings struct {
	ClientUrl                   string
	ContractAddress             string
	RedisHost                   string
	RedisPort                   string
	RedisDB                     string
	AuthReadToken               string
	SlackReportingUrl           string
	TxRelayerUrl                string
	TxRelayerAuthWriteToken     string
	APIHost                     string
	DataMarketAddresses         []string
	DataMarketContractAddresses []common.Address
	BatchSize                   int
	BlockTime                   int
	HttpTimeout                 int
	PeriodicEligibleCountAlerts bool
	PastDaysBuffer              int
	RetryLimits                 int
	RewardsUpdateBatchSize      int
	RewardsUpdateEpochInterval  int64
	AttestorQueuePushEnabled    bool
	InitCleanupEnabled          bool
	ContractQueryTimeout        int64
	BlockFetchTimeout           int64
	EventProcessingTimeout      int64
	BatchProcessingTimeout      int64
	MemoryProfilingInterval     int
	DataMarketMigration         struct {
		Enabled       bool
		Mappings      []DataMarketMigrationEntry
		DaysToMigrate int
	}
}

func LoadConfig() {
	dataMarketAddresses := getEnv("DATA_MARKET_ADDRESSES", "[]")
	dataMarketAddressesList := []string{}

	err := json.Unmarshal([]byte(dataMarketAddresses), &dataMarketAddressesList)
	if err != nil {
		log.Fatalf("Failed to parse DATA_MARKET_ADDRESSES environment variable: %v", err)
	}
	if len(dataMarketAddressesList) == 0 {
		log.Fatalf("DATA_MARKET_ADDRESSES environment variable has an empty array")
	}

	periodicEligibleCountAlerts, periodicEligibleCountAlertsErr := strconv.ParseBool(getEnv("PERIODIC_ELIGIBLE_COUNT_ALERTS", "true"))
	if periodicEligibleCountAlertsErr != nil {
		log.Fatalf("Failed to parse PERIODIC_ELIGIBLE_COUNT_ALERTS environment variable: %v", periodicEligibleCountAlertsErr)
	}

	attestorQueuePushEnabled, attestorQueuePushEnabledErr := strconv.ParseBool(getEnv("ATTESTOR_QUEUE_PUSH_ENABLED", "false"))
	if attestorQueuePushEnabledErr != nil {
		log.Fatalf("Failed to parse ATTESTOR_QUEUE_PUSH_ENABLED environment variable: %v", attestorQueuePushEnabledErr)
	}

	initCleanupEnabled, initCleanupEnabledErr := strconv.ParseBool(getEnv("INIT_CLEANUP_ENABLED", "false"))
	if initCleanupEnabledErr != nil {
		log.Fatalf("Failed to parse INIT_CLEANUP_ENABLED environment variable: %v", initCleanupEnabledErr)
	}

	// Migration settings
	migrationEnabled := getEnv("ENABLE_MARKET_MIGRATION", "false") == "true"
	migrationMappings := getEnv("MARKET_MIGRATION_MAPPINGS", "")
	daysToMigrate, daysToMigrateErr := strconv.Atoi(getEnv("MARKET_MIGRATION_DAYS", "1"))
	if daysToMigrateErr != nil {
		log.Printf("Invalid MARKET_MIGRATION_DAYS value, using default of 1: %v", daysToMigrateErr)
		daysToMigrate = 1
	}
	if daysToMigrate < 1 {
		log.Printf("MARKET_MIGRATION_DAYS must be at least 1, using default of 1")
		daysToMigrate = 1
	}

	config := Settings{
		ClientUrl:                   getEnv("PROST_RPC_URL", ""),
		ContractAddress:             getEnv("PROTOCOL_STATE_CONTRACT", ""),
		RedisHost:                   getEnv("REDIS_HOST", ""),
		RedisPort:                   getEnv("REDIS_PORT", ""),
		RedisDB:                     getEnv("REDIS_DB", ""),
		AuthReadToken:               getEnv("AUTH_READ_TOKEN", ""),
		SlackReportingUrl:           getEnv("SLACK_REPORTING_URL", ""),
		TxRelayerUrl:                getEnv("TX_RELAYER_URL", ""),
		TxRelayerAuthWriteToken:     getEnv("TX_RELAYER_AUTH_WRITE_TOKEN", ""),
		DataMarketAddresses:         dataMarketAddressesList,
		PeriodicEligibleCountAlerts: periodicEligibleCountAlerts,
		APIHost:                     getEnv("API_HOST", ""),
		AttestorQueuePushEnabled:    attestorQueuePushEnabled,
		InitCleanupEnabled:          initCleanupEnabled,
	}

	if config.AuthReadToken == "" {
		log.Fatalf("AUTH_READ_TOKEN environment variable is not set")
	}

	for _, addr := range config.DataMarketAddresses {
		config.DataMarketContractAddresses = append(config.DataMarketContractAddresses, common.HexToAddress(addr))
	}

	batchSize, batchSizeParseErr := strconv.Atoi(getEnv("BATCH_SIZE", ""))
	if batchSizeParseErr != nil {
		log.Fatalf("Failed to parse BATCH_SIZE environment variable: %v", batchSizeParseErr)
	}
	config.BatchSize = batchSize

	rewardsUpdateBatchSize, rewardsUpdateBatchSizeParseErr := strconv.Atoi(getEnv("REWARDS_UPDATE_BATCH_SIZE", ""))
	if rewardsUpdateBatchSizeParseErr != nil {
		log.Fatalf("Failed to parse REWARDS_UPDATE_BATCH_SIZE environment variable: %v", rewardsUpdateBatchSizeParseErr)
	}
	config.RewardsUpdateBatchSize = rewardsUpdateBatchSize

	rewardsUpdateEpochInterval, rewardsUpdateEpochIntervalParseErr := strconv.Atoi(getEnv("REWARDS_UPDATE_EPOCH_INTERVAL", ""))
	if rewardsUpdateEpochIntervalParseErr != nil {
		log.Fatalf("Failed to parse REWARDS_UPDATE_EPOCH_INTERVAL environment variable: %v", rewardsUpdateEpochIntervalParseErr)
	}
	config.RewardsUpdateEpochInterval = int64(rewardsUpdateEpochInterval)

	blockTime, blockTimeParseErr := strconv.Atoi(getEnv("BLOCK_TIME", ""))
	if blockTimeParseErr != nil {
		log.Fatalf("Failed to parse BLOCK_TIME environment variable: %v", blockTimeParseErr)
	}
	config.BlockTime = blockTime

	httpTimeout, timeoutParseErr := strconv.Atoi(getEnv("HTTP_TIMEOUT", ""))
	if timeoutParseErr != nil {
		log.Fatalf("Failed to parse HTTP_TIMEOUT environment variable: %v", timeoutParseErr)
	}
	config.HttpTimeout = httpTimeout

	pastDaysBuffer, pastDaysBufferParseErr := strconv.Atoi(getEnv("PAST_DAYS_BUFFER", "5"))
	if pastDaysBufferParseErr != nil {
		log.Fatalf("Failed to parse PAST_DAYS_BUFFER environment variable: %v", pastDaysBufferParseErr)
	}
	config.PastDaysBuffer = pastDaysBuffer

	retryLimits, retryLimitsParseErr := strconv.Atoi(getEnv("RETRY_LIMITS", "3"))
	if retryLimitsParseErr != nil {
		log.Fatalf("Failed to parse RETRY_LIMITS environment variable: %v", retryLimitsParseErr)
	}
	config.RetryLimits = retryLimits

	contractQueryTimeout, contractQueryTimeoutParseErr := strconv.ParseInt(getEnv("CONTRACT_QUERY_TIMEOUT", ""), 10, 64)
	if contractQueryTimeoutParseErr != nil {
		log.Fatalf("Failed to parse CONTRACT_QUERY_TIMEOUT environment variable: %v", contractQueryTimeoutParseErr)
	}
	config.ContractQueryTimeout = contractQueryTimeout

	blockFetchTimeout, blockFetchTimeoutErr := strconv.ParseInt(getEnv("BLOCK_FETCH_TIMEOUT", ""), 10, 64)
	if blockFetchTimeoutErr != nil {
		log.Fatalf("Failed to parse BLOCK_FETCH_TIMEOUT environment variable: %v", blockFetchTimeoutErr)
	}
	config.BlockFetchTimeout = blockFetchTimeout

	eventProcessingTimeout, eventTimeoutErr := strconv.ParseInt(getEnv("EVENT_PROCESSING_TIMEOUT", ""), 10, 64)
	if eventTimeoutErr != nil {
		log.Fatalf("Failed to parse EVENT_PROCESSING_TIMEOUT environment variable: %v", eventTimeoutErr)
	}
	config.EventProcessingTimeout = eventProcessingTimeout

	batchProcessingTimeout, batchTimeoutErr := strconv.ParseInt(getEnv("BATCH_PROCESSING_TIMEOUT", ""), 10, 64)
	if batchTimeoutErr != nil {
		log.Fatalf("Failed to parse BATCH_PROCESSING_TIMEOUT environment variable: %v", batchTimeoutErr)
	}
	config.BatchProcessingTimeout = batchProcessingTimeout

	memoryProfilingInterval, memoryProfilingIntervalParseErr := strconv.Atoi(getEnv("MEMORY_PROFILING_INTERVAL", "15"))
	if memoryProfilingIntervalParseErr != nil {
		log.Fatalf("Failed to parse MEMORY_PROFILING_INTERVAL environment variable: %v", memoryProfilingIntervalParseErr)
	}
	config.MemoryProfilingInterval = memoryProfilingInterval

	if migrationEnabled {
		if migrationMappings == "" {
			log.Fatal("Migration is enabled but no mappings are configured")
		}

		// Parse mappings in format "old1:new1,old2:new2,..."
		mappings := strings.Split(migrationMappings, ",")
		for _, mapping := range mappings {
			addresses := strings.Split(mapping, ":")
			if len(addresses) != 2 {
				log.Fatalf("Invalid migration mapping format: %s", mapping)
			}

			config.DataMarketMigration.Mappings = append(
				config.DataMarketMigration.Mappings,
				DataMarketMigrationEntry{
					OldMarketAddress: common.HexToAddress(addresses[0]),
					NewMarketAddress: common.HexToAddress(addresses[1]),
				},
			)
		}
		config.DataMarketMigration.Enabled = true
		config.DataMarketMigration.DaysToMigrate = daysToMigrate
	}

	SettingsObj = &config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
