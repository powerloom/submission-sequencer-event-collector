package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs/clients"
	"submission-sequencer-collector/pkgs/prost"
	"submission-sequencer-collector/pkgs/redis"
	"submission-sequencer-collector/pkgs/service"
	"submission-sequencer-collector/pkgs/utils"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Initiate logger
	utils.InitLogger()

	// Load the config object
	config.LoadConfig()

	// Initialize reporting service
	clients.InitializeReportingClient(config.SettingsObj.SlackReportingUrl, 5*time.Second)

	// Initialize tx relayer service
	clients.InitializeTxClient(config.SettingsObj.TxRelayerUrl, time.Duration(config.SettingsObj.HttpTimeout)*time.Second)

	// Setup redis
	redis.RedisClient = redis.NewRedisClient()

	// Create root context before any client configuration
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := prost.ConfigureClient(ctx); err != nil {
		log.Fatal(err)
	}

	// Set up the RPC client, contract, and ABI instance
	if err := prost.ConfigureContractInstance(ctx); err != nil {
		log.Fatal(err)
	}
	prost.ConfigureABI()

	// Load the state variables from the protocol state contract
	if err := prost.LoadContractStateVariables(ctx); err != nil {
		log.Fatal(err)
	}

	// Create a root context that can be cancelled
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup

	// Run initial cleanup for all data markets
	for _, dataMarketAddress := range config.SettingsObj.DataMarketAddresses {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			if err := prost.CleanupSubmissionSet(ctx, addr); err != nil {
				log.Printf("Initial cleanup failed for %s: %v", addr, err)
			}
		}(dataMarketAddress)
	}

	wg.Add(1)
	go service.StartApiServer() // Start API Server

	wg.Add(1)
	go prost.StartFetchingBlocks(ctx) // Pass the context

	// Add signal handling for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	cancel() // Signal all goroutines to shut down
	wg.Wait()
}
