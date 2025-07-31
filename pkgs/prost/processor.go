package prost

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs"
	"submission-sequencer-collector/pkgs/clients"
	"submission-sequencer-collector/pkgs/redis"
	"sync"
	"time"

	// External dependencies
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

// Timeout Hierarchy:
//
// Block Processing Loop (continuous)
// ├── blockFetchTimeout (5s)
// │   └── HeaderByNumber
// │
// ├── blockProcessTimeout (30s)
// │   ├── fetchBlock
// │   │   └── BlockByNumber with retries
// │   │
// │   ├── ProcessEvents (30s)
// │   │   ├── FilterLogs with retries
// │   │   ├── handleEpochReleasedEvent
// │   │   └── handleSnapshotBatchSubmittedEvent
// │   │
// │   ├── processEpochDeadlinesForDataMarkets (120s)
// │   │   └── For each data market:
// │   │       └── checkAndTriggerBatchPreparation (90s)
// │   │           └── triggerBatchPreparation
// │   │
// │   └── Redis block hash storage (5s)

// Timeout durations for various operations
var (
	blockFetchTimeout      time.Duration
	eventProcessingTimeout time.Duration
	batchProcessingTimeout time.Duration
	contractQueryTimeout   time.Duration
)

var windowManager *WindowManager

// Note: All type definitions have been moved to types.go in the same package (prost).
// The following types are imported implicitly:
// - EpochMarkerDetails
// - BatchDetails
// - SubmissionDetails
// - DayTransitionEpochInfo
// - RelayerRequestBody

func triggerBatchPreparation(ctx context.Context, dataMarketAddress string, epochID *big.Int, startBlockNum, endBlockNum int64) error {
	blockCount := endBlockNum - startBlockNum + 1
	log.Infof("🚀 Starting batch preparation for epoch %s, data market %s, processing %d blocks from %d to %d",
		epochID.String(), dataMarketAddress, blockCount, startBlockNum, endBlockNum)

	// Collect all headers first
	headers := make([]string, 0, blockCount)
	for blockNum := startBlockNum; blockNum <= endBlockNum; blockNum++ {
		blockHash, err := redis.Get(ctx, redis.BlockHashByNumber(blockNum))
		if err != nil {
			log.Errorf("❌ Failed to fetch block hash for block %d: %v", blockNum, err)
			continue
		}
		headers = append(headers, blockHash)
	}

	// Get all submission keys for all headers at once
	submissionCtx, cancel := context.WithTimeout(ctx, eventProcessingTimeout)
	submissionKeys, err := getValidSubmissionKeys(submissionCtx, epochID.Uint64(), headers, dataMarketAddress)
	cancel()
	if err != nil {
		log.Errorf("❌ Failed to fetch submission keys for epoch %s: %v", epochID.String(), err)
		return fmt.Errorf("failed to fetch submission keys: %w", err)
	}
	log.Infof("🔑 Retrieved %d submission keys for epoch %s", len(submissionKeys), epochID.String())

	// Update submission count once with all keys
	if err := UpdateSlotSubmissionCount(ctx, epochID, dataMarketAddress, submissionKeys); err != nil {
		log.Errorf("❌ Failed to update submission counts: %v", err)
		return fmt.Errorf("failed to update submission counts: %w", err)
	}

	// Create project map once
	projectMap := constructProjectMap(submissionKeys)
	batches := arrangeSubmissionKeysInBatches(projectMap)
	log.Infof("📊 Created %d batches from %d projects", len(batches), len(projectMap))

	// Store batch count in Redis
	err = redis.SetWithExpiration(ctx, redis.GetBatchCountKey(dataMarketAddress, epochID.String()),
		strconv.Itoa(len(batches)), 24*time.Hour)
	if err != nil {
		log.Errorf("❌ Failed to store batch count: %v", err)
		return fmt.Errorf("failed to store batch count: %w", err)
	}

	// Send batch size once
	if err = SendBatchSizeToRelayer(dataMarketAddress, epochID, len(batches)); err != nil {
		log.Errorf("🚨 Failed to send batch size to relayer: %v", err)
		return fmt.Errorf("failed to send batch size: %w", err)
	}
	log.Infof("📨 Batch size %d sent successfully for epoch %s in data market %s", len(batches), epochID.String(), dataMarketAddress)

	// Process batches concurrently
	var wg sync.WaitGroup
	for i, batch := range batches {
		i, batch := i, batch
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Create submission details
			details := SubmissionDetails{
				EpochID:           epochID,
				BatchID:           i + 1,
				Batch:             batch,
				DataMarketAddress: dataMarketAddress,
			}

			jsonData, err := json.Marshal(details)
			if err != nil {
				log.Errorf("❌ Failed to marshal submission details for batch %d: %v", i+1, err)
				return
			}

			// Push to Redis
			err = redis.LPush(ctx, "finalizerQueue", jsonData).Err()
			if err != nil {
				log.Errorf("❌ Failed to push batch %d to finalizer queue: %v", i+1, err)
				return
			}

			// Store batch details
			batchJSON, err := json.Marshal(batch)
			if err != nil {
				log.Errorf("❌ Failed to marshal batch %d details: %v", i+1, err)
				return
			}

			err = redis.StoreBatchDetails(ctx, dataMarketAddress, epochID.String(),
				strconv.Itoa(i+1), string(batchJSON))
			if err != nil {
				log.Errorf("❌ Failed to store batch %d details: %v", i+1, err)
				return
			}

			log.Infof("✅ Processed batch %d successfully", i+1)
		}()
	}

	wg.Wait()
	log.Infof("✅ Completed all batch processing for epoch %s in data market %s", epochID.String(), dataMarketAddress)
	return nil
}

// Calculate and update total submission count for a data market address on protocol state
func UpdateSlotSubmissionCount(ctx context.Context, epochID *big.Int, dataMarketAddress string, submissionKeys []string) error {
	// Fetch the current day
	currentDay, err := FetchCurrentDay(ctx, common.HexToAddress(dataMarketAddress))
	log.Infof("Current day for data market %s: %s", dataMarketAddress, currentDay.String())
	if err != nil {
		log.Errorf("Failed to fetch current day for data market %s: %v", dataMarketAddress, err)
		return err
	}

	if epochID.Int64()%config.SettingsObj.RewardsUpdateEpochInterval == 0 {
		// Send eligible nodes count to the relayer if the periodic eligible count alerts are set to true
		if config.SettingsObj.PeriodicEligibleCountAlerts {
			// Fetch the slotIDs whose eligible submissions are recorded for the current day
			eligibleNodesByDayKeys := redis.EligibleNodesByDayKey(dataMarketAddress, currentDay.String())
			slotIDs := redis.GetSetKeys(ctx, eligibleNodesByDayKeys)

			// Construct the message with eligible node count and slot IDs
			alertMsg := fmt.Sprintf("🔔 Epoch %d: Eligible node count for data market %s on day %s: %d\nSlot IDs: %v", epochID, dataMarketAddress, currentDay.String(), len(slotIDs), slotIDs)

			// Send the alert
			clients.SendFailureNotification(pkgs.SendEligibleNodesCount, alertMsg, time.Now().String(), "High")
			log.Info(alertMsg)
		}

		// Send the updateRewards request to the relayer, including the count of eligible nodes for the specified buffer days period
		for day := 1; day <= int(math.Min(float64(config.SettingsObj.PastDaysBuffer), float64(currentDay.Int64()))); day++ {
			// Calculate the day to check
			dayToCheck := new(big.Int).Sub(currentDay, big.NewInt(int64(day)))

			// Skip processing if day is less than equal to zero
			if dayToCheck.Int64() <= 0 {
				continue
			}

			// Initialize retry attempts
			retryCount := 0
			for retryCount < config.SettingsObj.RetryLimits {
				// Fetch the contract instance for the data market address
				instance := DataMarketInstances[dataMarketAddress]

				// Fetch eligible node count from data market contract
				count, err := instance.EligibleNodesForDay(&bind.CallOpts{Context: ctx}, dayToCheck)

				// If count is non-zero and no error, break the retry loop
				if err == nil && count != nil && count.Uint64() > 0 {
					log.Infof("✅ Contract Query successful: Eligible node count for data market %s on day %s: %d", dataMarketAddress, dayToCheck.String(), count.Uint64())
					break
				}

				// Log error if contract call failed
				if err != nil {
					log.Errorf("Failed to fetch eligible node count from contract for data market %s on day %s: %v", dataMarketAddress, dayToCheck.String(), err)
				}

				// Calculate the difference between currentDay and dayToCheck
				dayDifference := new(big.Int).Sub(currentDay, dayToCheck)

				// Skip cached count and recalculation when the day has rolled over and epochID is within the buffer range
				// Instead of using passed ctx for Redis ops
				epochsInADay, err := redis.GetEpochsInADay(ctx, dataMarketAddress)
				if err != nil {
					log.Errorf("Failed to fetch epochs in a day for data market %s: %v", dataMarketAddress, err)
					return err
				}

				if dayDifference.Int64() == 1 && int(epochID.Int64())%int(epochsInADay.Int64()) <= BufferEpochs {
					log.Infof("Skipping cached count and recalculation for data market %s on day %s due to epochID %s being in buffer range", dataMarketAddress, dayToCheck.String(), epochID.String())
					break
				}

				// Fetch the current day value from Redis
				cachedCount, err := redis.GetSetCardinality(ctx, redis.EligibleNodesByDayKey(dataMarketAddress, dayToCheck.String()))
				if err != nil {
					errorMsg := fmt.Sprintf("❌ Error fetching cached eligible node count for data market %s on day %s: %v", dataMarketAddress, dayToCheck.String(), err)
					clients.SendFailureNotification(pkgs.SendEligibleNodesCount, errorMsg, time.Now().String(), "Medium")
					log.Error(errorMsg)
					return err
				}

				if cachedCount > 0 {
					log.Infof("Cached eligible node count found for data market %s on day %s: %d", dataMarketAddress, dayToCheck.String(), cachedCount)

					// Attempt to update using cached value
					if err = SendUpdateRewardsToRelayer(ctx, dataMarketAddress, []*big.Int{}, []*big.Int{}, dayToCheck.String(), cachedCount); err != nil {
						errorMsg := fmt.Sprintf("🚨 Failed to send rewards update for data market %s on day %s using cached count: %v", dataMarketAddress, dayToCheck.String(), err)
						clients.SendFailureNotification(pkgs.SendUpdateRewardsToRelayer, errorMsg, time.Now().String(), "High")
						log.Error(errorMsg)
						return err
					}

					successMsg := fmt.Sprintf("✅ Successfully updated rewards using cached count: Eligible node count for data market %s on day %s: %d", dataMarketAddress, dayToCheck.String(), cachedCount)
					clients.SendFailureNotification(pkgs.SendEligibleNodesCount, successMsg, time.Now().String(), "High")
					log.Info(successMsg)

					break
				} else {
					// Check if a zero count update has already been sent
					sent, err := redis.GetBooleanValue(ctx, redis.ZeroCountUpdateKey(dataMarketAddress, dayToCheck.String()))
					if err != nil {
						log.Errorf("Error checking zero count update status for data market %s on day %s: %v", dataMarketAddress, dayToCheck.String(), err)
						return err
					}
					if sent {
						log.Infof("Skipping zero count update for data market %s on day %s as it has already been sent", dataMarketAddress, dayToCheck.String())
						break
					}

					// Fetch eligible slots using the helper function instead of contract calls
					eligibleNodesCount, eligibleSlotIDs := fetchEligibleSlotIDs(ctx, dataMarketAddress, dayToCheck.String())

					// Prepare arrays for the relayer
					slotIDs := make([]*big.Int, 0)
					submissionsList := make([]*big.Int, 0)

					// Get submission counts from Redis for eligible slots
					for _, slotID := range eligibleSlotIDs {
						slotIDBigInt, _ := new(big.Int).SetString(slotID, 10)

						key := redis.EligibleSlotSubmissionKey(dataMarketAddress, slotID, dayToCheck.String())
						slotSubmissionCount, err := redis.Get(ctx, key)

						if err != nil {
							log.Errorf("Failed to fetch eligible submission count for slotID %s, data market %s on day %s: %v",
								slotID, dataMarketAddress, dayToCheck.String(), err)
							continue
						}

						if slotSubmissionCount != "" {
							submissionCountBigInt, ok := new(big.Int).SetString(slotSubmissionCount, 10)
							if !ok {
								log.Errorf("Failed to convert slot submission count %s to big.Int", slotSubmissionCount)
								continue
							}

							slotIDs = append(slotIDs, slotIDBigInt)
							submissionsList = append(submissionsList, submissionCountBigInt)
						}
					}

					log.Infof("Recalculated eligible nodes count for data market %s on day %s: %d", dataMarketAddress, dayToCheck.String(), eligibleNodesCount)

					// Check if recalculated eligible nodes count is zero
					if eligibleNodesCount == 0 {
						log.Infof("Eligible node count is zero for data market %s on day %s", dataMarketAddress, dayToCheck.String())

						if err := redis.SetBooleanValue(ctx, redis.ZeroCountUpdateKey(dataMarketAddress, dayToCheck.String()), true, 24*time.Hour); err != nil {
							log.Errorf("Error marking zero count update as sent for data market %s on day %s: %v", dataMarketAddress, dayToCheck.String(), err)
							return err
						}

						log.Infof("Marked zero count update as sent for data market %s on day %s. Skipping update to relayer.", dataMarketAddress, dayToCheck.String())
						break
					}

					// Send actual submission counts to relayer instead of empty arrays
					if err = SendUpdateRewardsToRelayer(ctx, dataMarketAddress, slotIDs, submissionsList, dayToCheck.String(), eligibleNodesCount); err != nil {
						errorMsg := fmt.Sprintf("🚨 Failed to send rewards update for data market %s on day %s using recalculated count: %v", dataMarketAddress, dayToCheck.String(), err)
						clients.SendFailureNotification(pkgs.SendUpdateRewardsToRelayer, errorMsg, time.Now().String(), "High")
						log.Error(errorMsg)
						retryCount++
						continue
					}

					successMsg := fmt.Sprintf("✅ Successfully updated rewards using recalculated count: Eligible node count for data market %s on day %s: %d", dataMarketAddress, dayToCheck.String(), eligibleNodesCount)
					clients.SendFailureNotification(pkgs.SendEligibleNodesCount, successMsg, time.Now().String(), "High")
					log.Info(successMsg)

					break
				}
			}
		}
	}

	// Process day transitions and store corresponding epoch marker details
	if err := handleDayTransition(ctx, dataMarketAddress, currentDay, epochID); err != nil {
		log.Errorf("Error handling day transition for data market %s: %v", dataMarketAddress, err)
		return err
	}

	// Verify and trigger updateRewards to relayer when the current epoch matches the buffer epoch for any data market
	var wg sync.WaitGroup
	errChan := make(chan error, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := sendFinalRewards(epochID); err != nil {
			select {
			case errChan <- err:
			default:
				log.Errorf("Error sending final rewards for epoch %s: %v", epochID, err)
			}
		}
	}()

	// Fetch day size for the specified data market address from Redis
	daySize, err := redis.GetDaySize(ctx, dataMarketAddress)
	if err != nil {
		log.Errorf("Failed to fetch day size for data market %s: %v", dataMarketAddress, err)
		return err
	}

	// Fetch epochs in a day for the specified data market address from Redis
	epochsInADay, err := redis.GetEpochsInADay(ctx, dataMarketAddress)
	if err != nil {
		log.Errorf("Failed to fetch epochs in a day for data market %s: %v", dataMarketAddress, err)
		return err
	}

	// Calculate expiration time
	expirationTime := getExpirationTime(epochID.Int64(), daySize.Int64(), epochsInADay.Int64())

	// Set the current day in Redis with the calculated expiration duration
	if err := redis.SetWithExpiration(ctx, redis.GetCurrentDayKey(dataMarketAddress), currentDay.String(), time.Until(expirationTime)); err != nil {
		return fmt.Errorf("failed to cache day value for data market %s in Redis: %v", dataMarketAddress, err)
	}

	// Set the last known day in Redis (for detecting day transition)
	if err := redis.Set(ctx, redis.LastKnownDay(dataMarketAddress), currentDay.String()); err != nil {
		return fmt.Errorf("failed to cache last known day value for data market %s in Redis: %v", dataMarketAddress, err)
	}

	// Increment the slot submissions count for a data market in Redis
	for _, submissionKey := range submissionKeys {
		parts := strings.Split(submissionKey, ".")
		slotID := parts[3]

		count, err := redis.Incr(ctx, redis.SlotSubmissionKey(dataMarketAddress, slotID, currentDay.String()))
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to increment submission count for slotID %s, epoch %s in data market %s: %s", slotID, epochID, dataMarketAddress, err.Error())
			clients.SendFailureNotification(pkgs.UpdateSlotSubmissionCount, errorMsg, time.Now().String(), "High")
			log.Error(errorMsg)
			return err
		}

		log.Infof("📈 Slot submission count updated successfully for slotID %s, epoch %s in data market %s: %d", slotID, epochID, dataMarketAddress, count)
	}

	return nil
}

func handleDayTransition(ctx context.Context, dataMarketAddress string, currentDay, epochID *big.Int) error {
	// Fetch the last known day value from Redis
	lastKnownDay, err := redis.Get(ctx, redis.LastKnownDay(dataMarketAddress))
	log.Infof("Last known day for data market %s: %s", dataMarketAddress, lastKnownDay)
	if err != nil {
		log.Errorf("Error fetching last known day value for data market %s from Redis: %v", dataMarketAddress, err)
		return err
	}

	// Check for day transition
	if lastKnownDay != "" && lastKnownDay != currentDay.String() {
		log.Infof("Day transition detected for data market %s: %s -> %s", dataMarketAddress, lastKnownDay, currentDay.String())

		// Prepare the day transition epoch marker details
		dayTransitionEpochDetails := DayTransitionEpochInfo{
			LastKnownDay: lastKnownDay,
			CurrentEpoch: epochID.Int64(),
			BufferEpoch:  epochID.Int64() + int64(BufferEpochs),
		}

		dayTransitionEpochDetailsJSON, err := json.Marshal(dayTransitionEpochDetails)
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to marshal day transtion epoch details for epochID %s, data market %s: %s", epochID.String(), dataMarketAddress, err.Error())
			clients.SendFailureNotification(pkgs.HandleDayTransition, errorMsg, time.Now().String(), "High")
			log.Error(errorMsg)
			return err
		}

		// Store the day transition epoch details in Redis
		if err := redis.StoreDayTransitionEpochDetails(ctx, dataMarketAddress, epochID.String(), string(dayTransitionEpochDetailsJSON)); err != nil {
			errorMsg := fmt.Sprintf("Failed to store day transition epoch marker details for epochID %s, data market %s in Redis: %s", epochID.String(), dataMarketAddress, err.Error())
			clients.SendFailureNotification(pkgs.HandleDayTransition, errorMsg, time.Now().String(), "High")
			log.Error(errorMsg)
		}

		log.Infof("✅ Successfully stored day transition epoch marker details for epochID %s, data market %s in Redis", epochID.String(), dataMarketAddress)
	}

	return nil
}

func sendRewardUpdates(ctx context.Context, dataMarketAddress, epochID string) error {
	// Fetch the current day
	currentDay, err := FetchCurrentDay(ctx, common.HexToAddress(dataMarketAddress))
	if err != nil {
		log.Errorf("Failed to fetch current day for data market %s: %v", dataMarketAddress, err)
		return err
	}

	// Prepare data for relayer
	slotIDs := make([]*big.Int, 0)
	submissionsList := make([]*big.Int, 0)

	cachedTotalNodesCount, err := redis.Get(ctx, redis.TotalNodesCountKey())
	if err != nil {
		log.Errorf("❌ Failed to fetch total nodes count for data market %s: %v", dataMarketAddress, err)
		return err
	}

	totalNodesCount, ok := new(big.Int).SetString(cachedTotalNodesCount, 10)
	if !ok {
		log.Errorf("❌ Failed to convert total nodes count %s to big.Int", cachedTotalNodesCount)
		return err
	}

	for slotID := int64(1); slotID <= totalNodesCount.Int64(); slotID++ {
		// Get the eligible submission count from Redis
		key := redis.EligibleSlotSubmissionKey(dataMarketAddress, big.NewInt(slotID).String(), currentDay.String())
		slotSubmissionCount, err := redis.Get(ctx, key)
		if err != nil {
			log.Errorf("Failed to fetch eligible submission count for slotID %d, epoch %s within data market %s: %v", slotID, epochID, dataMarketAddress, err)
			continue
		}

		// Check that submission count value is not empty
		if slotSubmissionCount != "" {
			submissionCountBigInt, ok := new(big.Int).SetString(slotSubmissionCount, 10)
			if !ok {
				log.Errorf("Failed to convert slot submission count %s to big.Int", slotSubmissionCount)
				continue
			}

			// Add the slotID and submission count to the respective arrays
			slotIDs = append(slotIDs, big.NewInt(slotID))
			submissionsList = append(submissionsList, submissionCountBigInt)
		}
	}

	// Use asyncBatchArrays instead of batchArrays
	asyncBatchArrays(dataMarketAddress, currentDay.String(), slotIDs, submissionsList, 0)

	return nil
}

func sendFinalRewards(currentEpoch *big.Int) error {
	log.Infof("🔍 Initiating day transition check for current epoch: %s", currentEpoch.String())
	var wg sync.WaitGroup

	// Create error channel to collect errors from goroutines
	errChan := make(chan error, len(config.SettingsObj.DataMarketAddresses))

	// Fetch and process day transition epoch markers set concurrently for each data market address
	for _, dataMarketAddress := range config.SettingsObj.DataMarketAddresses {
		dataMarketAddress := dataMarketAddress // Capture for goroutine
		wg.Add(1)

		// Process each data market concurrently
		go func() {
			defer wg.Done()

			// Create independent context for this market
			marketCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
			defer cancel()

			// Process the data market
			epochMarkerKeys, err := redis.RedisClient.SMembers(marketCtx, redis.DayRolloverEpochMarkerSet(dataMarketAddress)).Result()
			if err != nil {
				log.Errorf("Failed to fetch day transition epoch markers from Redis for data market %s: %s", dataMarketAddress, err)
				errChan <- fmt.Errorf("failed to fetch day transition epoch markers: %w", err)
				return
			}

			log.Infof("✅ Fetched %d day transition epoch marker keys for data market %s: %v", len(epochMarkerKeys), dataMarketAddress, epochMarkerKeys)

			// Process each day transition epoch marker key for this data market address
			for _, epochMarkerKey := range epochMarkerKeys {
				select {
				case <-marketCtx.Done():
					errChan <- marketCtx.Err()
					return
				default:
					// Retrieve the day transition epoch marker details from Redis
					dayTransitionMarkerDetailsJSON, err := redis.RedisClient.Get(marketCtx, redis.DayRolloverEpochMarkerDetails(dataMarketAddress, epochMarkerKey)).Result()
					if err != nil {
						log.Errorf("Failed to fetch day transition epoch marker details from Redis for key %s: %s", epochMarkerKey, err)
						continue
					}

					var dayTransitionMarkerInfo DayTransitionEpochInfo
					if err := json.Unmarshal([]byte(dayTransitionMarkerDetailsJSON), &dayTransitionMarkerInfo); err != nil {
						log.Errorf("Failed to unmarshal day transition epoch marker details for key %s: %s", epochMarkerKey, err)
						continue
					}

					log.Debugf("📊 Day transition epoch marker details for key %s: %+v", epochMarkerKey, dayTransitionMarkerInfo)

					// Check if the current epoch matches the buffer epoch
					if currentEpoch.Int64() == dayTransitionMarkerInfo.BufferEpoch {
						lastKnownDay := dayTransitionMarkerInfo.LastKnownDay

						// Fetch the eligible nodes count and slotIDs for the last known day (prev day)
						eligibleNodesCount, eligibleSlotIDs := fetchEligibleSlotIDs(marketCtx, dataMarketAddress, lastKnownDay)

						log.Infof("✅ Successfully fetched eligible nodes count for data market %s on day %s: %d", dataMarketAddress, lastKnownDay, eligibleNodesCount)

						// Prepare data for relayer
						slotIDs := make([]*big.Int, 0)
						submissionsList := make([]*big.Int, 0)

						for _, slotID := range eligibleSlotIDs {
							slotIDBigInt, _ := new(big.Int).SetString(slotID, 10)

							key := redis.EligibleSlotSubmissionKey(dataMarketAddress, slotID, lastKnownDay)
							slotSubmissionCount, err := redis.Get(marketCtx, key)
							if err != nil {
								log.Errorf("Failed to fetch eligible submission count for slotID %s, data market %s on day %s: %v",
									slotID, dataMarketAddress, lastKnownDay, err)
								continue
							}

							if slotSubmissionCount != "" {
								submissionCountBigInt, ok := new(big.Int).SetString(slotSubmissionCount, 10)
								if !ok {
									log.Errorf("Failed to convert slot submission count %s to big.Int", slotSubmissionCount)
									continue
								}

								slotIDs = append(slotIDs, slotIDBigInt)
								submissionsList = append(submissionsList, submissionCountBigInt)
							}
						}

						// Call non-blocking batch processing
						asyncBatchArrays(dataMarketAddress, lastKnownDay, slotIDs, submissionsList, eligibleNodesCount)

						// Remove the epochID and its day transition details from Redis
						epochID := new(big.Int).Sub(currentEpoch, big.NewInt(int64(BufferEpochs)))
						if err := redis.RemoveDayTransitionEpochFromRedis(marketCtx, dataMarketAddress, epochID.String()); err != nil {
							log.Errorf("Error removing day transition epoch %s data from Redis for data market %s: %v", epochID.String(), dataMarketAddress, err)
							errChan <- fmt.Errorf("failed to remove day transition epoch data: %w", err)
						}

						log.Infof("🧹 Successfully removed day transition epoch %s data from Redis for data market %s", epochID.String(), dataMarketAddress)

						// Send alert message about eligible nodes count update
						alertMsg := fmt.Sprintf("🔔 Day Transition Update: Eligible nodes count for data market %s has been updated for day %s: %d", dataMarketAddress, lastKnownDay, eligibleNodesCount)
						clients.SendFailureNotification(pkgs.SendEligibleNodesCount, alertMsg, time.Now().String(), "High")
						log.Info(alertMsg)
					}
				}
			}
			log.Infof("Completed processing for data market %s at epoch: %d", dataMarketAddress, currentEpoch.Uint64())
		}()
	}

	// Wait for all data market goroutines to finish with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
		close(errChan)
	}()

	select {
	case <-done:
		// Collect any errors
		var errs []error
		for err := range errChan {
			if err != nil {
				errs = append(errs, err)
			}
		}

		if len(errs) > 0 {
			return fmt.Errorf("multiple errors occurred during send rewards for epoch %s: %v", currentEpoch.String(), errs)
		}
		return nil
	case <-time.After(3 * time.Minute):
		log.Warnf("⚠️ sendFinalRewards timed out waiting for markets to process for epoch %s", currentEpoch.String())
		return fmt.Errorf("timed out processing day transition for epoch %s", currentEpoch.String())
	}
}

// this batches reward updates to the relayer
func asyncBatchArrays(dataMarketAddress, currentDay string, slotIDs, submissionsList []*big.Int, eligibleNodesCount int) {
	// Create a semaphore to limit concurrent batches
	batchSemaphore := make(chan struct{}, 5) // Limit to 5 concurrent batches
	maxBatches := 100                        // Maximum number of batches to process

	// Create a self-managed goroutine that doesn't block the caller
	go func() {
		// Create independent context with timeout
		batchCtx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer cancel()

		// Fetch the batch size from config
		batchSize := config.SettingsObj.RewardsUpdateBatchSize
		var wg sync.WaitGroup
		errChan := make(chan error, len(slotIDs)/batchSize+1)

		// Process the data in batches
		for start := 0; start < len(slotIDs) && start/batchSize < maxBatches; start += batchSize {
			end := start + batchSize
			if end > len(slotIDs) {
				end = len(slotIDs)
			}

			// Create new slices for each batch - these are small and short-lived
			batchSlotIDs := slotIDs[start:end]
			batchSubmissions := submissionsList[start:end]

			// Acquire semaphore with timeout
			select {
			case batchSemaphore <- struct{}{}:
				// Got permission to proceed
			case <-batchCtx.Done():
				log.Warnf("⚠️ Context cancelled while waiting for batch semaphore for market %s day %s",
					dataMarketAddress, currentDay)
				return
			case <-time.After(30 * time.Second):
				log.Warnf("⚠️ Timeout waiting for batch semaphore for market %s day %s",
					dataMarketAddress, currentDay)
				continue
			}

			wg.Add(1)
			go func(start, end int, batchSlotIDs, batchSubmissions []*big.Int) {
				defer wg.Done()
				defer func() { <-batchSemaphore }() // Release semaphore when done

				// Request context as child of batch context
				reqCtx, reqCancel := context.WithTimeout(batchCtx, 30*time.Second)
				defer reqCancel()

				if err := SendUpdateRewardsToRelayer(reqCtx, dataMarketAddress, batchSlotIDs, batchSubmissions, currentDay, eligibleNodesCount); err != nil {
					errorMsg := fmt.Sprintf("🚨 Relayer batch error in sending rewards update batch %d-%d for data market %s on day %s: %v", start, end, dataMarketAddress, currentDay, err)
					clients.SendFailureNotification(pkgs.SendUpdateRewardsToRelayer, errorMsg, time.Now().String(), "High")
					log.Error(errorMsg)
					select {
					case errChan <- err:
					case <-reqCtx.Done():
					}
				} else {
					log.Infof("✅ Successfully sent rewards update batch %d-%d for data market %s on day %s",
						start, end, dataMarketAddress, currentDay)
				}
			}(start, end, batchSlotIDs, batchSubmissions)

			// Small delay between batch starts to prevent overwhelming the relayer
			select {
			case <-time.After(100 * time.Millisecond):
			case <-batchCtx.Done():
				return
			}
		}

		// Wait for all batches to complete or timeout
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
			close(errChan)
		}()

		select {
		case <-done:
			// Check for any errors
			var errors []error
			for err := range errChan {
				errors = append(errors, err)
			}

			if len(errors) > 0 {
				log.Errorf("❌ %d rewards update batch processing errors for data market %s on day %s",
					len(errors), dataMarketAddress, currentDay)
			} else {
				log.Infof("✅ Successfully sent all %d rewards update batches to relayer for data market %s on day %s",
					(len(slotIDs)+batchSize-1)/batchSize, dataMarketAddress, currentDay)
			}
		case <-time.After(2*time.Minute + 30*time.Second):
			log.Warnf("⚠️ Timed out waiting for rewards update batches to complete for market %s day %s",
				dataMarketAddress, currentDay)
		}
	}()
}

func InitializeTimeouts() {
	blockFetchTimeout = time.Second * time.Duration(config.SettingsObj.BlockFetchTimeout)
	eventProcessingTimeout = time.Second * time.Duration(config.SettingsObj.EventProcessingTimeout)
	batchProcessingTimeout = time.Second * time.Duration(config.SettingsObj.BatchProcessingTimeout)
	contractQueryTimeout = time.Second * time.Duration(config.SettingsObj.ContractQueryTimeout)
}

func InitializeSubmissionWindowProcessor() {
	windowManager = NewWindowManager()
}
