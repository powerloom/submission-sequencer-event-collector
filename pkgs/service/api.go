package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs/prost"
	"submission-sequencer-collector/pkgs/redis"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "submission-sequencer-collector/pkgs/service/docs"

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
