package prost

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"runtime"
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
	"golang.org/x/sync/errgroup"
)

// Timeout durations for various operations
const (
	eventProcessingTimeout  = 30 * time.Second
	batchPreparationTimeout = 60 * time.Second
	marketProcessingTimeout = 30 * time.Second
	batchProcessingTimeout  = 120 * time.Second
)

var (
	// Memory pools for frequently allocated slices
	headerPool = sync.Pool{
		New: func() any {
			slice := make([]string, 0, 1000) // Typical block range size
			return &slice
		},
	}
	slotIDPool = sync.Pool{
		New: func() any {
			slice := make([]*big.Int, 0, config.SettingsObj.RewardsUpdateBatchSize)
			return &slice
		},
	}
	submissionsPool = sync.Pool{
		New: func() any {
			slice := make([]*big.Int, 0, config.SettingsObj.RewardsUpdateBatchSize)
			return &slice
		},
	}
)

// Note: All type definitions have been moved to types.go in the same package (prost).
// The following types are imported implicitly:
// - EpochMarkerDetails
// - BatchDetails
// - SubmissionDetails
// - DayTransitionEpochInfo
// - RelayerRequestBody

func processMarketData(ctx context.Context, dataMarketAddress string, currentBlockNum int64) error {
	// Panic recovery
	defer func() {
		if r := recover(); r != nil {
			stack := make([]byte, 4096)
			stack = stack[:runtime.Stack(stack, false)]
			errMsg := fmt.Sprintf("Panic in market data processing: %v\n%s", r, stack)
			log.Error(errMsg)
			errAlert := fmt.Errorf("panic in market data processing: %v", r)
			clients.SendFailureNotification(pkgs.ProcessMarketData, errAlert.Error(), time.Now().String(), "High")
		}
	}()

	// Check context before starting
	if err := ctx.Err(); err != nil {
		return err
	}

	// Fetch all the epoch marker keys from Redis for this data market address
	epochMarkerKeys, err := redis.RedisClient.SMembers(ctx, redis.EpochMarkerSet(dataMarketAddress)).Result()
	if err != nil {
		return fmt.Errorf("failed to fetch epoch markers: %w", err)
	}

	log.Infof("Fetched %d epoch marker keys from cache for data market %s: %v", len(epochMarkerKeys), dataMarketAddress, epochMarkerKeys)

	// Create error group for coordinated error handling of epoch processing
	g, epochCtx := errgroup.WithContext(ctx)

	// Process each epoch marker concurrently with panic recovery
	for _, epochMarkerKey := range epochMarkerKeys {
		epochMarkerKey := epochMarkerKey
		g.Go(func() (err error) {
			// Panic recovery for each goroutine
			defer func() {
				if r := recover(); r != nil {
					stack := make([]byte, 4096)
					stack = stack[:runtime.Stack(stack, false)]
					errMsg := fmt.Sprintf("Panic in epoch marker processing: %v\n%s", r, stack)
					log.Error(errMsg)
					clients.SendFailureNotification(pkgs.ProcessMarketData, errMsg, time.Now().String(), "High")
					err = fmt.Errorf("panic in epoch marker processing: %v", r)
				}
			}()

			// Create Redis context with timeout
			redisCtx, cancel := context.WithTimeout(epochCtx, redisOperationTimeout)
			defer cancel()

			// Retrieve the epoch marker details from Redis
			epochMarkerDetailsJSON, err := redis.RedisClient.Get(redisCtx, redis.EpochMarkerDetails(dataMarketAddress, epochMarkerKey)).Result()
			if err != nil {
				errMsg := fmt.Sprintf("Failed to fetch epoch marker details from Redis for key %s: %s", epochMarkerKey, err)
				clients.SendFailureNotification(pkgs.CheckAndTriggerBatchPreparation, errMsg, time.Now().String(), "High")
				log.Error(errMsg)
				return fmt.Errorf("failed to fetch epoch marker details: %w", err)
			}

			var epochMarkerDetails EpochMarkerDetails
			if err := json.Unmarshal([]byte(epochMarkerDetailsJSON), &epochMarkerDetails); err != nil {
				errMsg := fmt.Sprintf("Failed to unmarshal epoch marker details for key %s: %s", epochMarkerKey, err)
				clients.SendFailureNotification(pkgs.CheckAndTriggerBatchPreparation, errMsg, time.Now().String(), "High")
				log.Error(errMsg)
				return fmt.Errorf("failed to unmarshal epoch marker details: %w", err)
			}

			// Check if the current block number matches the submission limit block number for this epoch
			if currentBlockNum == epochMarkerDetails.SubmissionLimitBlockNumber {
				log.Infof("🔄 Initiating batch preparation for epoch %s, data market %s at submission limit block number: %d", epochMarkerKey, dataMarketAddress, currentBlockNum)

				// Convert the epoch ID string to big.Int for further processing
				epochID, ok := big.NewInt(0).SetString(epochMarkerKey, 10)
				if !ok {
					log.Errorf("Failed to convert epochID %s to big.Int for data market %s", epochMarkerKey, dataMarketAddress)
					return fmt.Errorf("failed to convert epochID %s to big.Int", epochMarkerKey)
				}

				// Create batch preparation context
				batchCtx, cancel := context.WithTimeout(epochCtx, batchPreparationTimeout)
				defer cancel()

				select {
				case <-epochCtx.Done():
					return epochCtx.Err()
				case workerPool <- struct{}{}: // Acquire worker with context awareness
					defer func() { <-workerPool }() // Release worker
					if err := triggerBatchPreparation(batchCtx, dataMarketAddress, epochID, epochMarkerDetails.EpochReleaseBlockNumber, currentBlockNum); err != nil {
						return fmt.Errorf("failed to trigger batch preparation: %w", err)
					}
					return nil
				}
			}
			return nil
		})
	}

	// Wait for all epoch processing to complete
	return g.Wait()
}

func triggerBatchPreparation(ctx context.Context, dataMarketAddress string, epochID *big.Int, startBlockNum, endBlockNum int64) error {
	// Initialize headers from pool
	headers := *(headerPool.Get().(*[]string))

	// Declare err here so it's in scope for the deferred function
	var err error

	// Deferred cleanup with panic recovery
	defer func() {
		headerPool.Put(&headers)

		if r := recover(); r != nil {
			stack := make([]byte, 4096)
			stack = stack[:runtime.Stack(stack, false)]
			errMsg := fmt.Sprintf("Panic in batch preparation: %v\n%s", r, stack)
			log.Error(errMsg)
			clients.SendFailureNotification(pkgs.TriggerBatchPreparation, errMsg, time.Now().String(), "High")
			// Convert panic to error
			err = fmt.Errorf("panic in batch preparation: %v", r)
		}
	}()

	// Calculate the total number of blocks in the range
	blockCount := endBlockNum - startBlockNum + 1
	log.Infof("🚀 Starting batch preparation for epoch %s, data market %s, processing %d blocks from %d to %d", epochID.String(), dataMarketAddress, blockCount, startBlockNum, endBlockNum)

	// Create error group for block header processing
	g, gctx := errgroup.WithContext(ctx)

	// Process block headers concurrently
	for blockNum := startBlockNum; blockNum <= endBlockNum; blockNum++ {
		blockNum := blockNum // Capture for closure
		g.Go(func() error {
			select {
			case <-gctx.Done():
				return gctx.Err()
			default:
				// Create Redis context with timeout
				redisCtx, cancel := context.WithTimeout(gctx, redisOperationTimeout)
				defer cancel()

				// Generate the Redis key for the current block number
				blockKey := redis.BlockHashByNumber(blockNum)

				// Fetch the block hash from Redis using the generated key
				blockHashValue, err := redis.Get(redisCtx, blockKey)
				if err != nil {
					errMsg := fmt.Sprintf("Failed to fetch block hash for block %d: %s", blockNum, err.Error())
					clients.SendFailureNotification(pkgs.TriggerBatchPreparation, errMsg, time.Now().String(), "High")
					log.Error(errMsg)
					return fmt.Errorf("failed to fetch block hash: %w", err)
				}

				// Convert the block hash from string to common.Hash type
				blockHash := common.HexToHash(blockHashValue)

				// Add the block hash to the headers slice with mutex protection
				headers = append(headers, blockHash.Hex())
				return nil
			}
		})
	}

	// Wait for all block header processing to complete
	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to process block headers: %w", err)
	}

	log.Infof("📦 Collected %d headers for epoch %s in data market %s", len(headers), epochID.String(), dataMarketAddress)

	// Create submission keys context with timeout
	submissionCtx, cancel := context.WithTimeout(ctx, eventProcessingTimeout)
	defer cancel()

	// Fetch valid submission keys for the epoch
	submissionKeys, err := getValidSubmissionKeys(submissionCtx, epochID.Uint64(), headers, dataMarketAddress)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch submission keys for epoch %s in data market %s: %s", epochID.String(), dataMarketAddress, err.Error())
		clients.SendFailureNotification(pkgs.TriggerBatchPreparation, errMsg, time.Now().String(), "High")
		log.Error(errMsg)
		return fmt.Errorf("failed to fetch submission keys: %w", err)
	}

	log.Infof("🔑 Retrieved %d valid submission keys for epoch %s in data market %s", len(submissionKeys), epochID.String(), dataMarketAddress)

	// Update total submission count with timeout
	updateCtx, updateCancel := context.WithTimeout(ctx, eventProcessingTimeout)
	defer updateCancel()

	if err := UpdateSlotSubmissionCount(updateCtx, epochID, dataMarketAddress, submissionKeys); err != nil {
		log.Errorf("Failed to update slot submission counts for epoch %s in data market %s: %s", epochID, dataMarketAddress, err.Error())
		return fmt.Errorf("failed to update slot submission counts: %w", err)
	}

	// Construct the project map [ProjectID -> SubmissionKeys]
	projectMap := constructProjectMap(submissionKeys)
	log.Infof("📊 Project map created with %d projects for epoch %s in data market %s", len(projectMap), epochID.String(), dataMarketAddress)

	// Arrange the projectMap into batches of submission keys
	batches := arrangeSubmissionKeysInBatches(projectMap)
	log.Infof("🔄 Arranged %d batches of submission keys for epoch %s in data market %s", len(batches), epochID.String(), dataMarketAddress)

	// Store batch count in Redis with timeout
	redisCtx, redisCancel := context.WithTimeout(ctx, redisOperationTimeout)
	defer redisCancel()

	if err := redis.SetWithExpiration(redisCtx, redis.GetBatchCountKey(dataMarketAddress, epochID.String()), strconv.Itoa(len(batches)), 24*time.Hour); err != nil {
		log.Errorf("Failed to set batch count for epoch %s, data market %s in Redis: %s", epochID.String(), dataMarketAddress, err.Error())
		return fmt.Errorf("failed to set batch count: %w", err)
	}

	// Send batch size to relayer
	if err = SendBatchSizeToRelayer(dataMarketAddress, epochID, len(batches)); err != nil {
		errMsg := fmt.Sprintf("🚨 Failed to send submission batch size for epoch %s in data market %s to relayer: %s", epochID.String(), dataMarketAddress, err.Error())
		clients.SendFailureNotification(pkgs.TriggerBatchPreparation, errMsg, time.Now().String(), "High")
		log.Error(errMsg)
		return fmt.Errorf("failed to send batch size to relayer: %w", err)
	}

	log.Infof("📨 Batch size %d sent successfully for epoch %s in data market %s", len(batches), epochID.String(), dataMarketAddress)

	// Create error group for batch processing
	batchGroup, batchCtx := errgroup.WithContext(ctx)

	// Process batches concurrently
	for i, batch := range batches {
		i, batch := i, batch // Capture for closure
		batchGroup.Go(func() error {
			select {
			case <-batchCtx.Done():
				return batchCtx.Err()
			default:
				// Create submission details
				submissionDetails := SubmissionDetails{
					EpochID:           epochID,
					BatchID:           i + 1,
					Batch:             batch,
					DataMarketAddress: dataMarketAddress,
				}

				// Serialize the struct to JSON
				jsonData, err := json.Marshal(submissionDetails)
				if err != nil {
					errMsg := fmt.Sprintf("Serialization failed for submission details of batch %d, epoch %s in data market %s: %v", i+1, epochID.String(), dataMarketAddress, err)
					clients.SendFailureNotification(pkgs.TriggerBatchPreparation, errMsg, time.Now().String(), "High")
					log.Error(errMsg)
					return fmt.Errorf("failed to marshal submission details: %w", err)
				}

				// Create Redis context with timeout for each operation
				redisCtx, cancel := context.WithTimeout(batchCtx, redisOperationTimeout)
				defer cancel()

				// Push the serialized data to Redis
				if err := redis.LPush(redisCtx, "finalizerQueue", jsonData).Err(); err != nil {
					errMsg := fmt.Sprintf("Error pushing submission details of batch %d to Redis for epoch %s in data market %s to finalizer queue in Redis: %v", i+1, epochID.String(), dataMarketAddress, err)
					clients.SendFailureNotification(pkgs.TriggerBatchPreparation, errMsg, time.Now().String(), "High")
					log.Error(errMsg)
					return fmt.Errorf("failed to push submission details to finalizer queue: %w", err)
				}

				// Serialize the batch details to JSON
				batchJSONData, err := json.Marshal(batch)
				if err != nil {
					errMsg := fmt.Sprintf("Serialization failed for batch details of batch %d, epoch %s in data market %s: %v", i+1, epochID.String(), dataMarketAddress, err)
					clients.SendFailureNotification(pkgs.TriggerBatchPreparation, errMsg, time.Now().String(), "High")
					log.Error(errMsg)
					return fmt.Errorf("failed to marshal batch details: %w", err)
				}

				// Convert the batch ID to a big integer
				batchID := big.NewInt(int64(i + 1))

				// Create new Redis context for storing batch details
				storeCtx, storeCancel := context.WithTimeout(batchCtx, redisOperationTimeout)
				defer storeCancel()

				// Store the batch details with a key generated from dataMarketAddress, epochID, and batchID
				if err := redis.StoreBatchDetails(storeCtx, dataMarketAddress, epochID.String(), batchID.String(), string(batchJSONData)); err != nil {
					log.Errorf("Failed to store details for batch %d of epoch %s in data market %s: %v", batchID.Int64(), epochID.String(), dataMarketAddress, err)
					return fmt.Errorf("failed to store batch details: %w", err)
				}

				log.Infof("✅ Batch %d successfully pushed to Redis and stored for epoch %s in data market %s", batchID.Int64(), epochID.String(), dataMarketAddress)
				return nil
			}
		})
	}

	// Wait for all batch processing to complete
	if err := batchGroup.Wait(); err != nil {
		return fmt.Errorf("batch processing failed: %w", err)
	}

	return nil
}

// Calculate and update total submission count for a data market address
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
				var count *big.Int
				if output, err := MustQuery(ctx, func() (*big.Int, error) {
					return instance.EligibleNodesForDay(&bind.CallOpts{}, dayToCheck)
				}); err == nil {
					count = output
				}

				// If count is non-zero, break the retry loop
				if count != nil && count.Uint64() > 0 {
					log.Infof("✅ Contract Query successful: Eligible node count for data market %s on day %s: %d", dataMarketAddress, dayToCheck.String(), count.Uint64())
					break
				}

				// Calculate the difference between currentDay and dayToCheck
				dayDifference := new(big.Int).Sub(currentDay, dayToCheck)

				// Skip cached count and recalculation when the day has rolled over and epochID is within the buffer range
				epochsInADay, err := redis.GetEpochsInADay(ctx, dataMarketAddress)
				if err != nil {
					log.Errorf("Failed to fetch epochs in a day for data market %s: %v", dataMarketAddress, err)
					return err
				}

				if dayDifference.Int64() == 1 && int(epochID.Int64())%int(epochsInADay.Int64()) <= BufferEpochs {
					log.Infof("Skipping cached count and recalculation for data market %s on day %s due to epochID %s being in buffer range", dataMarketAddress, dayToCheck.String(), epochID.String())
					break
				}

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

					// Fallback to recalculation if cached value is not found
					eligibleNodes := 0

					// Fetch daily snapshot quota for the specified data market address from Redis
					dailySnapshotQuota, err := redis.GetDailySnapshotQuota(ctx, dataMarketAddress)
					if err != nil {
						log.Errorf("❌ Failed to fetch daily snapshot quota for data market %s: %v", dataMarketAddress, err)
						return err
					}
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
						var slotSubmissionCount *big.Int
						if output, err := MustQuery(ctx, func() (*big.Int, error) {
							return Instance.SlotSubmissionCount(&bind.CallOpts{}, common.HexToAddress(dataMarketAddress), big.NewInt(int64(slotID)), dayToCheck)
						}); err == nil {
							slotSubmissionCount = output
						}

						if slotSubmissionCount != nil && slotSubmissionCount.Uint64() >= dailySnapshotQuota.Uint64() {
							eligibleNodes++
						}
					}

					log.Infof("Recalculated eligible nodes count for data market %s on day %s: %d", dataMarketAddress, dayToCheck.String(), eligibleNodes)

					// Check if recalculated eligible nodes count is zero
					if eligibleNodes == 0 {
						log.Infof("Eligible node count is zero for data market %s on day %s", dataMarketAddress, dayToCheck.String())

						if err := redis.SetBooleanValue(ctx, redis.ZeroCountUpdateKey(dataMarketAddress, dayToCheck.String()), true, 24*time.Hour); err != nil {
							log.Errorf("Error marking zero count update as sent for data market %s on day %s: %v", dataMarketAddress, dayToCheck.String(), err)
							return err
						}

						log.Infof("Marked zero count update as sent for data market %s on day %s. Skipping update to relayer.", dataMarketAddress, dayToCheck.String())
						break
					}

					// Attempt to update using recalculated value
					if err = SendUpdateRewardsToRelayer(ctx, dataMarketAddress, []*big.Int{}, []*big.Int{}, dayToCheck.String(), eligibleNodes); err != nil {
						errorMsg := fmt.Sprintf("🚨 Failed to send rewards update for data market %s on day %s using recalculated count: %v", dataMarketAddress, dayToCheck.String(), err)
						clients.SendFailureNotification(pkgs.SendUpdateRewardsToRelayer, errorMsg, time.Now().String(), "High")
						log.Error(errorMsg)
						retryCount++
						continue
					}

					successMsg := fmt.Sprintf("✅ Successfully updated rewards using recalculated count: Eligible node count for data market %s on day %s: %d", dataMarketAddress, dayToCheck.String(), eligibleNodes)
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
	go func() {
		if err := sendFinalRewards(ctx, epochID); err != nil {
			log.Errorf("Error sending final rewards for epoch %s: %v", epochID, err)
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

	batchArrays(ctx, dataMarketAddress, currentDay.String(), slotIDs, submissionsList, 0)

	return nil
}

func sendFinalRewards(ctx context.Context, currentEpoch *big.Int) error {
	log.Infof("🔍 Initiating day transition check for current epoch: %s", currentEpoch.String())
	var wg sync.WaitGroup

	// Create error channel to collect errors from goroutines
	errChan := make(chan error, len(config.SettingsObj.DataMarketAddresses))

	// Fetch and process day transition epoch markers set concurrently for each data market address
	for _, dataMarketAddress := range config.SettingsObj.DataMarketAddresses {
		wg.Add(1)
		go func(dataMarketAddress string) {
			defer wg.Done()

			// Create timeout context as child of parent context
			timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()

			// Process the data market
			epochMarkerKeys, err := redis.RedisClient.SMembers(timeoutCtx, redis.DayRolloverEpochMarkerSet(dataMarketAddress)).Result()
			if err != nil {
				log.Errorf("Failed to fetch day transition epoch markers from Redis for data market %s: %s", dataMarketAddress, err)
				errChan <- fmt.Errorf("failed to fetch day transition epoch markers: %w", err)
				return
			}

			log.Infof("✅ Fetched %d day transition epoch marker keys for data market %s: %v", len(epochMarkerKeys), dataMarketAddress, epochMarkerKeys)

			// Process each day transition epoch marker key for this data market address
			for _, epochMarkerKey := range epochMarkerKeys {
				select {
				case <-ctx.Done():
					errChan <- ctx.Err()
					return
				default:
					// Retrieve the day transition epoch marker details from Redis
					epochMarkerDetailsJSON, err := redis.RedisClient.Get(timeoutCtx, redis.DayRolloverEpochMarkerDetails(dataMarketAddress, epochMarkerKey)).Result()
					if err != nil {
						log.Errorf("Failed to fetch day transition epoch marker details from Redis for key %s: %s", epochMarkerKey, err)
						continue
					}

					var epochMarkerDetails DayTransitionEpochInfo
					if err := json.Unmarshal([]byte(epochMarkerDetailsJSON), &epochMarkerDetails); err != nil {
						log.Errorf("Failed to unmarshal day transition epoch marker details for key %s: %s", epochMarkerKey, err)
						continue
					}

					log.Debugf("📊 Day transition epoch marker details for key %s: %+v", epochMarkerKey, epochMarkerDetails)

					// Check if the current epoch matches the buffer epoch
					if currentEpoch.Int64() == epochMarkerDetails.BufferEpoch {
						lastKnownDay := epochMarkerDetails.LastKnownDay

						// Fetch the eligible nodes count and slotIDs for the last known day (prev day)
						eligibleNodesCount, eligibleSlotIDs := fetchEligibleSlotIDs(timeoutCtx, dataMarketAddress, lastKnownDay)

						log.Infof("✅ Successfully fetched eligible nodes count for data market %s on day %s: %d", dataMarketAddress, lastKnownDay, eligibleNodesCount)

						// Prepare data for relayer
						slotIDs := make([]*big.Int, 0)
						submissionsList := make([]*big.Int, 0)

						for _, slotID := range eligibleSlotIDs {
							slotIDBigInt, _ := new(big.Int).SetString(slotID, 10)
							lastKnownDayBigInt, _ := new(big.Int).SetString(lastKnownDay, 10)

							var submissionCount *big.Int
							if output, err := MustQuery(timeoutCtx, func() (*big.Int, error) {
								return Instance.SlotSubmissionCount(&bind.CallOpts{}, common.HexToAddress(dataMarketAddress), slotIDBigInt, lastKnownDayBigInt)
							}); err == nil {
								submissionCount = output
							}

							slotIDs = append(slotIDs, slotIDBigInt)
							submissionsList = append(submissionsList, submissionCount)
						}

						// Call batchArrays with the parent context
						batchArrays(timeoutCtx, dataMarketAddress, lastKnownDay, slotIDs, submissionsList, eligibleNodesCount)

						// Remove the epochID and its day transition details from Redis
						epochID := new(big.Int).Sub(currentEpoch, big.NewInt(int64(BufferEpochs)))
						if err := redis.RemoveDayTransitionEpochFromRedis(timeoutCtx, dataMarketAddress, epochID.String()); err != nil {
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
		}(dataMarketAddress)
	}

	// Wait for all data market goroutines to finish
	wg.Wait()
	close(errChan)

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
}

func batchArrays(ctx context.Context, dataMarketAddress, currentDay string, slotIDs, submissionsList []*big.Int, eligibleNodesCount int) {
	// Fetch the batch size from config
	batchSize := config.SettingsObj.RewardsUpdateBatchSize
	var wg sync.WaitGroup
	errChan := make(chan error, len(slotIDs)/batchSize+1)

	// Process the data in batches
	for start := 0; start < len(slotIDs); start += batchSize {
		end := start + batchSize
		if end > len(slotIDs) {
			end = len(slotIDs)
		}

		// Get slices from pool and type assert correctly
		slotIDsBatchPtr := slotIDPool.Get().(*[]*big.Int)
		submissionsBatchPtr := submissionsPool.Get().(*[]*big.Int)

		// Clear the slices before reuse
		*slotIDsBatchPtr = (*slotIDsBatchPtr)[:0]
		*submissionsBatchPtr = (*submissionsBatchPtr)[:0]

		// Copy data to pooled slices
		*slotIDsBatchPtr = append(*slotIDsBatchPtr, slotIDs[start:end]...)
		*submissionsBatchPtr = append(*submissionsBatchPtr, submissionsList[start:end]...)

		wg.Add(1)
		workerPool <- struct{}{} // Acquire worker
		go func(start, end int, slotIDsBatch, submissionsBatch *[]*big.Int) {
			defer wg.Done()
			defer func() { <-workerPool }() // Release worker
			defer func() {
				// Clear slices before returning to pool
				*slotIDsBatch = (*slotIDsBatch)[:0]
				*submissionsBatch = (*submissionsBatch)[:0]
				slotIDPool.Put(slotIDsBatch)
				submissionsPool.Put(submissionsBatch)
			}()

			batchCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()

			if err := SendUpdateRewardsToRelayer(batchCtx, dataMarketAddress, *slotIDsBatch, *submissionsBatch, currentDay, eligibleNodesCount); err != nil {
				errorMsg := fmt.Sprintf("🚨 Relayer batch error in batch %d-%d for data market %s on day %s: %v", start, end, dataMarketAddress, currentDay, err)
				clients.SendFailureNotification(pkgs.SendUpdateRewardsToRelayer, errorMsg, time.Now().String(), "High")
				log.Error(errorMsg)
				errChan <- err
			}
		}(start, end, slotIDsBatchPtr, submissionsBatchPtr)
	}

	// Wait for all batches to complete
	wg.Wait()
	close(errChan)

	// Check for any errors
	for err := range errChan {
		if err != nil {
			log.Errorf("Batch processing error for data market %s on day %s: %v", dataMarketAddress, currentDay, err)
		}
	}

	log.Infof("✅ Successfully sent %d batches to relayer for data market %s on day %s", len(slotIDs)/batchSize, dataMarketAddress, currentDay)
}
