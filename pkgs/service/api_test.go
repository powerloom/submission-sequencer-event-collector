//nolint:errcheck
package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs/prost"
	"submission-sequencer-collector/pkgs/redis"
	"submission-sequencer-collector/pkgs/utils"

	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
)

var mr *miniredis.Miniredis

func TestMain(m *testing.M) {
	var err error
	mr, err = miniredis.Run()
	if err != nil {
		log.Fatalf("could not start miniredis: %v", err)
	}

	// Initialize the config settings
	config.SettingsObj = &config.Settings{
		ContractAddress:     "0xE88E5f64AEB483d7057645326AdDFA24A3B312DF",
		ClientUrl:           "https://rpc-prost1m.powerloom.io",
		DataMarketAddresses: []string{"0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"},
		RedisHost:           mr.Host(),
		RedisPort:           mr.Port(),
		RedisDB:             "0",
	}

	utils.InitLogger()
	redis.RedisClient = redis.NewRedisClient()

	prost.ConfigureClient()
	prost.ConfigureContractInstance()

	m.Run()

	mr.Close()
}

func TestHandleTotalSubmissions(t *testing.T) {
	// Set the authentication read token
	config.SettingsObj.AuthReadToken = "valid-token"

	// Set the current day
	redis.Set(context.Background(), redis.GetCurrentDayKey("0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"), "5")

	// Set total submission count for each day
	redis.Set(context.Background(), redis.TotalSubmissionsCountKey("5", "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"), "100")
	redis.Set(context.Background(), redis.TotalSubmissionsCountKey("4", "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"), "80")
	redis.Set(context.Background(), redis.TotalSubmissionsCountKey("3", "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"), "150")
	redis.Set(context.Background(), redis.TotalSubmissionsCountKey("2", "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"), "60")
	redis.Set(context.Background(), redis.TotalSubmissionsCountKey("1", "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"), "50")

	tests := []struct {
		name       string
		body       string
		statusCode int
		response   []DailySubmissions
	}{
		{
			name:       "Valid token, past days 1",
			body:       `{"token": "valid-token", "past_days": 1, "data_market_address": "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"}`,
			statusCode: http.StatusOK,
			response: []DailySubmissions{
				{Day: 5, Submissions: 100},
			},
		},
		{
			name:       "Valid token, past days 3",
			body:       `{"token": "valid-token", "past_days": 3, "data_market_address": "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"}`,
			statusCode: http.StatusOK,
			response: []DailySubmissions{
				{Day: 5, Submissions: 100},
				{Day: 4, Submissions: 80},
				{Day: 3, Submissions: 150},
			},
		},
		{
			name:       "Valid token, total submissions till date",
			body:       `{"token": "valid-token", "past_days": 5, "data_market_address": "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"}`,
			statusCode: http.StatusOK,
			response: []DailySubmissions{
				{Day: 5, Submissions: 100},
				{Day: 4, Submissions: 80},
				{Day: 3, Submissions: 150},
				{Day: 2, Submissions: 60},
				{Day: 1, Submissions: 50},
			},
		},
		{
			name:       "Valid token, negative past days",
			body:       `{"token": "valid-token", "past_days": -1, "data_market_address": "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"}`,
			statusCode: http.StatusBadRequest,
			response:   nil,
		},
		{
			name:       "Invalid token",
			body:       `{"token": "invalid-token", "past_days": 1, "data_market_address": "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200c"}`,
			statusCode: http.StatusUnauthorized,
			response:   nil,
		},
		{
			name:       "Invalid Data Market Address",
			body:       `{"token": "valid-token", "past_days": 1, "data_market_address": "0x0C2E22fe7526fAeF28E7A58c84f8723dEFcE200d"}`,
			statusCode: http.StatusBadRequest,
			response:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/totalSubmissions", strings.NewReader(tt.body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handleTotalSubmissions)
			testHandler := RequestMiddleware(handler)
			testHandler.ServeHTTP(rr, req)

			responseBody := rr.Body.String()
			t.Log("Response Body:", responseBody)

			assert.Equal(t, tt.statusCode, rr.Code)

			if tt.statusCode == http.StatusOK {
				var response struct {
					Info struct {
						Success  bool               `json:"success"`
						Response []DailySubmissions `json:"response"`
					} `json:"info"`
					RequestID string `json:"request_id"`
				}

				err := json.NewDecoder(rr.Body).Decode(&response)
				assert.NoError(t, err)

				err = json.Unmarshal([]byte(responseBody), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.response, response.Info.Response)
			}
		})
	}
}
