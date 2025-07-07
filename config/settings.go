package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	rpchelper "github.com/powerloom/go-rpc-helper"
	"github.com/powerloom/go-rpc-helper/reporting"

	"github.com/ethereum/go-ethereum/common"
)

var SettingsObj *Settings

type DataMarketMigrationEntry struct {
	OldMarketAddress common.Address
	NewMarketAddress common.Address
}

type Settings struct {
	// RPC Helper Configuration
	RPCNodes        []string `json:"rpc_nodes"`
	ArchiveRPCNodes []string `json:"archive_rpc_nodes"`
	MaxRetries      int      `json:"max_retries"`
	RetryDelayMs    int      `json:"retry_delay_ms"`
	MaxRetryDelayS  int      `json:"max_retry_delay_s"`
	RequestTimeoutS int      `json:"request_timeout_s"`

	// Legacy fields (keeping for backward compatibility during transition)
	ClientUrl                        string
	ContractAddress                  string
	RedisHost                        string
	RedisPort                        string
	RedisDB                          string
	AuthReadToken                    string
	SlackReportingUrl                string
	TxRelayerUrl                     string
	TxRelayerAuthWriteToken          string
	APIHost                          string
	DataMarketAddresses              []string
	DataMarketContractAddresses      []common.Address
	BatchSize                        int
	BlockTime                        int
	HttpTimeout                      int
	PeriodicEligibleCountAlerts      bool
	PastDaysBuffer                   int
	RetryLimits                      int
	RewardsUpdateBatchSize           int
	RewardsUpdateEpochInterval       int64
	AttestorQueuePushEnabled         bool
	InitCleanupEnabled               bool
	ConcurrentSubmissionCountUpdates int
	ContractQueryTimeout             int64
	BlockFetchTimeout                int64
	EventProcessingTimeout           int64
	BatchProcessingTimeout           int64
	MemoryProfilingInterval          int
	DataMarketMigration              struct {
		Enabled       bool
		Mappings      []DataMarketMigrationEntry
		DaysToMigrate int
	}
}

func LoadConfig() {
	// Parse RPC nodes from environment variable
	rpcNodesStr := getEnv("RPC_NODES", "[]")
	var rpcNodes []string
	err := json.Unmarshal([]byte(rpcNodesStr), &rpcNodes)
	if err != nil {
		log.Fatalf("Failed to parse RPC_NODES environment variable: %v", err)
	}
	if len(rpcNodes) == 0 {
		// Fallback to legacy PROST_RPC_URL for backward compatibility
		legacyRPCURL := getEnv("PROST_RPC_URL", "")
		if legacyRPCURL != "" {
			rpcNodes = []string{legacyRPCURL}
		} else {
			log.Fatalf("RPC_NODES environment variable has an empty array and no PROST_RPC_URL fallback")
		}
	}

	// Clean quotes from RPC node URLs
	for i, url := range rpcNodes {
		rpcNodes[i] = strings.Trim(url, "\"")
	}

	// Parse archive RPC nodes from environment variable (optional)
	archiveRPCNodesStr := getEnv("ARCHIVE_RPC_NODES", "[]")
	var archiveRPCNodes []string
	err = json.Unmarshal([]byte(archiveRPCNodesStr), &archiveRPCNodes)
	if err != nil {
		log.Fatalf("Failed to parse ARCHIVE_RPC_NODES environment variable: %v", err)
	}

	// Clean quotes from archive RPC node URLs
	for i, url := range archiveRPCNodes {
		archiveRPCNodes[i] = strings.Trim(url, "\"")
	}

	dataMarketAddresses := getEnv("DATA_MARKET_ADDRESSES", "[]")
	dataMarketAddressesList := []string{}

	err = json.Unmarshal([]byte(dataMarketAddresses), &dataMarketAddressesList)
	if err != nil {
		log.Fatalf("Failed to parse DATA_MARKET_ADDRESSES environment variable: %v", err)
	}
	if len(dataMarketAddressesList) == 0 {
		log.Fatalf("DATA_MARKET_ADDRESSES environment variable has an empty array")
	}

	// Clean quotes from data market addresses
	for i, addr := range dataMarketAddressesList {
		dataMarketAddressesList[i] = strings.Trim(addr, "\"")
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
	concurrentSubmissionCountUpdates, concurrentSubmissionCountUpdatesErr := strconv.Atoi(getEnv("CONCURRENT_SUBMISSION_COUNT_UPDATES", "10"))
	if concurrentSubmissionCountUpdatesErr != nil {
		log.Printf("Invalid CONCURRENT_SUBMISSION_COUNT_UPDATES value, using default of 10: %v", concurrentSubmissionCountUpdatesErr)
		concurrentSubmissionCountUpdates = 10
	}

	config := Settings{
		// RPC Helper Configuration
		RPCNodes:        rpcNodes,
		ArchiveRPCNodes: archiveRPCNodes,
		MaxRetries:      getEnvInt("RPC_MAX_RETRIES", 3),
		RetryDelayMs:    getEnvInt("RPC_RETRY_DELAY_MS", 500),
		MaxRetryDelayS:  getEnvInt("RPC_MAX_RETRY_DELAY_S", 30),
		RequestTimeoutS: getEnvInt("RPC_REQUEST_TIMEOUT_S", 30),

		// Legacy configuration (keeping for backward compatibility)
		ClientUrl:                        getEnv("PROST_RPC_URL", ""),
		ContractAddress:                  strings.Trim(getEnv("PROTOCOL_STATE_CONTRACT", ""), "\""),
		RedisHost:                        getEnv("REDIS_HOST", ""),
		RedisPort:                        getEnv("REDIS_PORT", ""),
		RedisDB:                          getEnv("REDIS_DB", ""),
		AuthReadToken:                    getEnv("AUTH_READ_TOKEN", ""),
		SlackReportingUrl:                getEnv("SLACK_REPORTING_URL", ""),
		TxRelayerUrl:                     getEnv("TX_RELAYER_URL", ""),
		TxRelayerAuthWriteToken:          getEnv("TX_RELAYER_AUTH_WRITE_TOKEN", ""),
		DataMarketAddresses:              dataMarketAddressesList,
		PeriodicEligibleCountAlerts:      periodicEligibleCountAlerts,
		APIHost:                          getEnv("API_HOST", ""),
		AttestorQueuePushEnabled:         attestorQueuePushEnabled,
		InitCleanupEnabled:               initCleanupEnabled,
		ConcurrentSubmissionCountUpdates: concurrentSubmissionCountUpdates,
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

func getEnvInt(key string, defaultValue int) int {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Failed to parse %s environment variable: %v", key, err)
	}
	return intValue
}

func (s *Settings) ToRPCConfig() *rpchelper.RPCConfig {
	config := &rpchelper.RPCConfig{
		Nodes: func() []rpchelper.NodeConfig {
			var nodes []rpchelper.NodeConfig
			for _, url := range s.RPCNodes {
				nodes = append(nodes, rpchelper.NodeConfig{URL: url})
			}
			return nodes
		}(),
		ArchiveNodes: func() []rpchelper.NodeConfig {
			var nodes []rpchelper.NodeConfig
			for _, url := range s.ArchiveRPCNodes {
				nodes = append(nodes, rpchelper.NodeConfig{URL: url})
			}
			return nodes
		}(),
		MaxRetries:     s.MaxRetries,
		RetryDelay:     time.Duration(s.RetryDelayMs) * time.Millisecond,
		MaxRetryDelay:  time.Duration(s.MaxRetryDelayS) * time.Second,
		RequestTimeout: time.Duration(s.RequestTimeoutS) * time.Second,
	}

	// Configure webhook if SlackReportingUrl is provided
	if s.SlackReportingUrl != "" {
		log.Printf("Configuring webhook alerts with URL: %s", s.SlackReportingUrl)
		config.WebhookConfig = &reporting.WebhookConfig{
			URL:     s.SlackReportingUrl,
			Timeout: 30 * time.Second,
			Retries: 3,
		}
	} else {
		log.Printf("No webhook URL configured - SLACK_REPORTING_URL environment variable not set")
	}

	return config
}
