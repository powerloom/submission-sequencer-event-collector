package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"sort"
	"strconv"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs/prost"
	"submission-sequencer-collector/pkgs/redis"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/protobuf/encoding/protojson"

	_ "submission-sequencer-collector/pkgs/service/docs"

	"submission-sequencer-collector/pkgs"

	log "github.com/sirupsen/logrus"
)

// @title My API Documentation
// @version 1.0
// @description This API handles submissions and provides Swagger documentation
// @termsOfService http://yourterms.com

// @contact.name API Support
// @contact.url http://www.yoursupport.com
// @contact.email support@example.com

// @host localhost:8080
// @BasePath /

type SubmissionsRequest struct {
	Token             string `json:"token"`
	SlotID            int    `json:"slot_id"`
	PastDays          int    `json:"past_days"`
	DataMarketAddress string `json:"data_market_address"`
}

type DailySubmissions struct {
	Day         int   `json:"day"`
	Submissions int64 `json:"submissions"`
}

type EligibleNodesRequest struct {
	Token             string `json:"token"`
	EpochID           int    `json:"epoch_id"`
	PastDays          int    `json:"past_days"`
	DataMarketAddress string `json:"data_market_address"`
}

type EpochDataMarketRequest struct {
	Token             string `json:"token"`
	EpochID           int    `json:"epoch_id"`
	DataMarketAddress string `json:"data_market_address"`
}

type EligibleNodes struct {
	Day     int      `json:"day"`
	Count   int      `json:"eligible_nodes_count"`
	SlotIDs []string `json:"slot_ids"`
}

type BatchCount struct {
	TotalBatches int `json:"total_batches"`
}

type SubmissionDetails struct {
	SubmissionID   string                   `json:"submission_id"`
	SubmissionData *pkgs.SnapshotSubmission `json:"submission_data"`
}

type EpochSubmissionSummary struct {
	SubmissionCount int                 `json:"epoch_submission_count"`
	Submissions     []SubmissionDetails `json:"submissions"`
}

type InfoType[K any] struct {
	Success  bool `json:"success"`
	Response K    `json:"response"`
}

type ResponseArray[K any] []K

type Response[K any] struct {
	Info      InfoType[K] `json:"info"`
	RequestID string      `json:"request_id"`
}

func getDailySubmissions(dataMarketAddress string, slotID int, day *big.Int) int64 {
	if val, err := redis.Get(context.Background(), redis.SlotSubmissionKey(dataMarketAddress, strconv.Itoa(slotID), day.String())); err != nil || val == "" {
		subs, err := prost.MustQuery[*big.Int](context.Background(), func() (*big.Int, error) {
			subs, err := prost.Instance.SlotSubmissionCount(&bind.CallOpts{}, common.HexToAddress(dataMarketAddress), big.NewInt(int64(slotID)), day)
			return subs, err
		})
		if err != nil {
			log.Errorln("Could not fetch submissions from contract: ", err.Error())
			return 0
		}

		return subs.Int64()
	} else {
		submissions, _ := new(big.Int).SetString(val, 10)
		return submissions.Int64()
	}
}

func getDailyEligibleSubmissions(dataMarketAddress string, slotID int, day *big.Int) int64 {
	// Construct the key for fetching eligible submissions from Redis
	key := redis.EligibleSlotSubmissionKey(dataMarketAddress, strconv.Itoa(slotID), day.String())

	// Attempt to get the value from Redis
	val, err := redis.Get(context.Background(), key)
	if err != nil {
		log.Errorf("Failed to fetch eligible submissions for key %s: %v", key, err)
		return 0
	}

	// If the value is found, parse it as a big integer
	if val != "" {
		eligibleSubmissions, ok := new(big.Int).SetString(val, 10)
		if !ok {
			log.Errorf("Failed to parse eligible submissions for key %s: invalid integer format", key)
			return 0
		}

		return eligibleSubmissions.Int64()
	}

	// Return 0 if no value is found in Redis
	return 0
}

func getEligibleSlotIDs(dataMarketAddress string, day *big.Int) []string {
	// Construct the key for fetching eligible slotIDs from Redis
	key := redis.EligibleSlotSubmissionsByDayKey(dataMarketAddress, day.String())

	// Attempt to get the set values from Redis
	eligibleSlotIDs := redis.GetSetKeys(context.Background(), key)

	// Return the list of slotIDs
	return eligibleSlotIDs
}

func getEpochSubmissions(epochSubmissionKey string) (map[string]string, error) {
	// Use HGetAll to retrieve all key-value pairs in the hash
	submissions, err := redis.RedisClient.HGetAll(context.Background(), epochSubmissionKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch epoch submission details from Redis: %v", err)
	}

	return submissions, nil
}

// handleTotalSubmissions godoc
// @Summary Get total submissions
// @Description Retrieves total submission counts for a specific data market address across a specified number of past days
// @Tags Submissions
// @Accept json
// @Produce json
// @Param request body SubmissionsRequest true "Submissions request payload"
// @Success 200 {object} Response[ResponseArray[DailySubmissions]]
// @Failure 400 {string} string "Bad Request: Invalid input parameters (e.g., past days < 1, invalid slotID or invalid data market address)"
// @Failure 401 {string} string "Unauthorized: Incorrect token"
// @Router /totalSubmissions [post]
func handleTotalSubmissions(w http.ResponseWriter, r *http.Request) {
	var request SubmissionsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	if request.PastDays < 1 {
		http.Error(w, "Past days should be at least 1", http.StatusBadRequest)
		return
	}

	slotID := request.SlotID
	if slotID < 1 || slotID > 10000 {
		http.Error(w, fmt.Sprintf("Invalid slotID: %d", slotID), http.StatusBadRequest)
		return
	}

	isValid := false
	for _, dataMarketAddress := range config.SettingsObj.DataMarketAddresses {
		if request.DataMarketAddress == dataMarketAddress {
			isValid = true
			break
		}
	}

	if !isValid {
		http.Error(w, "Invalid Data Market Address!", http.StatusBadRequest)
		return
	}

	day, err := prost.FetchCurrentDay(common.HexToAddress(request.DataMarketAddress))
	if err != nil {
		http.Error(w, "Failed to fetch current day", http.StatusBadRequest)
		return
	}

	currentDay := new(big.Int).Set(day)
	submissionsResponse := make([]DailySubmissions, request.PastDays)

	var wg sync.WaitGroup
	ch := make(chan DailySubmissions, request.PastDays)

	for i := 0; i < request.PastDays; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			day := new(big.Int).Sub(currentDay, big.NewInt(int64(i)))
			subs := getDailySubmissions(request.DataMarketAddress, request.SlotID, day)
			ch <- DailySubmissions{Day: int(day.Int64()), Submissions: subs}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for submission := range ch {
		submissionsResponse[int(currentDay.Int64())-submission.Day] = submission
	}

	info := InfoType[ResponseArray[DailySubmissions]]{
		Success:  true,
		Response: submissionsResponse,
	}

	response := Response[ResponseArray[DailySubmissions]]{
		Info:      info,
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// handleEligibleSubmissions godoc
// @Summary Get eligible submissions
// @Description Retrieves eligible submission counts for a specific data market address across a specified number of past days
// @Tags Submissions
// @Accept json
// @Produce json
// @Param request body SubmissionsRequest true "Submissions request payload"
// @Success 200 {object} Response[ResponseArray[DailySubmissions]]
// @Failure 400 {string} string "Bad Request: Invalid input parameters (e.g., past days < 1, invalid slotID or invalid data market address)"
// @Failure 401 {string} string "Unauthorized: Incorrect token"
// @Router /eligibleSubmissions [post]
func handleEligibleSubmissions(w http.ResponseWriter, r *http.Request) {
	var request SubmissionsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	if request.PastDays < 1 {
		http.Error(w, "Past days should be at least 1", http.StatusBadRequest)
		return
	}

	slotID := request.SlotID
	if slotID < 1 || slotID > 10000 {
		http.Error(w, fmt.Sprintf("Invalid slotID: %d", slotID), http.StatusBadRequest)
		return
	}

	isValid := false
	for _, dataMarketAddress := range config.SettingsObj.DataMarketAddresses {
		if request.DataMarketAddress == dataMarketAddress {
			isValid = true
			break
		}
	}

	if !isValid {
		http.Error(w, "Invalid Data Market Address!", http.StatusBadRequest)
		return
	}

	day, err := prost.FetchCurrentDay(common.HexToAddress(request.DataMarketAddress))
	if err != nil {
		http.Error(w, "Failed to fetch current day", http.StatusBadRequest)
		return
	}

	currentDay := new(big.Int).Set(day)
	submissionsResponse := make([]DailySubmissions, request.PastDays)

	var wg sync.WaitGroup
	ch := make(chan DailySubmissions, request.PastDays)

	for i := 0; i < request.PastDays; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			day := new(big.Int).Sub(currentDay, big.NewInt(int64(i)))
			subs := getDailyEligibleSubmissions(request.DataMarketAddress, request.SlotID, day)
			ch <- DailySubmissions{Day: int(day.Int64()), Submissions: subs}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for submission := range ch {
		submissionsResponse[int(currentDay.Int64())-submission.Day] = submission
	}

	info := InfoType[ResponseArray[DailySubmissions]]{
		Success:  true,
		Response: submissionsResponse,
	}

	response := Response[ResponseArray[DailySubmissions]]{
		Info:      info,
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// handleEligibleNodesCount godoc
// @Summary Get eligible nodes count
// @Description Retrieves the total count of eligible nodes along with their corresponding slot IDs for a specified data market address and epochID across a specified number of past days
// @Tags Eligible Nodes Count
// @Accept json
// @Produce json
// @Param request body EligibleNodesRequest true "Eligible nodes count payload"
// @Success 200 {object} Response[ResponseArray[EligibleNodes]]
// @Failure 400 {string} string "Bad Request: Invalid input parameters (e.g., past days < 1, missing or invalid epochID, or invalid data market address)"
// @Failure 401 {string} string "Unauthorized: Incorrect token"
// @Router /eligibleNodesCount [post]
func handleEligibleNodesCount(w http.ResponseWriter, r *http.Request) {
	var request EligibleNodesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	if request.PastDays < 1 {
		http.Error(w, "Past days should be at least 1", http.StatusBadRequest)
		return
	}

	epochID := request.EpochID
	if epochID <= 0 {
		http.Error(w, "EpochID is missing or invalid", http.StatusBadRequest)
		return
	}

	isValid := false
	for _, dataMarketAddress := range config.SettingsObj.DataMarketAddresses {
		if request.DataMarketAddress == dataMarketAddress {
			isValid = true
			break
		}
	}

	if !isValid {
		http.Error(w, "Invalid Data Market Address!", http.StatusBadRequest)
		return
	}

	day, err := prost.FetchCurrentDay(common.HexToAddress(request.DataMarketAddress))
	if err != nil {
		http.Error(w, "Failed to fetch current day", http.StatusBadRequest)
		return
	}

	currentDay := new(big.Int).Set(day)
	eligibleNodesResponse := make([]EligibleNodes, request.PastDays)

	var wg sync.WaitGroup
	ch := make(chan EligibleNodes, request.PastDays)

	for i := 0; i < request.PastDays; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			day := new(big.Int).Sub(currentDay, big.NewInt(int64(i)))
			slotIDs := getEligibleSlotIDs(request.DataMarketAddress, day) // Fetch eligible slot IDs for the day
			ch <- EligibleNodes{
				Day:     int(day.Int64()),
				Count:   len(slotIDs),
				SlotIDs: slotIDs,
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for eligibleNode := range ch {
		eligibleNodesResponse[currentDay.Int64()-int64(eligibleNode.Day)] = eligibleNode
	}

	info := InfoType[ResponseArray[EligibleNodes]]{
		Success:  true,
		Response: eligibleNodesResponse,
	}

	response := Response[ResponseArray[EligibleNodes]]{
		Info:      info,
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// handleBatchCount godoc
// @Summary Get total batch count
// @Description Retrieves the total number of batches created within a specific epoch for a given data market address
// @Tags Batch Count
// @Accept json
// @Produce json
// @Param request body EpochDataMarketRequest true "Epoch data market request payload"
// @Success 200 {object} Response[BatchCount]
// @Failure 400 {string} string "Bad Request: Invalid input parameters (e.g., missing or invalid epochID, or invalid data market address)"
// @Failure 401 {string} string "Unauthorized: Incorrect token"
// @Router /batchCount [post]
func handleBatchCount(w http.ResponseWriter, r *http.Request) {
	var request EpochDataMarketRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	epochID := request.EpochID
	if epochID <= 0 {
		http.Error(w, "EpochID is missing or invalid", http.StatusBadRequest)
		return
	}

	isValid := false
	for _, dataMarketAddress := range config.SettingsObj.DataMarketAddresses {
		if request.DataMarketAddress == dataMarketAddress {
			isValid = true
			break
		}
	}

	if !isValid {
		http.Error(w, "Invalid Data Market Address!", http.StatusBadRequest)
		return
	}

	// Fetch the batch count from Redis
	batchCountKey := redis.GetBatchCountKey(request.DataMarketAddress, strconv.Itoa(request.EpochID))
	batchCountStr, err := redis.Get(context.Background(), batchCountKey)
	if err != nil {
		http.Error(w, "Internal Server Error: Failed to fetch batch count", http.StatusInternalServerError)
		return
	}

	// Convert batch count to integer
	batchCount, err := strconv.Atoi(batchCountStr)
	if err != nil {
		http.Error(w, "Internal Server Error: Invalid batch count format", http.StatusInternalServerError)
		return
	}

	info := InfoType[BatchCount]{
		Success: true,
		Response: BatchCount{
			TotalBatches: batchCount,
		},
	}

	response := Response[BatchCount]{
		Info:      info,
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleEpochSubmissionDetails(w http.ResponseWriter, r *http.Request) {
	var request EpochDataMarketRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	epochID := request.EpochID
	if epochID <= 0 {
		http.Error(w, "EpochID is missing or invalid", http.StatusBadRequest)
		return
	}

	isValid := false
	for _, dataMarketAddress := range config.SettingsObj.DataMarketAddresses {
		if request.DataMarketAddress == dataMarketAddress {
			isValid = true
			break
		}
	}

	if !isValid {
		http.Error(w, "Invalid Data Market Address!", http.StatusBadRequest)
		return
	}

	// Fetch the epoch submission count from Redis
	submissionCountKey := redis.EpochSubmissionsCount(request.DataMarketAddress, uint64(request.EpochID))
	submissionCountStr, err := redis.Get(context.Background(), submissionCountKey)
	if err != nil {
		http.Error(w, "Internal Server Error: Failed to fetch epoch submission count", http.StatusInternalServerError)
		return
	}

	// Convert submission count to integer
	submissionCount, err := strconv.Atoi(submissionCountStr)
	if err != nil {
		http.Error(w, "Internal Server Error: Invalid epoch submission format", http.StatusInternalServerError)
		return
	}

	// Fetch the epoch submission details from Redis
	epochSubmissionsKey := redis.EpochSubmissionsKey(request.DataMarketAddress, uint64(request.EpochID))
	epochSubmissionDetails, err := getEpochSubmissions(epochSubmissionsKey)
	if err != nil {
		http.Error(w, "Internal Server Error: Failed to fetch epoch submission details", http.StatusInternalServerError)
		return
	}

	submissionDetailsList := make([]SubmissionDetails, 0)
	for submissionID, submissionJSON := range epochSubmissionDetails {
		// Unmarshal the JSON into the SnapshotSubmission struct
		submissionData := pkgs.SnapshotSubmission{}
		err = protojson.Unmarshal([]byte(submissionJSON), &submissionData)
		if err != nil {
			log.Errorf("Failed to unmarshal submission details for ID %s: %v", submissionID, err)
			continue // Skip this submission and move to the next
		}

		// Create a SubmissionDetails object
		details := SubmissionDetails{
			SubmissionID:   submissionID,
			SubmissionData: &submissionData,
		}

		// Append the details to the list
		submissionDetailsList = append(submissionDetailsList, details)
	}

	sort.Slice(submissionDetailsList, func(i, j int) bool {
		return submissionDetailsList[i].SubmissionID < submissionDetailsList[j].SubmissionID
	})

	info := InfoType[EpochSubmissionSummary]{
		Success: true,
		Response: EpochSubmissionSummary{
			SubmissionCount: submissionCount,
			Submissions:     submissionDetailsList,
		},
	}

	response := Response[EpochSubmissionSummary]{
		Info:      info,
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "request_id", requestID)
		r = r.WithContext(ctx)

		log.WithField("request_id", requestID).Infof("Request started for: %s", r.URL.Path)

		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r)

		log.WithField("request_id", requestID).Infof("Request ended")
	})
}

func StartApiServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/totalSubmissions", handleTotalSubmissions)
	mux.HandleFunc("/eligibleSubmissions", handleEligibleSubmissions)
	mux.HandleFunc("/eligibleNodesCount", handleEligibleNodesCount)
	mux.HandleFunc("/batchCount", handleBatchCount)
	mux.HandleFunc("/epochSubmissionDetails", handleEpochSubmissionDetails)

	handler := RequestMiddleware(mux)

	// Serve Swagger UI with the middleware
	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)

	mux.Handle("/swagger/", RequestMiddleware(swaggerHandler))

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
