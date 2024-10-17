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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	Client       *ethclient.Client
	Instance     *contract.Contract
	CurrentBlock *types.Block
	ContractABI  abi.ABI
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

// calculateSubmissionLimitBlock computes the block number when the submission window ends
func calculateSubmissionLimitBlock(epochReleaseBlock *big.Int) (*big.Int, error) {
	// Fetch snapshot submission limit from the contract
	submissionLimit, err := Instance.SnapshotSubmissionWindow(&bind.CallOpts{}, config.SettingsObj.DataMarketContractAddress)
	if err != nil {
		clients.SendFailureNotification("Contract query error [calculateSubmissionLimitBlock]", fmt.Sprintf("Failed to fetch snapshot submission limit: %s", err.Error()), time.Now().String(), "High")
		log.Errorf("Failed to fetch snapshot submission limit: %s\n", err.Error())
		return nil, err
	}

	// Add the submission limit to the epoch release block number
	submissionLimitBlockNum := new(big.Int).Add(epochReleaseBlock, submissionLimit)

	log.Debugln("Snapshot Submission Limit Block Number: ", submissionLimitBlockNum)

	return submissionLimitBlockNum, nil
}
