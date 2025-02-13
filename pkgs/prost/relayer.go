package prost

import (
	"context"
	"math/big"
	"submission-sequencer-collector/pkgs/clients"
	"time"

	"github.com/cenkalti/backoff"
	log "github.com/sirupsen/logrus"
)

func SendBatchSizeToRelayer(dataMarketAddress string, epochID *big.Int, batchSize int) error {
	// Define the operation that will be retried
	operation := func() error {
		// Attempt to submit the batch size
		err := clients.SendSubmissionBatchSize(dataMarketAddress, epochID, batchSize)
		if err != nil {
			log.Errorf("Error sending submission batch size for epoch %s, data market %s: %v", epochID, dataMarketAddress, err)
			return err // Return error to trigger retry
		}

		log.Infof("Successfully submitted batch size for epoch %s, data market %s", epochID, dataMarketAddress)
		return nil // Successful submission, no need for further retries
	}

	// Customize the backoff configuration
	backoffConfig := backoff.NewExponentialBackOff()
	backoffConfig.InitialInterval = 1 * time.Second // Start with a 1-second delay
	backoffConfig.Multiplier = 1.5                  // Increase interval by 1.5x after each retry
	backoffConfig.MaxInterval = 4 * time.Second     // Set max interval between retries
	backoffConfig.MaxElapsedTime = 10 * time.Second // Retry for a maximum of 10 seconds

	// Limit retries to 3 times within 10 seconds
	if err := backoff.Retry(operation, backoff.WithMaxRetries(backoffConfig, 3)); err != nil {
		log.Errorf("Failed to submit batch size for epoch %s, data market %s after multiple retries: %v", epochID, dataMarketAddress, err)
		return err
	}

	return nil
}

func SendUpdateRewardsToRelayer(ctx context.Context, dataMarketAddress string, slotIDs, submissionsList []*big.Int, day string, eligibleNodes int) error {
	// Define the operation that will be retried
	operation := func() error {
		// Pass ctx to SendUpdateRewardsRequest
		err := clients.SendUpdateRewardsRequest(ctx, dataMarketAddress, slotIDs, submissionsList, day, eligibleNodes)
		if err != nil {
			log.Errorf("Error sending final updateRewards request for data market %s on day %s: %v. Retrying...", dataMarketAddress, day, err)
			return err // Return error to trigger retry
		}

		log.Infof("📤 Successfully sent final updateRewards request for data market %s on day %s", dataMarketAddress, day)
		return nil // Successful submission, no need for further retries
	}

	// Customize the backoff configuration
	backoffConfig := backoff.NewExponentialBackOff()
	backoffConfig.InitialInterval = 1 * time.Second // Start with a 1-second delay
	backoffConfig.Multiplier = 1.5                  // Increase interval by 1.5x after each retry
	backoffConfig.MaxInterval = 4 * time.Second     // Set max interval between retries
	backoffConfig.MaxElapsedTime = 10 * time.Second // Retry for a maximum of 10 seconds

	// Use context with backoff
	if err := backoff.Retry(operation, backoff.WithContext(backoffConfig, ctx)); err != nil {
		log.Errorf("Failed to send final updateRewards request after retries for data market %s on day %s: %v", dataMarketAddress, day, err)
		return err
	}

	return nil
}
