package prost

import (
	"context"
	"crypto/tls"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"submission-sequencer-collector/config"
	"submission-sequencer-collector/pkgs/clients"
	"submission-sequencer-collector/pkgs/contract"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	Client      *ethclient.Client
	Instance    *contract.Contract
	ContractABI abi.ABI
)

func ConfigureClient() {
	rpcClient, err := rpc.DialOptions(context.Background(), config.SettingsObj.ClientUrl, rpc.WithHTTPClient(&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}))
	if err != nil {
		log.Errorf("Failed to connect to client: %s", err)
		log.Fatal(err)
	}

	Client = ethclient.NewClient(rpcClient)
}

func ConfigureContractInstance() {
	Instance, _ = contract.NewContract(common.HexToAddress(config.SettingsObj.ContractAddress), Client)
}

func ConfigureABI() {
	contractABI, err := abi.JSON(strings.NewReader(contract.ContractMetaData.ABI))
	if err != nil {
		log.Errorf("Failed to configure contract ABI: %s", err)
		log.Fatal(err)
	}

	ContractABI = contractABI
}

func UpdateSubmissionLimit(curBlock *big.Int) *big.Int {
	// Fetch snapshot submission window from the contract
	window, err := Instance.SnapshotSubmissionWindow(&bind.CallOpts{}, config.SettingsObj.DataMarketContractAddress)
	if err != nil {
		clients.SendFailureNotification("Contract query error [UpdateSubmissionLimit]", fmt.Sprintf("Failed to fetch snapshot submission window: %s", err.Error()), time.Now().String(), "High")
		log.Errorf("Failed to fetch snapshot submission window: %s\n", err.Error())
	}

	submissionLimit := new(big.Int).Add(curBlock, window)
	submissionLimit = submissionLimit.Add(submissionLimit, big.NewInt(1))

	log.Debugln("Snapshot Submission Limit: ", submissionLimit)

	return submissionLimit
}
