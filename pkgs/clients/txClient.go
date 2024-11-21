package clients

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"submission-sequencer-collector/config"
	"time"
)

type TxRelayerClient struct {
	url    string
	client *http.Client
}

type SubmissionBatchSizeRequest struct {
	EpochID           *big.Int `json:"epochID"`
	Size              int      `json:"batchSize"`
	AuthToken         string   `json:"authToken"`
	DataMarketAddress string   `json:"dataMarketAddress"`
}

type UpdateRewardsRequest struct {
	DataMarketAddress string     `json:"dataMarketAddress"`
	SlotIDs           []*big.Int `json:"slotIDs"`
	SubmissionsList   []*big.Int `json:"submissionsList"`
	Day               *big.Int   `json:"day"`
	EligibleNodes     int        `json:"eligibleNodes"`
	AuthToken         string     `json:"authToken"`
}

var txRelayerClient *TxRelayerClient

// InitializeTxClient initializes the TxRelayerClient with the provided URL and timeout
func InitializeTxClient(url string, timeout time.Duration) {
	txRelayerClient = &TxRelayerClient{
		url: url,
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

// SendSubmissionBatchSize sends the size of the submission batch for a given epoch
func SendSubmissionBatchSize(dataMarketAddress string, epochID *big.Int, batchSize int) error {
	request := SubmissionBatchSizeRequest{
		EpochID:           epochID,
		Size:              batchSize,
		DataMarketAddress: dataMarketAddress,
		AuthToken:         config.SettingsObj.TxRelayerAuthWriteToken,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("unable to marshal batch size request: %w", err)
	}

	url := fmt.Sprintf("%s/submitBatchSize", txRelayerClient.url)

	resp, err := txRelayerClient.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("unable to send submission batch size request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send submission batch size request, status code: %d", resp.StatusCode)
	}

	return nil
}

// SendUpdateRewardsRequest sends rewards update data to the transaction relayer service
func SendUpdateRewardsRequest(dataMarketAddress string, slotIDs, submissionsList []*big.Int, currentDay string, eligibleNodes int) error {
	day, ok := new(big.Int).SetString(currentDay, 10)
	if !ok {
		return fmt.Errorf("unable to parse current day '%s' into big.Int for update rewards request", currentDay)
	}

	request := UpdateRewardsRequest{
		DataMarketAddress: dataMarketAddress,
		SlotIDs:           slotIDs,
		SubmissionsList:   submissionsList,
		Day:               day,
		EligibleNodes:     eligibleNodes,
		AuthToken:         config.SettingsObj.TxRelayerAuthWriteToken,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("unable to marshal update rewards request: %w", err)
	}

	url := fmt.Sprintf("%s/submitUpdateRewards", txRelayerClient.url)

	resp, err := txRelayerClient.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("unable to send update rewards request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send update rewards request, status code: %d", resp.StatusCode)
	}

	return nil
}
