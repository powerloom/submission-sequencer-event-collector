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
