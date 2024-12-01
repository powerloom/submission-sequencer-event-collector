// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dataMarketContract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PowerloomDataMarketSlotInfo is an auto generated low-level Go binding around an user-defined struct.
type PowerloomDataMarketSlotInfo struct {
	SlotId                  *big.Int
	SnapshotterAddress      common.Address
	RewardPoints            *big.Int
	CurrentDaySnapshotCount *big.Int
}

// DataMarketContractMetaData contains all meta data concerning the DataMarketContract contract.
var DataMarketContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initializer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AdminsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BatchSubmissionsCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"snapshotterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DailyTaskCompletedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DayStartedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"DelayedAttestationSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DelayedBatchSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EpochReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"snapshotterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RewardsDistributedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SequencersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"SnapshotBatchAttestationSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotBatchFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotBatchSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TriggerBatchResubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorAttestationsInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"ValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAY_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EPOCH_SIZE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USE_BLOCK_NUMBER_AS_EPOCH_ID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationSubmissionWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"attestationsReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedCidsRootHash\",\"type\":\"bytes32\"}],\"name\":\"attestationsReceivedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"batchCidAttestationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batchCidDivergentValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"batchCidDivergentValidatorsLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"batchCidSequencerAttestation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"finalizedCidsRootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batchCidToProjects\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"projectids\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"batchCidToProjectsLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchSubmissionWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"checkDynamicConsensusAttestations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"checkSlotTaskStatusForDay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailySnapshotQuota\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dayCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deploymentBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"}],\"name\":\"eligibleNodesForDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"eligibleNodes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"endBatchSubmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochIdToBatchCids\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"batchCids\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"epochIdToBatchSubmissionsCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"batchSubmissionsCompleted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochEnd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochsInADay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"forceCompleteConsensusAttestations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"TRIGGER_BATCH_RESUBMISSION\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"forceSkipEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"getBatchCidToProjects\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getEpochIdToBatchCids\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequencers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"getSlotInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"snapshotterAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentDaySnapshotCount\",\"type\":\"uint256\"}],\"internalType\":\"structPowerloomDataMarket.SlotInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSequencersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSnapshotterCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"epochSize\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainBlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useBlockNumberAsEpochId\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_protocolStateAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"lastFinalizedSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"lastSequencerFinalizedSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dayCounter\",\"type\":\"uint256\"}],\"name\":\"loadCurrentDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"snapshotCount\",\"type\":\"uint256\"}],\"name\":\"loadSlotSubmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"maxAttestationFinalizedRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"finalizedCidsRootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"maxAttestationsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"maxSnapshotsCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAttestationsForConsensus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"projectFirstEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolState\",\"outputs\":[{\"internalType\":\"contractIPowerloomProtocolState\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"releaseEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sequencerId\",\"type\":\"string\"}],\"name\":\"setSequencerId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"slotRewardPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"slotRewardPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"}],\"name\":\"slotSubmissionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"snapshotStatus\",\"outputs\":[{\"internalType\":\"enumPowerloomDataMarket.SnapshotStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshotSubmissionWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedCidsRootHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"SNAPSHOT_BATCH_ATTESTATION_SUBMITTED\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"projectIds\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"snapshotCids\",\"type\":\"string[]\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedCidsRootHash\",\"type\":\"bytes32\"}],\"name\":\"submitSubmissionBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"SNAPSHOT_BATCH_SUBMITTED\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"DELAYED_BATCH_SUBMITTED\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumPowerloomDataMarket.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_status\",\"type\":\"bool[]\"}],\"name\":\"updateAddresses\",\"outputs\":[{\"internalType\":\"enumPowerloomDataMarket.Role\",\"name\":\"ROLE\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newattestationSubmissionWindow\",\"type\":\"uint256\"}],\"name\":\"updateAttestationSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newbatchSubmissionWindow\",\"type\":\"uint256\"}],\"name\":\"updateBatchSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dailySnapshotQuota\",\"type\":\"uint256\"}],\"name\":\"updateDailySnapshotQuota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_daySize\",\"type\":\"uint256\"}],\"name\":\"updateDaySize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eligibleNodes\",\"type\":\"uint256\"}],\"name\":\"updateEligibleNodesForDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateEpochManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAttestationsForConsensus\",\"type\":\"uint256\"}],\"name\":\"updateMinAttestationsForConsensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_protocolState\",\"type\":\"address\"}],\"name\":\"updateProtocolState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRewardPoolSize\",\"type\":\"uint256\"}],\"name\":\"updateRewardPoolSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submissions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"updateRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newsnapshotSubmissionWindow\",\"type\":\"uint256\"}],\"name\":\"updateSnapshotSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DataMarketContractABI is the input ABI used to generate the binding from.
// Deprecated: Use DataMarketContractMetaData.ABI instead.
var DataMarketContractABI = DataMarketContractMetaData.ABI

// DataMarketContract is an auto generated Go binding around an Ethereum contract.
type DataMarketContract struct {
	DataMarketContractCaller     // Read-only binding to the contract
	DataMarketContractTransactor // Write-only binding to the contract
	DataMarketContractFilterer   // Log filterer for contract events
}

// DataMarketContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataMarketContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataMarketContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataMarketContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataMarketContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataMarketContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataMarketContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataMarketContractSession struct {
	Contract     *DataMarketContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DataMarketContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataMarketContractCallerSession struct {
	Contract *DataMarketContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DataMarketContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataMarketContractTransactorSession struct {
	Contract     *DataMarketContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DataMarketContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataMarketContractRaw struct {
	Contract *DataMarketContract // Generic contract binding to access the raw methods on
}

// DataMarketContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataMarketContractCallerRaw struct {
	Contract *DataMarketContractCaller // Generic read-only contract binding to access the raw methods on
}

// DataMarketContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataMarketContractTransactorRaw struct {
	Contract *DataMarketContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataMarketContract creates a new instance of DataMarketContract, bound to a specific deployed contract.
func NewDataMarketContract(address common.Address, backend bind.ContractBackend) (*DataMarketContract, error) {
	contract, err := bindDataMarketContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataMarketContract{DataMarketContractCaller: DataMarketContractCaller{contract: contract}, DataMarketContractTransactor: DataMarketContractTransactor{contract: contract}, DataMarketContractFilterer: DataMarketContractFilterer{contract: contract}}, nil
}

// NewDataMarketContractCaller creates a new read-only instance of DataMarketContract, bound to a specific deployed contract.
func NewDataMarketContractCaller(address common.Address, caller bind.ContractCaller) (*DataMarketContractCaller, error) {
	contract, err := bindDataMarketContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractCaller{contract: contract}, nil
}

// NewDataMarketContractTransactor creates a new write-only instance of DataMarketContract, bound to a specific deployed contract.
func NewDataMarketContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DataMarketContractTransactor, error) {
	contract, err := bindDataMarketContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractTransactor{contract: contract}, nil
}

// NewDataMarketContractFilterer creates a new log filterer instance of DataMarketContract, bound to a specific deployed contract.
func NewDataMarketContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DataMarketContractFilterer, error) {
	contract, err := bindDataMarketContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractFilterer{contract: contract}, nil
}

// bindDataMarketContract binds a generic wrapper to an already deployed contract.
func bindDataMarketContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DataMarketContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataMarketContract *DataMarketContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataMarketContract.Contract.DataMarketContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataMarketContract *DataMarketContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataMarketContract.Contract.DataMarketContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataMarketContract *DataMarketContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataMarketContract.Contract.DataMarketContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataMarketContract *DataMarketContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataMarketContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataMarketContract *DataMarketContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataMarketContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataMarketContract *DataMarketContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataMarketContract.Contract.contract.Transact(opts, method, params...)
}

// DAYSIZE is a free data retrieval call binding the contract method 0xae423daf.
//
// Solidity: function DAY_SIZE() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) DAYSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "DAY_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAYSIZE is a free data retrieval call binding the contract method 0xae423daf.
//
// Solidity: function DAY_SIZE() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) DAYSIZE() (*big.Int, error) {
	return _DataMarketContract.Contract.DAYSIZE(&_DataMarketContract.CallOpts)
}

// DAYSIZE is a free data retrieval call binding the contract method 0xae423daf.
//
// Solidity: function DAY_SIZE() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) DAYSIZE() (*big.Int, error) {
	return _DataMarketContract.Contract.DAYSIZE(&_DataMarketContract.CallOpts)
}

// EPOCHSIZE is a free data retrieval call binding the contract method 0x62656003.
//
// Solidity: function EPOCH_SIZE() view returns(uint8)
func (_DataMarketContract *DataMarketContractCaller) EPOCHSIZE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "EPOCH_SIZE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// EPOCHSIZE is a free data retrieval call binding the contract method 0x62656003.
//
// Solidity: function EPOCH_SIZE() view returns(uint8)
func (_DataMarketContract *DataMarketContractSession) EPOCHSIZE() (uint8, error) {
	return _DataMarketContract.Contract.EPOCHSIZE(&_DataMarketContract.CallOpts)
}

// EPOCHSIZE is a free data retrieval call binding the contract method 0x62656003.
//
// Solidity: function EPOCH_SIZE() view returns(uint8)
func (_DataMarketContract *DataMarketContractCallerSession) EPOCHSIZE() (uint8, error) {
	return _DataMarketContract.Contract.EPOCHSIZE(&_DataMarketContract.CallOpts)
}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x351b6155.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) SOURCECHAINBLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "SOURCE_CHAIN_BLOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x351b6155.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) SOURCECHAINBLOCKTIME() (*big.Int, error) {
	return _DataMarketContract.Contract.SOURCECHAINBLOCKTIME(&_DataMarketContract.CallOpts)
}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x351b6155.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) SOURCECHAINBLOCKTIME() (*big.Int, error) {
	return _DataMarketContract.Contract.SOURCECHAINBLOCKTIME(&_DataMarketContract.CallOpts)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) SOURCECHAINID() (*big.Int, error) {
	return _DataMarketContract.Contract.SOURCECHAINID(&_DataMarketContract.CallOpts)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _DataMarketContract.Contract.SOURCECHAINID(&_DataMarketContract.CallOpts)
}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x2d46247b.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID() view returns(bool)
func (_DataMarketContract *DataMarketContractCaller) USEBLOCKNUMBERASEPOCHID(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "USE_BLOCK_NUMBER_AS_EPOCH_ID")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x2d46247b.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID() view returns(bool)
func (_DataMarketContract *DataMarketContractSession) USEBLOCKNUMBERASEPOCHID() (bool, error) {
	return _DataMarketContract.Contract.USEBLOCKNUMBERASEPOCHID(&_DataMarketContract.CallOpts)
}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x2d46247b.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID() view returns(bool)
func (_DataMarketContract *DataMarketContractCallerSession) USEBLOCKNUMBERASEPOCHID() (bool, error) {
	return _DataMarketContract.Contract.USEBLOCKNUMBERASEPOCHID(&_DataMarketContract.CallOpts)
}

// AttestationSubmissionWindow is a free data retrieval call binding the contract method 0x25129a3f.
//
// Solidity: function attestationSubmissionWindow() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) AttestationSubmissionWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "attestationSubmissionWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationSubmissionWindow is a free data retrieval call binding the contract method 0x25129a3f.
//
// Solidity: function attestationSubmissionWindow() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) AttestationSubmissionWindow() (*big.Int, error) {
	return _DataMarketContract.Contract.AttestationSubmissionWindow(&_DataMarketContract.CallOpts)
}

// AttestationSubmissionWindow is a free data retrieval call binding the contract method 0x25129a3f.
//
// Solidity: function attestationSubmissionWindow() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) AttestationSubmissionWindow() (*big.Int, error) {
	return _DataMarketContract.Contract.AttestationSubmissionWindow(&_DataMarketContract.CallOpts)
}

// AttestationsReceived is a free data retrieval call binding the contract method 0xa27559df.
//
// Solidity: function attestationsReceived(string batchCid, address ) view returns(bool)
func (_DataMarketContract *DataMarketContractCaller) AttestationsReceived(opts *bind.CallOpts, batchCid string, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "attestationsReceived", batchCid, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestationsReceived is a free data retrieval call binding the contract method 0xa27559df.
//
// Solidity: function attestationsReceived(string batchCid, address ) view returns(bool)
func (_DataMarketContract *DataMarketContractSession) AttestationsReceived(batchCid string, arg1 common.Address) (bool, error) {
	return _DataMarketContract.Contract.AttestationsReceived(&_DataMarketContract.CallOpts, batchCid, arg1)
}

// AttestationsReceived is a free data retrieval call binding the contract method 0xa27559df.
//
// Solidity: function attestationsReceived(string batchCid, address ) view returns(bool)
func (_DataMarketContract *DataMarketContractCallerSession) AttestationsReceived(batchCid string, arg1 common.Address) (bool, error) {
	return _DataMarketContract.Contract.AttestationsReceived(&_DataMarketContract.CallOpts, batchCid, arg1)
}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0xa34797e0.
//
// Solidity: function attestationsReceivedCount(string batchCid, bytes32 finalizedCidsRootHash) view returns(uint256 count)
func (_DataMarketContract *DataMarketContractCaller) AttestationsReceivedCount(opts *bind.CallOpts, batchCid string, finalizedCidsRootHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "attestationsReceivedCount", batchCid, finalizedCidsRootHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0xa34797e0.
//
// Solidity: function attestationsReceivedCount(string batchCid, bytes32 finalizedCidsRootHash) view returns(uint256 count)
func (_DataMarketContract *DataMarketContractSession) AttestationsReceivedCount(batchCid string, finalizedCidsRootHash [32]byte) (*big.Int, error) {
	return _DataMarketContract.Contract.AttestationsReceivedCount(&_DataMarketContract.CallOpts, batchCid, finalizedCidsRootHash)
}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0xa34797e0.
//
// Solidity: function attestationsReceivedCount(string batchCid, bytes32 finalizedCidsRootHash) view returns(uint256 count)
func (_DataMarketContract *DataMarketContractCallerSession) AttestationsReceivedCount(batchCid string, finalizedCidsRootHash [32]byte) (*big.Int, error) {
	return _DataMarketContract.Contract.AttestationsReceivedCount(&_DataMarketContract.CallOpts, batchCid, finalizedCidsRootHash)
}

// BatchCidAttestationStatus is a free data retrieval call binding the contract method 0x26ec61e3.
//
// Solidity: function batchCidAttestationStatus(string batchCid) view returns(bool)
func (_DataMarketContract *DataMarketContractCaller) BatchCidAttestationStatus(opts *bind.CallOpts, batchCid string) (bool, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "batchCidAttestationStatus", batchCid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchCidAttestationStatus is a free data retrieval call binding the contract method 0x26ec61e3.
//
// Solidity: function batchCidAttestationStatus(string batchCid) view returns(bool)
func (_DataMarketContract *DataMarketContractSession) BatchCidAttestationStatus(batchCid string) (bool, error) {
	return _DataMarketContract.Contract.BatchCidAttestationStatus(&_DataMarketContract.CallOpts, batchCid)
}

// BatchCidAttestationStatus is a free data retrieval call binding the contract method 0x26ec61e3.
//
// Solidity: function batchCidAttestationStatus(string batchCid) view returns(bool)
func (_DataMarketContract *DataMarketContractCallerSession) BatchCidAttestationStatus(batchCid string) (bool, error) {
	return _DataMarketContract.Contract.BatchCidAttestationStatus(&_DataMarketContract.CallOpts, batchCid)
}

// BatchCidDivergentValidators is a free data retrieval call binding the contract method 0x90e70151.
//
// Solidity: function batchCidDivergentValidators(string batchCid, uint256 ) view returns(address validators)
func (_DataMarketContract *DataMarketContractCaller) BatchCidDivergentValidators(opts *bind.CallOpts, batchCid string, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "batchCidDivergentValidators", batchCid, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchCidDivergentValidators is a free data retrieval call binding the contract method 0x90e70151.
//
// Solidity: function batchCidDivergentValidators(string batchCid, uint256 ) view returns(address validators)
func (_DataMarketContract *DataMarketContractSession) BatchCidDivergentValidators(batchCid string, arg1 *big.Int) (common.Address, error) {
	return _DataMarketContract.Contract.BatchCidDivergentValidators(&_DataMarketContract.CallOpts, batchCid, arg1)
}

// BatchCidDivergentValidators is a free data retrieval call binding the contract method 0x90e70151.
//
// Solidity: function batchCidDivergentValidators(string batchCid, uint256 ) view returns(address validators)
func (_DataMarketContract *DataMarketContractCallerSession) BatchCidDivergentValidators(batchCid string, arg1 *big.Int) (common.Address, error) {
	return _DataMarketContract.Contract.BatchCidDivergentValidators(&_DataMarketContract.CallOpts, batchCid, arg1)
}

// BatchCidDivergentValidatorsLen is a free data retrieval call binding the contract method 0x6f87bcef.
//
// Solidity: function batchCidDivergentValidatorsLen(string batchCid) view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) BatchCidDivergentValidatorsLen(opts *bind.CallOpts, batchCid string) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "batchCidDivergentValidatorsLen", batchCid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchCidDivergentValidatorsLen is a free data retrieval call binding the contract method 0x6f87bcef.
//
// Solidity: function batchCidDivergentValidatorsLen(string batchCid) view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) BatchCidDivergentValidatorsLen(batchCid string) (*big.Int, error) {
	return _DataMarketContract.Contract.BatchCidDivergentValidatorsLen(&_DataMarketContract.CallOpts, batchCid)
}

// BatchCidDivergentValidatorsLen is a free data retrieval call binding the contract method 0x6f87bcef.
//
// Solidity: function batchCidDivergentValidatorsLen(string batchCid) view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) BatchCidDivergentValidatorsLen(batchCid string) (*big.Int, error) {
	return _DataMarketContract.Contract.BatchCidDivergentValidatorsLen(&_DataMarketContract.CallOpts, batchCid)
}

// BatchCidSequencerAttestation is a free data retrieval call binding the contract method 0xfe18613b.
//
// Solidity: function batchCidSequencerAttestation(string batchCid) view returns(bytes32 finalizedCidsRootHash)
func (_DataMarketContract *DataMarketContractCaller) BatchCidSequencerAttestation(opts *bind.CallOpts, batchCid string) ([32]byte, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "batchCidSequencerAttestation", batchCid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchCidSequencerAttestation is a free data retrieval call binding the contract method 0xfe18613b.
//
// Solidity: function batchCidSequencerAttestation(string batchCid) view returns(bytes32 finalizedCidsRootHash)
func (_DataMarketContract *DataMarketContractSession) BatchCidSequencerAttestation(batchCid string) ([32]byte, error) {
	return _DataMarketContract.Contract.BatchCidSequencerAttestation(&_DataMarketContract.CallOpts, batchCid)
}

// BatchCidSequencerAttestation is a free data retrieval call binding the contract method 0xfe18613b.
//
// Solidity: function batchCidSequencerAttestation(string batchCid) view returns(bytes32 finalizedCidsRootHash)
func (_DataMarketContract *DataMarketContractCallerSession) BatchCidSequencerAttestation(batchCid string) ([32]byte, error) {
	return _DataMarketContract.Contract.BatchCidSequencerAttestation(&_DataMarketContract.CallOpts, batchCid)
}

// BatchCidToProjects is a free data retrieval call binding the contract method 0x3b762661.
//
// Solidity: function batchCidToProjects(string batchCid, uint256 ) view returns(string projectids)
func (_DataMarketContract *DataMarketContractCaller) BatchCidToProjects(opts *bind.CallOpts, batchCid string, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "batchCidToProjects", batchCid, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BatchCidToProjects is a free data retrieval call binding the contract method 0x3b762661.
//
// Solidity: function batchCidToProjects(string batchCid, uint256 ) view returns(string projectids)
func (_DataMarketContract *DataMarketContractSession) BatchCidToProjects(batchCid string, arg1 *big.Int) (string, error) {
	return _DataMarketContract.Contract.BatchCidToProjects(&_DataMarketContract.CallOpts, batchCid, arg1)
}

// BatchCidToProjects is a free data retrieval call binding the contract method 0x3b762661.
//
// Solidity: function batchCidToProjects(string batchCid, uint256 ) view returns(string projectids)
func (_DataMarketContract *DataMarketContractCallerSession) BatchCidToProjects(batchCid string, arg1 *big.Int) (string, error) {
	return _DataMarketContract.Contract.BatchCidToProjects(&_DataMarketContract.CallOpts, batchCid, arg1)
}

// BatchCidToProjectsLen is a free data retrieval call binding the contract method 0xc25564e1.
//
// Solidity: function batchCidToProjectsLen(string batchCid) view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) BatchCidToProjectsLen(opts *bind.CallOpts, batchCid string) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "batchCidToProjectsLen", batchCid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchCidToProjectsLen is a free data retrieval call binding the contract method 0xc25564e1.
//
// Solidity: function batchCidToProjectsLen(string batchCid) view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) BatchCidToProjectsLen(batchCid string) (*big.Int, error) {
	return _DataMarketContract.Contract.BatchCidToProjectsLen(&_DataMarketContract.CallOpts, batchCid)
}

// BatchCidToProjectsLen is a free data retrieval call binding the contract method 0xc25564e1.
//
// Solidity: function batchCidToProjectsLen(string batchCid) view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) BatchCidToProjectsLen(batchCid string) (*big.Int, error) {
	return _DataMarketContract.Contract.BatchCidToProjectsLen(&_DataMarketContract.CallOpts, batchCid)
}

// BatchSubmissionWindow is a free data retrieval call binding the contract method 0xb398c290.
//
// Solidity: function batchSubmissionWindow() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) BatchSubmissionWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "batchSubmissionWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchSubmissionWindow is a free data retrieval call binding the contract method 0xb398c290.
//
// Solidity: function batchSubmissionWindow() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) BatchSubmissionWindow() (*big.Int, error) {
	return _DataMarketContract.Contract.BatchSubmissionWindow(&_DataMarketContract.CallOpts)
}

// BatchSubmissionWindow is a free data retrieval call binding the contract method 0xb398c290.
//
// Solidity: function batchSubmissionWindow() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) BatchSubmissionWindow() (*big.Int, error) {
	return _DataMarketContract.Contract.BatchSubmissionWindow(&_DataMarketContract.CallOpts)
}

// CheckDynamicConsensusAttestations is a free data retrieval call binding the contract method 0x418705da.
//
// Solidity: function checkDynamicConsensusAttestations(string batchCid, uint256 epochId) view returns(bool)
func (_DataMarketContract *DataMarketContractCaller) CheckDynamicConsensusAttestations(opts *bind.CallOpts, batchCid string, epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "checkDynamicConsensusAttestations", batchCid, epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDynamicConsensusAttestations is a free data retrieval call binding the contract method 0x418705da.
//
// Solidity: function checkDynamicConsensusAttestations(string batchCid, uint256 epochId) view returns(bool)
func (_DataMarketContract *DataMarketContractSession) CheckDynamicConsensusAttestations(batchCid string, epochId *big.Int) (bool, error) {
	return _DataMarketContract.Contract.CheckDynamicConsensusAttestations(&_DataMarketContract.CallOpts, batchCid, epochId)
}

// CheckDynamicConsensusAttestations is a free data retrieval call binding the contract method 0x418705da.
//
// Solidity: function checkDynamicConsensusAttestations(string batchCid, uint256 epochId) view returns(bool)
func (_DataMarketContract *DataMarketContractCallerSession) CheckDynamicConsensusAttestations(batchCid string, epochId *big.Int) (bool, error) {
	return _DataMarketContract.Contract.CheckDynamicConsensusAttestations(&_DataMarketContract.CallOpts, batchCid, epochId)
}

// CheckSlotTaskStatusForDay is a free data retrieval call binding the contract method 0xd1dd6ddd.
//
// Solidity: function checkSlotTaskStatusForDay(uint256 slotId, uint256 day) view returns(bool)
func (_DataMarketContract *DataMarketContractCaller) CheckSlotTaskStatusForDay(opts *bind.CallOpts, slotId *big.Int, day *big.Int) (bool, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "checkSlotTaskStatusForDay", slotId, day)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckSlotTaskStatusForDay is a free data retrieval call binding the contract method 0xd1dd6ddd.
//
// Solidity: function checkSlotTaskStatusForDay(uint256 slotId, uint256 day) view returns(bool)
func (_DataMarketContract *DataMarketContractSession) CheckSlotTaskStatusForDay(slotId *big.Int, day *big.Int) (bool, error) {
	return _DataMarketContract.Contract.CheckSlotTaskStatusForDay(&_DataMarketContract.CallOpts, slotId, day)
}

// CheckSlotTaskStatusForDay is a free data retrieval call binding the contract method 0xd1dd6ddd.
//
// Solidity: function checkSlotTaskStatusForDay(uint256 slotId, uint256 day) view returns(bool)
func (_DataMarketContract *DataMarketContractCallerSession) CheckSlotTaskStatusForDay(slotId *big.Int, day *big.Int) (bool, error) {
	return _DataMarketContract.Contract.CheckSlotTaskStatusForDay(&_DataMarketContract.CallOpts, slotId, day)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 begin, uint256 end, uint256 epochId)
func (_DataMarketContract *DataMarketContractCaller) CurrentEpoch(opts *bind.CallOpts) (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "currentEpoch")

	outstruct := new(struct {
		Begin   *big.Int
		End     *big.Int
		EpochId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Begin = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 begin, uint256 end, uint256 epochId)
func (_DataMarketContract *DataMarketContractSession) CurrentEpoch() (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	return _DataMarketContract.Contract.CurrentEpoch(&_DataMarketContract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256 begin, uint256 end, uint256 epochId)
func (_DataMarketContract *DataMarketContractCallerSession) CurrentEpoch() (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	return _DataMarketContract.Contract.CurrentEpoch(&_DataMarketContract.CallOpts)
}

// DailySnapshotQuota is a free data retrieval call binding the contract method 0x0220499f.
//
// Solidity: function dailySnapshotQuota() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) DailySnapshotQuota(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "dailySnapshotQuota")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailySnapshotQuota is a free data retrieval call binding the contract method 0x0220499f.
//
// Solidity: function dailySnapshotQuota() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) DailySnapshotQuota() (*big.Int, error) {
	return _DataMarketContract.Contract.DailySnapshotQuota(&_DataMarketContract.CallOpts)
}

// DailySnapshotQuota is a free data retrieval call binding the contract method 0x0220499f.
//
// Solidity: function dailySnapshotQuota() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) DailySnapshotQuota() (*big.Int, error) {
	return _DataMarketContract.Contract.DailySnapshotQuota(&_DataMarketContract.CallOpts)
}

// DayCounter is a free data retrieval call binding the contract method 0x99332c5e.
//
// Solidity: function dayCounter() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) DayCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "dayCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DayCounter is a free data retrieval call binding the contract method 0x99332c5e.
//
// Solidity: function dayCounter() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) DayCounter() (*big.Int, error) {
	return _DataMarketContract.Contract.DayCounter(&_DataMarketContract.CallOpts)
}

// DayCounter is a free data retrieval call binding the contract method 0x99332c5e.
//
// Solidity: function dayCounter() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) DayCounter() (*big.Int, error) {
	return _DataMarketContract.Contract.DayCounter(&_DataMarketContract.CallOpts)
}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0xcf004217.
//
// Solidity: function deploymentBlockNumber() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) DeploymentBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "deploymentBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0xcf004217.
//
// Solidity: function deploymentBlockNumber() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) DeploymentBlockNumber() (*big.Int, error) {
	return _DataMarketContract.Contract.DeploymentBlockNumber(&_DataMarketContract.CallOpts)
}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0xcf004217.
//
// Solidity: function deploymentBlockNumber() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) DeploymentBlockNumber() (*big.Int, error) {
	return _DataMarketContract.Contract.DeploymentBlockNumber(&_DataMarketContract.CallOpts)
}

// EligibleNodesForDay is a free data retrieval call binding the contract method 0x3b54c84e.
//
// Solidity: function eligibleNodesForDay(uint256 dayId) view returns(uint256 eligibleNodes)
func (_DataMarketContract *DataMarketContractCaller) EligibleNodesForDay(opts *bind.CallOpts, dayId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "eligibleNodesForDay", dayId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EligibleNodesForDay is a free data retrieval call binding the contract method 0x3b54c84e.
//
// Solidity: function eligibleNodesForDay(uint256 dayId) view returns(uint256 eligibleNodes)
func (_DataMarketContract *DataMarketContractSession) EligibleNodesForDay(dayId *big.Int) (*big.Int, error) {
	return _DataMarketContract.Contract.EligibleNodesForDay(&_DataMarketContract.CallOpts, dayId)
}

// EligibleNodesForDay is a free data retrieval call binding the contract method 0x3b54c84e.
//
// Solidity: function eligibleNodesForDay(uint256 dayId) view returns(uint256 eligibleNodes)
func (_DataMarketContract *DataMarketContractCallerSession) EligibleNodesForDay(dayId *big.Int) (*big.Int, error) {
	return _DataMarketContract.Contract.EligibleNodesForDay(&_DataMarketContract.CallOpts, dayId)
}

// EpochIdCounter is a free data retrieval call binding the contract method 0x23e78077.
//
// Solidity: function epochIdCounter() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) EpochIdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "epochIdCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochIdCounter is a free data retrieval call binding the contract method 0x23e78077.
//
// Solidity: function epochIdCounter() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) EpochIdCounter() (*big.Int, error) {
	return _DataMarketContract.Contract.EpochIdCounter(&_DataMarketContract.CallOpts)
}

// EpochIdCounter is a free data retrieval call binding the contract method 0x23e78077.
//
// Solidity: function epochIdCounter() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) EpochIdCounter() (*big.Int, error) {
	return _DataMarketContract.Contract.EpochIdCounter(&_DataMarketContract.CallOpts)
}

// EpochIdToBatchCids is a free data retrieval call binding the contract method 0x22152353.
//
// Solidity: function epochIdToBatchCids(uint256 epochId, uint256 ) view returns(string batchCids)
func (_DataMarketContract *DataMarketContractCaller) EpochIdToBatchCids(opts *bind.CallOpts, epochId *big.Int, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "epochIdToBatchCids", epochId, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EpochIdToBatchCids is a free data retrieval call binding the contract method 0x22152353.
//
// Solidity: function epochIdToBatchCids(uint256 epochId, uint256 ) view returns(string batchCids)
func (_DataMarketContract *DataMarketContractSession) EpochIdToBatchCids(epochId *big.Int, arg1 *big.Int) (string, error) {
	return _DataMarketContract.Contract.EpochIdToBatchCids(&_DataMarketContract.CallOpts, epochId, arg1)
}

// EpochIdToBatchCids is a free data retrieval call binding the contract method 0x22152353.
//
// Solidity: function epochIdToBatchCids(uint256 epochId, uint256 ) view returns(string batchCids)
func (_DataMarketContract *DataMarketContractCallerSession) EpochIdToBatchCids(epochId *big.Int, arg1 *big.Int) (string, error) {
	return _DataMarketContract.Contract.EpochIdToBatchCids(&_DataMarketContract.CallOpts, epochId, arg1)
}

// EpochIdToBatchSubmissionsCompleted is a free data retrieval call binding the contract method 0xc9ba561b.
//
// Solidity: function epochIdToBatchSubmissionsCompleted(uint256 epochId) view returns(bool batchSubmissionsCompleted)
func (_DataMarketContract *DataMarketContractCaller) EpochIdToBatchSubmissionsCompleted(opts *bind.CallOpts, epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "epochIdToBatchSubmissionsCompleted", epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochIdToBatchSubmissionsCompleted is a free data retrieval call binding the contract method 0xc9ba561b.
//
// Solidity: function epochIdToBatchSubmissionsCompleted(uint256 epochId) view returns(bool batchSubmissionsCompleted)
func (_DataMarketContract *DataMarketContractSession) EpochIdToBatchSubmissionsCompleted(epochId *big.Int) (bool, error) {
	return _DataMarketContract.Contract.EpochIdToBatchSubmissionsCompleted(&_DataMarketContract.CallOpts, epochId)
}

// EpochIdToBatchSubmissionsCompleted is a free data retrieval call binding the contract method 0xc9ba561b.
//
// Solidity: function epochIdToBatchSubmissionsCompleted(uint256 epochId) view returns(bool batchSubmissionsCompleted)
func (_DataMarketContract *DataMarketContractCallerSession) EpochIdToBatchSubmissionsCompleted(epochId *big.Int) (bool, error) {
	return _DataMarketContract.Contract.EpochIdToBatchSubmissionsCompleted(&_DataMarketContract.CallOpts, epochId)
}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_DataMarketContract *DataMarketContractCaller) EpochInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "epochInfo", arg0)

	outstruct := new(struct {
		Timestamp   *big.Int
		Blocknumber *big.Int
		EpochEnd    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Blocknumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_DataMarketContract *DataMarketContractSession) EpochInfo(arg0 *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	return _DataMarketContract.Contract.EpochInfo(&_DataMarketContract.CallOpts, arg0)
}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_DataMarketContract *DataMarketContractCallerSession) EpochInfo(arg0 *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	return _DataMarketContract.Contract.EpochInfo(&_DataMarketContract.CallOpts, arg0)
}

// EpochManager is a free data retrieval call binding the contract method 0xe2d2bfe3.
//
// Solidity: function epochManager() view returns(address)
func (_DataMarketContract *DataMarketContractCaller) EpochManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "epochManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EpochManager is a free data retrieval call binding the contract method 0xe2d2bfe3.
//
// Solidity: function epochManager() view returns(address)
func (_DataMarketContract *DataMarketContractSession) EpochManager() (common.Address, error) {
	return _DataMarketContract.Contract.EpochManager(&_DataMarketContract.CallOpts)
}

// EpochManager is a free data retrieval call binding the contract method 0xe2d2bfe3.
//
// Solidity: function epochManager() view returns(address)
func (_DataMarketContract *DataMarketContractCallerSession) EpochManager() (common.Address, error) {
	return _DataMarketContract.Contract.EpochManager(&_DataMarketContract.CallOpts)
}

// EpochsInADay is a free data retrieval call binding the contract method 0xe3042b17.
//
// Solidity: function epochsInADay() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) EpochsInADay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "epochsInADay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochsInADay is a free data retrieval call binding the contract method 0xe3042b17.
//
// Solidity: function epochsInADay() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) EpochsInADay() (*big.Int, error) {
	return _DataMarketContract.Contract.EpochsInADay(&_DataMarketContract.CallOpts)
}

// EpochsInADay is a free data retrieval call binding the contract method 0xe3042b17.
//
// Solidity: function epochsInADay() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) EpochsInADay() (*big.Int, error) {
	return _DataMarketContract.Contract.EpochsInADay(&_DataMarketContract.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_DataMarketContract *DataMarketContractCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_DataMarketContract *DataMarketContractSession) GetAdmins() ([]common.Address, error) {
	return _DataMarketContract.Contract.GetAdmins(&_DataMarketContract.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_DataMarketContract *DataMarketContractCallerSession) GetAdmins() ([]common.Address, error) {
	return _DataMarketContract.Contract.GetAdmins(&_DataMarketContract.CallOpts)
}

// GetBatchCidToProjects is a free data retrieval call binding the contract method 0x8cdf3f67.
//
// Solidity: function getBatchCidToProjects(string batchCid) view returns(string[])
func (_DataMarketContract *DataMarketContractCaller) GetBatchCidToProjects(opts *bind.CallOpts, batchCid string) ([]string, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "getBatchCidToProjects", batchCid)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetBatchCidToProjects is a free data retrieval call binding the contract method 0x8cdf3f67.
//
// Solidity: function getBatchCidToProjects(string batchCid) view returns(string[])
func (_DataMarketContract *DataMarketContractSession) GetBatchCidToProjects(batchCid string) ([]string, error) {
	return _DataMarketContract.Contract.GetBatchCidToProjects(&_DataMarketContract.CallOpts, batchCid)
}

// GetBatchCidToProjects is a free data retrieval call binding the contract method 0x8cdf3f67.
//
// Solidity: function getBatchCidToProjects(string batchCid) view returns(string[])
func (_DataMarketContract *DataMarketContractCallerSession) GetBatchCidToProjects(batchCid string) ([]string, error) {
	return _DataMarketContract.Contract.GetBatchCidToProjects(&_DataMarketContract.CallOpts, batchCid)
}

// GetEpochIdToBatchCids is a free data retrieval call binding the contract method 0xcad1132a.
//
// Solidity: function getEpochIdToBatchCids(uint256 epochId) view returns(string[])
func (_DataMarketContract *DataMarketContractCaller) GetEpochIdToBatchCids(opts *bind.CallOpts, epochId *big.Int) ([]string, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "getEpochIdToBatchCids", epochId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetEpochIdToBatchCids is a free data retrieval call binding the contract method 0xcad1132a.
//
// Solidity: function getEpochIdToBatchCids(uint256 epochId) view returns(string[])
func (_DataMarketContract *DataMarketContractSession) GetEpochIdToBatchCids(epochId *big.Int) ([]string, error) {
	return _DataMarketContract.Contract.GetEpochIdToBatchCids(&_DataMarketContract.CallOpts, epochId)
}

// GetEpochIdToBatchCids is a free data retrieval call binding the contract method 0xcad1132a.
//
// Solidity: function getEpochIdToBatchCids(uint256 epochId) view returns(string[])
func (_DataMarketContract *DataMarketContractCallerSession) GetEpochIdToBatchCids(epochId *big.Int) ([]string, error) {
	return _DataMarketContract.Contract.GetEpochIdToBatchCids(&_DataMarketContract.CallOpts, epochId)
}

// GetEpochManager is a free data retrieval call binding the contract method 0xf3ac5e92.
//
// Solidity: function getEpochManager() view returns(address)
func (_DataMarketContract *DataMarketContractCaller) GetEpochManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "getEpochManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEpochManager is a free data retrieval call binding the contract method 0xf3ac5e92.
//
// Solidity: function getEpochManager() view returns(address)
func (_DataMarketContract *DataMarketContractSession) GetEpochManager() (common.Address, error) {
	return _DataMarketContract.Contract.GetEpochManager(&_DataMarketContract.CallOpts)
}

// GetEpochManager is a free data retrieval call binding the contract method 0xf3ac5e92.
//
// Solidity: function getEpochManager() view returns(address)
func (_DataMarketContract *DataMarketContractCallerSession) GetEpochManager() (common.Address, error) {
	return _DataMarketContract.Contract.GetEpochManager(&_DataMarketContract.CallOpts)
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns(address[])
func (_DataMarketContract *DataMarketContractCaller) GetSequencers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "getSequencers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns(address[])
func (_DataMarketContract *DataMarketContractSession) GetSequencers() ([]common.Address, error) {
	return _DataMarketContract.Contract.GetSequencers(&_DataMarketContract.CallOpts)
}

// GetSequencers is a free data retrieval call binding the contract method 0x125c5f16.
//
// Solidity: function getSequencers() view returns(address[])
func (_DataMarketContract *DataMarketContractCallerSession) GetSequencers() ([]common.Address, error) {
	return _DataMarketContract.Contract.GetSequencers(&_DataMarketContract.CallOpts)
}

// GetSlotInfo is a free data retrieval call binding the contract method 0xbe20f9ac.
//
// Solidity: function getSlotInfo(uint256 slotId) view returns((uint256,address,uint256,uint256))
func (_DataMarketContract *DataMarketContractCaller) GetSlotInfo(opts *bind.CallOpts, slotId *big.Int) (PowerloomDataMarketSlotInfo, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "getSlotInfo", slotId)

	if err != nil {
		return *new(PowerloomDataMarketSlotInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PowerloomDataMarketSlotInfo)).(*PowerloomDataMarketSlotInfo)

	return out0, err

}

// GetSlotInfo is a free data retrieval call binding the contract method 0xbe20f9ac.
//
// Solidity: function getSlotInfo(uint256 slotId) view returns((uint256,address,uint256,uint256))
func (_DataMarketContract *DataMarketContractSession) GetSlotInfo(slotId *big.Int) (PowerloomDataMarketSlotInfo, error) {
	return _DataMarketContract.Contract.GetSlotInfo(&_DataMarketContract.CallOpts, slotId)
}

// GetSlotInfo is a free data retrieval call binding the contract method 0xbe20f9ac.
//
// Solidity: function getSlotInfo(uint256 slotId) view returns((uint256,address,uint256,uint256))
func (_DataMarketContract *DataMarketContractCallerSession) GetSlotInfo(slotId *big.Int) (PowerloomDataMarketSlotInfo, error) {
	return _DataMarketContract.Contract.GetSlotInfo(&_DataMarketContract.CallOpts, slotId)
}

// GetTotalSequencersCount is a free data retrieval call binding the contract method 0xab65a54c.
//
// Solidity: function getTotalSequencersCount() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) GetTotalSequencersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "getTotalSequencersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSequencersCount is a free data retrieval call binding the contract method 0xab65a54c.
//
// Solidity: function getTotalSequencersCount() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) GetTotalSequencersCount() (*big.Int, error) {
	return _DataMarketContract.Contract.GetTotalSequencersCount(&_DataMarketContract.CallOpts)
}

// GetTotalSequencersCount is a free data retrieval call binding the contract method 0xab65a54c.
//
// Solidity: function getTotalSequencersCount() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) GetTotalSequencersCount() (*big.Int, error) {
	return _DataMarketContract.Contract.GetTotalSequencersCount(&_DataMarketContract.CallOpts)
}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) GetTotalSnapshotterCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "getTotalSnapshotterCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) GetTotalSnapshotterCount() (*big.Int, error) {
	return _DataMarketContract.Contract.GetTotalSnapshotterCount(&_DataMarketContract.CallOpts)
}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) GetTotalSnapshotterCount() (*big.Int, error) {
	return _DataMarketContract.Contract.GetTotalSnapshotterCount(&_DataMarketContract.CallOpts)
}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0x983d52e7.
//
// Solidity: function getTotalValidatorsCount() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) GetTotalValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "getTotalValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0x983d52e7.
//
// Solidity: function getTotalValidatorsCount() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) GetTotalValidatorsCount() (*big.Int, error) {
	return _DataMarketContract.Contract.GetTotalValidatorsCount(&_DataMarketContract.CallOpts)
}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0x983d52e7.
//
// Solidity: function getTotalValidatorsCount() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) GetTotalValidatorsCount() (*big.Int, error) {
	return _DataMarketContract.Contract.GetTotalValidatorsCount(&_DataMarketContract.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_DataMarketContract *DataMarketContractCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_DataMarketContract *DataMarketContractSession) GetValidators() ([]common.Address, error) {
	return _DataMarketContract.Contract.GetValidators(&_DataMarketContract.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_DataMarketContract *DataMarketContractCallerSession) GetValidators() ([]common.Address, error) {
	return _DataMarketContract.Contract.GetValidators(&_DataMarketContract.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_DataMarketContract *DataMarketContractCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_DataMarketContract *DataMarketContractSession) IsInitialized() (bool, error) {
	return _DataMarketContract.Contract.IsInitialized(&_DataMarketContract.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_DataMarketContract *DataMarketContractCallerSession) IsInitialized() (bool, error) {
	return _DataMarketContract.Contract.IsInitialized(&_DataMarketContract.CallOpts)
}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0x4ea16b0a.
//
// Solidity: function lastFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_DataMarketContract *DataMarketContractCaller) LastFinalizedSnapshot(opts *bind.CallOpts, projectId string) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "lastFinalizedSnapshot", projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0x4ea16b0a.
//
// Solidity: function lastFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_DataMarketContract *DataMarketContractSession) LastFinalizedSnapshot(projectId string) (*big.Int, error) {
	return _DataMarketContract.Contract.LastFinalizedSnapshot(&_DataMarketContract.CallOpts, projectId)
}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0x4ea16b0a.
//
// Solidity: function lastFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_DataMarketContract *DataMarketContractCallerSession) LastFinalizedSnapshot(projectId string) (*big.Int, error) {
	return _DataMarketContract.Contract.LastFinalizedSnapshot(&_DataMarketContract.CallOpts, projectId)
}

// LastSequencerFinalizedSnapshot is a free data retrieval call binding the contract method 0x7cd4adee.
//
// Solidity: function lastSequencerFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_DataMarketContract *DataMarketContractCaller) LastSequencerFinalizedSnapshot(opts *bind.CallOpts, projectId string) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "lastSequencerFinalizedSnapshot", projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSequencerFinalizedSnapshot is a free data retrieval call binding the contract method 0x7cd4adee.
//
// Solidity: function lastSequencerFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_DataMarketContract *DataMarketContractSession) LastSequencerFinalizedSnapshot(projectId string) (*big.Int, error) {
	return _DataMarketContract.Contract.LastSequencerFinalizedSnapshot(&_DataMarketContract.CallOpts, projectId)
}

// LastSequencerFinalizedSnapshot is a free data retrieval call binding the contract method 0x7cd4adee.
//
// Solidity: function lastSequencerFinalizedSnapshot(string projectId) view returns(uint256 epochId)
func (_DataMarketContract *DataMarketContractCallerSession) LastSequencerFinalizedSnapshot(projectId string) (*big.Int, error) {
	return _DataMarketContract.Contract.LastSequencerFinalizedSnapshot(&_DataMarketContract.CallOpts, projectId)
}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0x6ea11f2f.
//
// Solidity: function maxAttestationFinalizedRootHash(string batchCid) view returns(bytes32 finalizedCidsRootHash)
func (_DataMarketContract *DataMarketContractCaller) MaxAttestationFinalizedRootHash(opts *bind.CallOpts, batchCid string) ([32]byte, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "maxAttestationFinalizedRootHash", batchCid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0x6ea11f2f.
//
// Solidity: function maxAttestationFinalizedRootHash(string batchCid) view returns(bytes32 finalizedCidsRootHash)
func (_DataMarketContract *DataMarketContractSession) MaxAttestationFinalizedRootHash(batchCid string) ([32]byte, error) {
	return _DataMarketContract.Contract.MaxAttestationFinalizedRootHash(&_DataMarketContract.CallOpts, batchCid)
}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0x6ea11f2f.
//
// Solidity: function maxAttestationFinalizedRootHash(string batchCid) view returns(bytes32 finalizedCidsRootHash)
func (_DataMarketContract *DataMarketContractCallerSession) MaxAttestationFinalizedRootHash(batchCid string) ([32]byte, error) {
	return _DataMarketContract.Contract.MaxAttestationFinalizedRootHash(&_DataMarketContract.CallOpts, batchCid)
}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0x9e09fbce.
//
// Solidity: function maxAttestationsCount(string batchCid) view returns(uint256 count)
func (_DataMarketContract *DataMarketContractCaller) MaxAttestationsCount(opts *bind.CallOpts, batchCid string) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "maxAttestationsCount", batchCid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0x9e09fbce.
//
// Solidity: function maxAttestationsCount(string batchCid) view returns(uint256 count)
func (_DataMarketContract *DataMarketContractSession) MaxAttestationsCount(batchCid string) (*big.Int, error) {
	return _DataMarketContract.Contract.MaxAttestationsCount(&_DataMarketContract.CallOpts, batchCid)
}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0x9e09fbce.
//
// Solidity: function maxAttestationsCount(string batchCid) view returns(uint256 count)
func (_DataMarketContract *DataMarketContractCallerSession) MaxAttestationsCount(batchCid string) (*big.Int, error) {
	return _DataMarketContract.Contract.MaxAttestationsCount(&_DataMarketContract.CallOpts, batchCid)
}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0xc2b97d4c.
//
// Solidity: function maxSnapshotsCid(string projectId, uint256 epochId) view returns(string)
func (_DataMarketContract *DataMarketContractCaller) MaxSnapshotsCid(opts *bind.CallOpts, projectId string, epochId *big.Int) (string, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "maxSnapshotsCid", projectId, epochId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0xc2b97d4c.
//
// Solidity: function maxSnapshotsCid(string projectId, uint256 epochId) view returns(string)
func (_DataMarketContract *DataMarketContractSession) MaxSnapshotsCid(projectId string, epochId *big.Int) (string, error) {
	return _DataMarketContract.Contract.MaxSnapshotsCid(&_DataMarketContract.CallOpts, projectId, epochId)
}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0xc2b97d4c.
//
// Solidity: function maxSnapshotsCid(string projectId, uint256 epochId) view returns(string)
func (_DataMarketContract *DataMarketContractCallerSession) MaxSnapshotsCid(projectId string, epochId *big.Int) (string, error) {
	return _DataMarketContract.Contract.MaxSnapshotsCid(&_DataMarketContract.CallOpts, projectId, epochId)
}

// MinAttestationsForConsensus is a free data retrieval call binding the contract method 0x0f827df9.
//
// Solidity: function minAttestationsForConsensus() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) MinAttestationsForConsensus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "minAttestationsForConsensus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAttestationsForConsensus is a free data retrieval call binding the contract method 0x0f827df9.
//
// Solidity: function minAttestationsForConsensus() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) MinAttestationsForConsensus() (*big.Int, error) {
	return _DataMarketContract.Contract.MinAttestationsForConsensus(&_DataMarketContract.CallOpts)
}

// MinAttestationsForConsensus is a free data retrieval call binding the contract method 0x0f827df9.
//
// Solidity: function minAttestationsForConsensus() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) MinAttestationsForConsensus() (*big.Int, error) {
	return _DataMarketContract.Contract.MinAttestationsForConsensus(&_DataMarketContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataMarketContract *DataMarketContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataMarketContract *DataMarketContractSession) Owner() (common.Address, error) {
	return _DataMarketContract.Contract.Owner(&_DataMarketContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataMarketContract *DataMarketContractCallerSession) Owner() (common.Address, error) {
	return _DataMarketContract.Contract.Owner(&_DataMarketContract.CallOpts)
}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0xfa30dbe0.
//
// Solidity: function projectFirstEpochId(string projectId) view returns(uint256 epochId)
func (_DataMarketContract *DataMarketContractCaller) ProjectFirstEpochId(opts *bind.CallOpts, projectId string) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "projectFirstEpochId", projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0xfa30dbe0.
//
// Solidity: function projectFirstEpochId(string projectId) view returns(uint256 epochId)
func (_DataMarketContract *DataMarketContractSession) ProjectFirstEpochId(projectId string) (*big.Int, error) {
	return _DataMarketContract.Contract.ProjectFirstEpochId(&_DataMarketContract.CallOpts, projectId)
}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0xfa30dbe0.
//
// Solidity: function projectFirstEpochId(string projectId) view returns(uint256 epochId)
func (_DataMarketContract *DataMarketContractCallerSession) ProjectFirstEpochId(projectId string) (*big.Int, error) {
	return _DataMarketContract.Contract.ProjectFirstEpochId(&_DataMarketContract.CallOpts, projectId)
}

// ProtocolState is a free data retrieval call binding the contract method 0x0b88f3a2.
//
// Solidity: function protocolState() view returns(address)
func (_DataMarketContract *DataMarketContractCaller) ProtocolState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "protocolState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolState is a free data retrieval call binding the contract method 0x0b88f3a2.
//
// Solidity: function protocolState() view returns(address)
func (_DataMarketContract *DataMarketContractSession) ProtocolState() (common.Address, error) {
	return _DataMarketContract.Contract.ProtocolState(&_DataMarketContract.CallOpts)
}

// ProtocolState is a free data retrieval call binding the contract method 0x0b88f3a2.
//
// Solidity: function protocolState() view returns(address)
func (_DataMarketContract *DataMarketContractCallerSession) ProtocolState() (common.Address, error) {
	return _DataMarketContract.Contract.ProtocolState(&_DataMarketContract.CallOpts)
}

// RewardPoolSize is a free data retrieval call binding the contract method 0x211b827d.
//
// Solidity: function rewardPoolSize() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) RewardPoolSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "rewardPoolSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPoolSize is a free data retrieval call binding the contract method 0x211b827d.
//
// Solidity: function rewardPoolSize() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) RewardPoolSize() (*big.Int, error) {
	return _DataMarketContract.Contract.RewardPoolSize(&_DataMarketContract.CallOpts)
}

// RewardPoolSize is a free data retrieval call binding the contract method 0x211b827d.
//
// Solidity: function rewardPoolSize() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) RewardPoolSize() (*big.Int, error) {
	return _DataMarketContract.Contract.RewardPoolSize(&_DataMarketContract.CallOpts)
}

// RewardsEnabled is a free data retrieval call binding the contract method 0x1dafe16b.
//
// Solidity: function rewardsEnabled() view returns(bool)
func (_DataMarketContract *DataMarketContractCaller) RewardsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "rewardsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardsEnabled is a free data retrieval call binding the contract method 0x1dafe16b.
//
// Solidity: function rewardsEnabled() view returns(bool)
func (_DataMarketContract *DataMarketContractSession) RewardsEnabled() (bool, error) {
	return _DataMarketContract.Contract.RewardsEnabled(&_DataMarketContract.CallOpts)
}

// RewardsEnabled is a free data retrieval call binding the contract method 0x1dafe16b.
//
// Solidity: function rewardsEnabled() view returns(bool)
func (_DataMarketContract *DataMarketContractCallerSession) RewardsEnabled() (bool, error) {
	return _DataMarketContract.Contract.RewardsEnabled(&_DataMarketContract.CallOpts)
}

// SequencerId is a free data retrieval call binding the contract method 0x04a78fca.
//
// Solidity: function sequencerId() view returns(string)
func (_DataMarketContract *DataMarketContractCaller) SequencerId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "sequencerId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SequencerId is a free data retrieval call binding the contract method 0x04a78fca.
//
// Solidity: function sequencerId() view returns(string)
func (_DataMarketContract *DataMarketContractSession) SequencerId() (string, error) {
	return _DataMarketContract.Contract.SequencerId(&_DataMarketContract.CallOpts)
}

// SequencerId is a free data retrieval call binding the contract method 0x04a78fca.
//
// Solidity: function sequencerId() view returns(string)
func (_DataMarketContract *DataMarketContractCallerSession) SequencerId() (string, error) {
	return _DataMarketContract.Contract.SequencerId(&_DataMarketContract.CallOpts)
}

// SlotRewardPoints is a free data retrieval call binding the contract method 0x486429e7.
//
// Solidity: function slotRewardPoints(uint256 slotId) view returns(uint256 slotRewardPoints)
func (_DataMarketContract *DataMarketContractCaller) SlotRewardPoints(opts *bind.CallOpts, slotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "slotRewardPoints", slotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotRewardPoints is a free data retrieval call binding the contract method 0x486429e7.
//
// Solidity: function slotRewardPoints(uint256 slotId) view returns(uint256 slotRewardPoints)
func (_DataMarketContract *DataMarketContractSession) SlotRewardPoints(slotId *big.Int) (*big.Int, error) {
	return _DataMarketContract.Contract.SlotRewardPoints(&_DataMarketContract.CallOpts, slotId)
}

// SlotRewardPoints is a free data retrieval call binding the contract method 0x486429e7.
//
// Solidity: function slotRewardPoints(uint256 slotId) view returns(uint256 slotRewardPoints)
func (_DataMarketContract *DataMarketContractCallerSession) SlotRewardPoints(slotId *big.Int) (*big.Int, error) {
	return _DataMarketContract.Contract.SlotRewardPoints(&_DataMarketContract.CallOpts, slotId)
}

// SlotSubmissionCount is a free data retrieval call binding the contract method 0x9dbc5064.
//
// Solidity: function slotSubmissionCount(uint256 slotId, uint256 dayId) view returns(uint256 snapshotCount)
func (_DataMarketContract *DataMarketContractCaller) SlotSubmissionCount(opts *bind.CallOpts, slotId *big.Int, dayId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "slotSubmissionCount", slotId, dayId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotSubmissionCount is a free data retrieval call binding the contract method 0x9dbc5064.
//
// Solidity: function slotSubmissionCount(uint256 slotId, uint256 dayId) view returns(uint256 snapshotCount)
func (_DataMarketContract *DataMarketContractSession) SlotSubmissionCount(slotId *big.Int, dayId *big.Int) (*big.Int, error) {
	return _DataMarketContract.Contract.SlotSubmissionCount(&_DataMarketContract.CallOpts, slotId, dayId)
}

// SlotSubmissionCount is a free data retrieval call binding the contract method 0x9dbc5064.
//
// Solidity: function slotSubmissionCount(uint256 slotId, uint256 dayId) view returns(uint256 snapshotCount)
func (_DataMarketContract *DataMarketContractCallerSession) SlotSubmissionCount(slotId *big.Int, dayId *big.Int) (*big.Int, error) {
	return _DataMarketContract.Contract.SlotSubmissionCount(&_DataMarketContract.CallOpts, slotId, dayId)
}

// SnapshotStatus is a free data retrieval call binding the contract method 0x3aaf384d.
//
// Solidity: function snapshotStatus(string projectId, uint256 epochId) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractCaller) SnapshotStatus(opts *bind.CallOpts, projectId string, epochId *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "snapshotStatus", projectId, epochId)

	outstruct := new(struct {
		Status      uint8
		SnapshotCid string
		Timestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.SnapshotCid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SnapshotStatus is a free data retrieval call binding the contract method 0x3aaf384d.
//
// Solidity: function snapshotStatus(string projectId, uint256 epochId) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractSession) SnapshotStatus(projectId string, epochId *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	return _DataMarketContract.Contract.SnapshotStatus(&_DataMarketContract.CallOpts, projectId, epochId)
}

// SnapshotStatus is a free data retrieval call binding the contract method 0x3aaf384d.
//
// Solidity: function snapshotStatus(string projectId, uint256 epochId) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractCallerSession) SnapshotStatus(projectId string, epochId *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	return _DataMarketContract.Contract.SnapshotStatus(&_DataMarketContract.CallOpts, projectId, epochId)
}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0x059080f6.
//
// Solidity: function snapshotSubmissionWindow() view returns(uint256)
func (_DataMarketContract *DataMarketContractCaller) SnapshotSubmissionWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataMarketContract.contract.Call(opts, &out, "snapshotSubmissionWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0x059080f6.
//
// Solidity: function snapshotSubmissionWindow() view returns(uint256)
func (_DataMarketContract *DataMarketContractSession) SnapshotSubmissionWindow() (*big.Int, error) {
	return _DataMarketContract.Contract.SnapshotSubmissionWindow(&_DataMarketContract.CallOpts)
}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0x059080f6.
//
// Solidity: function snapshotSubmissionWindow() view returns(uint256)
func (_DataMarketContract *DataMarketContractCallerSession) SnapshotSubmissionWindow() (*big.Int, error) {
	return _DataMarketContract.Contract.SnapshotSubmissionWindow(&_DataMarketContract.CallOpts)
}

// EndBatchSubmissions is a paid mutator transaction binding the contract method 0x34c029a9.
//
// Solidity: function endBatchSubmissions(uint256 epochId) returns()
func (_DataMarketContract *DataMarketContractTransactor) EndBatchSubmissions(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "endBatchSubmissions", epochId)
}

// EndBatchSubmissions is a paid mutator transaction binding the contract method 0x34c029a9.
//
// Solidity: function endBatchSubmissions(uint256 epochId) returns()
func (_DataMarketContract *DataMarketContractSession) EndBatchSubmissions(epochId *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.EndBatchSubmissions(&_DataMarketContract.TransactOpts, epochId)
}

// EndBatchSubmissions is a paid mutator transaction binding the contract method 0x34c029a9.
//
// Solidity: function endBatchSubmissions(uint256 epochId) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) EndBatchSubmissions(epochId *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.EndBatchSubmissions(&_DataMarketContract.TransactOpts, epochId)
}

// ForceCompleteConsensusAttestations is a paid mutator transaction binding the contract method 0x116e130f.
//
// Solidity: function forceCompleteConsensusAttestations(string batchCid, uint256 epochId) returns(bool TRIGGER_BATCH_RESUBMISSION)
func (_DataMarketContract *DataMarketContractTransactor) ForceCompleteConsensusAttestations(opts *bind.TransactOpts, batchCid string, epochId *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "forceCompleteConsensusAttestations", batchCid, epochId)
}

// ForceCompleteConsensusAttestations is a paid mutator transaction binding the contract method 0x116e130f.
//
// Solidity: function forceCompleteConsensusAttestations(string batchCid, uint256 epochId) returns(bool TRIGGER_BATCH_RESUBMISSION)
func (_DataMarketContract *DataMarketContractSession) ForceCompleteConsensusAttestations(batchCid string, epochId *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.ForceCompleteConsensusAttestations(&_DataMarketContract.TransactOpts, batchCid, epochId)
}

// ForceCompleteConsensusAttestations is a paid mutator transaction binding the contract method 0x116e130f.
//
// Solidity: function forceCompleteConsensusAttestations(string batchCid, uint256 epochId) returns(bool TRIGGER_BATCH_RESUBMISSION)
func (_DataMarketContract *DataMarketContractTransactorSession) ForceCompleteConsensusAttestations(batchCid string, epochId *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.ForceCompleteConsensusAttestations(&_DataMarketContract.TransactOpts, batchCid, epochId)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0xf537a3e2.
//
// Solidity: function forceSkipEpoch(uint256 begin, uint256 end) returns()
func (_DataMarketContract *DataMarketContractTransactor) ForceSkipEpoch(opts *bind.TransactOpts, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "forceSkipEpoch", begin, end)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0xf537a3e2.
//
// Solidity: function forceSkipEpoch(uint256 begin, uint256 end) returns()
func (_DataMarketContract *DataMarketContractSession) ForceSkipEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.ForceSkipEpoch(&_DataMarketContract.TransactOpts, begin, end)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0xf537a3e2.
//
// Solidity: function forceSkipEpoch(uint256 begin, uint256 end) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) ForceSkipEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.ForceSkipEpoch(&_DataMarketContract.TransactOpts, begin, end)
}

// Initialize is a paid mutator transaction binding the contract method 0x0b372a31.
//
// Solidity: function initialize(address ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId, address _protocolStateAddress) returns()
func (_DataMarketContract *DataMarketContractTransactor) Initialize(opts *bind.TransactOpts, ownerAddress common.Address, epochSize uint8, sourceChainId *big.Int, sourceChainBlockTime *big.Int, useBlockNumberAsEpochId bool, _protocolStateAddress common.Address) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "initialize", ownerAddress, epochSize, sourceChainId, sourceChainBlockTime, useBlockNumberAsEpochId, _protocolStateAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x0b372a31.
//
// Solidity: function initialize(address ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId, address _protocolStateAddress) returns()
func (_DataMarketContract *DataMarketContractSession) Initialize(ownerAddress common.Address, epochSize uint8, sourceChainId *big.Int, sourceChainBlockTime *big.Int, useBlockNumberAsEpochId bool, _protocolStateAddress common.Address) (*types.Transaction, error) {
	return _DataMarketContract.Contract.Initialize(&_DataMarketContract.TransactOpts, ownerAddress, epochSize, sourceChainId, sourceChainBlockTime, useBlockNumberAsEpochId, _protocolStateAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x0b372a31.
//
// Solidity: function initialize(address ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId, address _protocolStateAddress) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) Initialize(ownerAddress common.Address, epochSize uint8, sourceChainId *big.Int, sourceChainBlockTime *big.Int, useBlockNumberAsEpochId bool, _protocolStateAddress common.Address) (*types.Transaction, error) {
	return _DataMarketContract.Contract.Initialize(&_DataMarketContract.TransactOpts, ownerAddress, epochSize, sourceChainId, sourceChainBlockTime, useBlockNumberAsEpochId, _protocolStateAddress)
}

// LoadCurrentDay is a paid mutator transaction binding the contract method 0x82cdfd43.
//
// Solidity: function loadCurrentDay(uint256 _dayCounter) returns()
func (_DataMarketContract *DataMarketContractTransactor) LoadCurrentDay(opts *bind.TransactOpts, _dayCounter *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "loadCurrentDay", _dayCounter)
}

// LoadCurrentDay is a paid mutator transaction binding the contract method 0x82cdfd43.
//
// Solidity: function loadCurrentDay(uint256 _dayCounter) returns()
func (_DataMarketContract *DataMarketContractSession) LoadCurrentDay(_dayCounter *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.LoadCurrentDay(&_DataMarketContract.TransactOpts, _dayCounter)
}

// LoadCurrentDay is a paid mutator transaction binding the contract method 0x82cdfd43.
//
// Solidity: function loadCurrentDay(uint256 _dayCounter) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) LoadCurrentDay(_dayCounter *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.LoadCurrentDay(&_DataMarketContract.TransactOpts, _dayCounter)
}

// LoadSlotSubmissions is a paid mutator transaction binding the contract method 0xbb8ab44b.
//
// Solidity: function loadSlotSubmissions(uint256 slotId, uint256 dayId, uint256 snapshotCount) returns()
func (_DataMarketContract *DataMarketContractTransactor) LoadSlotSubmissions(opts *bind.TransactOpts, slotId *big.Int, dayId *big.Int, snapshotCount *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "loadSlotSubmissions", slotId, dayId, snapshotCount)
}

// LoadSlotSubmissions is a paid mutator transaction binding the contract method 0xbb8ab44b.
//
// Solidity: function loadSlotSubmissions(uint256 slotId, uint256 dayId, uint256 snapshotCount) returns()
func (_DataMarketContract *DataMarketContractSession) LoadSlotSubmissions(slotId *big.Int, dayId *big.Int, snapshotCount *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.LoadSlotSubmissions(&_DataMarketContract.TransactOpts, slotId, dayId, snapshotCount)
}

// LoadSlotSubmissions is a paid mutator transaction binding the contract method 0xbb8ab44b.
//
// Solidity: function loadSlotSubmissions(uint256 slotId, uint256 dayId, uint256 snapshotCount) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) LoadSlotSubmissions(slotId *big.Int, dayId *big.Int, snapshotCount *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.LoadSlotSubmissions(&_DataMarketContract.TransactOpts, slotId, dayId, snapshotCount)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x132c290f.
//
// Solidity: function releaseEpoch(uint256 begin, uint256 end) returns(bool, bool)
func (_DataMarketContract *DataMarketContractTransactor) ReleaseEpoch(opts *bind.TransactOpts, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "releaseEpoch", begin, end)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x132c290f.
//
// Solidity: function releaseEpoch(uint256 begin, uint256 end) returns(bool, bool)
func (_DataMarketContract *DataMarketContractSession) ReleaseEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.ReleaseEpoch(&_DataMarketContract.TransactOpts, begin, end)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x132c290f.
//
// Solidity: function releaseEpoch(uint256 begin, uint256 end) returns(bool, bool)
func (_DataMarketContract *DataMarketContractTransactorSession) ReleaseEpoch(begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.ReleaseEpoch(&_DataMarketContract.TransactOpts, begin, end)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataMarketContract *DataMarketContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataMarketContract *DataMarketContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _DataMarketContract.Contract.RenounceOwnership(&_DataMarketContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataMarketContract *DataMarketContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DataMarketContract.Contract.RenounceOwnership(&_DataMarketContract.TransactOpts)
}

// SetSequencerId is a paid mutator transaction binding the contract method 0x05ebfc34.
//
// Solidity: function setSequencerId(string _sequencerId) returns()
func (_DataMarketContract *DataMarketContractTransactor) SetSequencerId(opts *bind.TransactOpts, _sequencerId string) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "setSequencerId", _sequencerId)
}

// SetSequencerId is a paid mutator transaction binding the contract method 0x05ebfc34.
//
// Solidity: function setSequencerId(string _sequencerId) returns()
func (_DataMarketContract *DataMarketContractSession) SetSequencerId(_sequencerId string) (*types.Transaction, error) {
	return _DataMarketContract.Contract.SetSequencerId(&_DataMarketContract.TransactOpts, _sequencerId)
}

// SetSequencerId is a paid mutator transaction binding the contract method 0x05ebfc34.
//
// Solidity: function setSequencerId(string _sequencerId) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) SetSequencerId(_sequencerId string) (*types.Transaction, error) {
	return _DataMarketContract.Contract.SetSequencerId(&_DataMarketContract.TransactOpts, _sequencerId)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0x0f6c3394.
//
// Solidity: function submitBatchAttestation(string batchCid, uint256 epochId, bytes32 finalizedCidsRootHash) returns(bool SNAPSHOT_BATCH_ATTESTATION_SUBMITTED)
func (_DataMarketContract *DataMarketContractTransactor) SubmitBatchAttestation(opts *bind.TransactOpts, batchCid string, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "submitBatchAttestation", batchCid, epochId, finalizedCidsRootHash)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0x0f6c3394.
//
// Solidity: function submitBatchAttestation(string batchCid, uint256 epochId, bytes32 finalizedCidsRootHash) returns(bool SNAPSHOT_BATCH_ATTESTATION_SUBMITTED)
func (_DataMarketContract *DataMarketContractSession) SubmitBatchAttestation(batchCid string, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _DataMarketContract.Contract.SubmitBatchAttestation(&_DataMarketContract.TransactOpts, batchCid, epochId, finalizedCidsRootHash)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0x0f6c3394.
//
// Solidity: function submitBatchAttestation(string batchCid, uint256 epochId, bytes32 finalizedCidsRootHash) returns(bool SNAPSHOT_BATCH_ATTESTATION_SUBMITTED)
func (_DataMarketContract *DataMarketContractTransactorSession) SubmitBatchAttestation(batchCid string, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _DataMarketContract.Contract.SubmitBatchAttestation(&_DataMarketContract.TransactOpts, batchCid, epochId, finalizedCidsRootHash)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0x06e07f92.
//
// Solidity: function submitSubmissionBatch(string batchCid, uint256 epochId, string[] projectIds, string[] snapshotCids, bytes32 finalizedCidsRootHash) returns(bool SNAPSHOT_BATCH_SUBMITTED, bool DELAYED_BATCH_SUBMITTED)
func (_DataMarketContract *DataMarketContractTransactor) SubmitSubmissionBatch(opts *bind.TransactOpts, batchCid string, epochId *big.Int, projectIds []string, snapshotCids []string, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "submitSubmissionBatch", batchCid, epochId, projectIds, snapshotCids, finalizedCidsRootHash)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0x06e07f92.
//
// Solidity: function submitSubmissionBatch(string batchCid, uint256 epochId, string[] projectIds, string[] snapshotCids, bytes32 finalizedCidsRootHash) returns(bool SNAPSHOT_BATCH_SUBMITTED, bool DELAYED_BATCH_SUBMITTED)
func (_DataMarketContract *DataMarketContractSession) SubmitSubmissionBatch(batchCid string, epochId *big.Int, projectIds []string, snapshotCids []string, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _DataMarketContract.Contract.SubmitSubmissionBatch(&_DataMarketContract.TransactOpts, batchCid, epochId, projectIds, snapshotCids, finalizedCidsRootHash)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0x06e07f92.
//
// Solidity: function submitSubmissionBatch(string batchCid, uint256 epochId, string[] projectIds, string[] snapshotCids, bytes32 finalizedCidsRootHash) returns(bool SNAPSHOT_BATCH_SUBMITTED, bool DELAYED_BATCH_SUBMITTED)
func (_DataMarketContract *DataMarketContractTransactorSession) SubmitSubmissionBatch(batchCid string, epochId *big.Int, projectIds []string, snapshotCids []string, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _DataMarketContract.Contract.SubmitSubmissionBatch(&_DataMarketContract.TransactOpts, batchCid, epochId, projectIds, snapshotCids, finalizedCidsRootHash)
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x95268408.
//
// Solidity: function toggleRewards() returns()
func (_DataMarketContract *DataMarketContractTransactor) ToggleRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "toggleRewards")
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x95268408.
//
// Solidity: function toggleRewards() returns()
func (_DataMarketContract *DataMarketContractSession) ToggleRewards() (*types.Transaction, error) {
	return _DataMarketContract.Contract.ToggleRewards(&_DataMarketContract.TransactOpts)
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x95268408.
//
// Solidity: function toggleRewards() returns()
func (_DataMarketContract *DataMarketContractTransactorSession) ToggleRewards() (*types.Transaction, error) {
	return _DataMarketContract.Contract.ToggleRewards(&_DataMarketContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DataMarketContract *DataMarketContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DataMarketContract *DataMarketContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DataMarketContract.Contract.TransferOwnership(&_DataMarketContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DataMarketContract.Contract.TransferOwnership(&_DataMarketContract.TransactOpts, newOwner)
}

// UpdateAddresses is a paid mutator transaction binding the contract method 0xa5240d23.
//
// Solidity: function updateAddresses(uint8 role, address[] _addresses, bool[] _status) returns(uint8 ROLE)
func (_DataMarketContract *DataMarketContractTransactor) UpdateAddresses(opts *bind.TransactOpts, role uint8, _addresses []common.Address, _status []bool) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateAddresses", role, _addresses, _status)
}

// UpdateAddresses is a paid mutator transaction binding the contract method 0xa5240d23.
//
// Solidity: function updateAddresses(uint8 role, address[] _addresses, bool[] _status) returns(uint8 ROLE)
func (_DataMarketContract *DataMarketContractSession) UpdateAddresses(role uint8, _addresses []common.Address, _status []bool) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateAddresses(&_DataMarketContract.TransactOpts, role, _addresses, _status)
}

// UpdateAddresses is a paid mutator transaction binding the contract method 0xa5240d23.
//
// Solidity: function updateAddresses(uint8 role, address[] _addresses, bool[] _status) returns(uint8 ROLE)
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateAddresses(role uint8, _addresses []common.Address, _status []bool) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateAddresses(&_DataMarketContract.TransactOpts, role, _addresses, _status)
}

// UpdateAttestationSubmissionWindow is a paid mutator transaction binding the contract method 0x82265a87.
//
// Solidity: function updateAttestationSubmissionWindow(uint256 newattestationSubmissionWindow) returns()
func (_DataMarketContract *DataMarketContractTransactor) UpdateAttestationSubmissionWindow(opts *bind.TransactOpts, newattestationSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateAttestationSubmissionWindow", newattestationSubmissionWindow)
}

// UpdateAttestationSubmissionWindow is a paid mutator transaction binding the contract method 0x82265a87.
//
// Solidity: function updateAttestationSubmissionWindow(uint256 newattestationSubmissionWindow) returns()
func (_DataMarketContract *DataMarketContractSession) UpdateAttestationSubmissionWindow(newattestationSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateAttestationSubmissionWindow(&_DataMarketContract.TransactOpts, newattestationSubmissionWindow)
}

// UpdateAttestationSubmissionWindow is a paid mutator transaction binding the contract method 0x82265a87.
//
// Solidity: function updateAttestationSubmissionWindow(uint256 newattestationSubmissionWindow) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateAttestationSubmissionWindow(newattestationSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateAttestationSubmissionWindow(&_DataMarketContract.TransactOpts, newattestationSubmissionWindow)
}

// UpdateBatchSubmissionWindow is a paid mutator transaction binding the contract method 0x31d44744.
//
// Solidity: function updateBatchSubmissionWindow(uint256 newbatchSubmissionWindow) returns()
func (_DataMarketContract *DataMarketContractTransactor) UpdateBatchSubmissionWindow(opts *bind.TransactOpts, newbatchSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateBatchSubmissionWindow", newbatchSubmissionWindow)
}

// UpdateBatchSubmissionWindow is a paid mutator transaction binding the contract method 0x31d44744.
//
// Solidity: function updateBatchSubmissionWindow(uint256 newbatchSubmissionWindow) returns()
func (_DataMarketContract *DataMarketContractSession) UpdateBatchSubmissionWindow(newbatchSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateBatchSubmissionWindow(&_DataMarketContract.TransactOpts, newbatchSubmissionWindow)
}

// UpdateBatchSubmissionWindow is a paid mutator transaction binding the contract method 0x31d44744.
//
// Solidity: function updateBatchSubmissionWindow(uint256 newbatchSubmissionWindow) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateBatchSubmissionWindow(newbatchSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateBatchSubmissionWindow(&_DataMarketContract.TransactOpts, newbatchSubmissionWindow)
}

// UpdateDailySnapshotQuota is a paid mutator transaction binding the contract method 0x6d26e630.
//
// Solidity: function updateDailySnapshotQuota(uint256 _dailySnapshotQuota) returns()
func (_DataMarketContract *DataMarketContractTransactor) UpdateDailySnapshotQuota(opts *bind.TransactOpts, _dailySnapshotQuota *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateDailySnapshotQuota", _dailySnapshotQuota)
}

// UpdateDailySnapshotQuota is a paid mutator transaction binding the contract method 0x6d26e630.
//
// Solidity: function updateDailySnapshotQuota(uint256 _dailySnapshotQuota) returns()
func (_DataMarketContract *DataMarketContractSession) UpdateDailySnapshotQuota(_dailySnapshotQuota *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateDailySnapshotQuota(&_DataMarketContract.TransactOpts, _dailySnapshotQuota)
}

// UpdateDailySnapshotQuota is a paid mutator transaction binding the contract method 0x6d26e630.
//
// Solidity: function updateDailySnapshotQuota(uint256 _dailySnapshotQuota) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateDailySnapshotQuota(_dailySnapshotQuota *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateDailySnapshotQuota(&_DataMarketContract.TransactOpts, _dailySnapshotQuota)
}

// UpdateDaySize is a paid mutator transaction binding the contract method 0x806daf2e.
//
// Solidity: function updateDaySize(uint256 _daySize) returns()
func (_DataMarketContract *DataMarketContractTransactor) UpdateDaySize(opts *bind.TransactOpts, _daySize *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateDaySize", _daySize)
}

// UpdateDaySize is a paid mutator transaction binding the contract method 0x806daf2e.
//
// Solidity: function updateDaySize(uint256 _daySize) returns()
func (_DataMarketContract *DataMarketContractSession) UpdateDaySize(_daySize *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateDaySize(&_DataMarketContract.TransactOpts, _daySize)
}

// UpdateDaySize is a paid mutator transaction binding the contract method 0x806daf2e.
//
// Solidity: function updateDaySize(uint256 _daySize) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateDaySize(_daySize *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateDaySize(&_DataMarketContract.TransactOpts, _daySize)
}

// UpdateEligibleNodesForDay is a paid mutator transaction binding the contract method 0x335e8a75.
//
// Solidity: function updateEligibleNodesForDay(uint256 day, uint256 eligibleNodes) returns()
func (_DataMarketContract *DataMarketContractTransactor) UpdateEligibleNodesForDay(opts *bind.TransactOpts, day *big.Int, eligibleNodes *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateEligibleNodesForDay", day, eligibleNodes)
}

// UpdateEligibleNodesForDay is a paid mutator transaction binding the contract method 0x335e8a75.
//
// Solidity: function updateEligibleNodesForDay(uint256 day, uint256 eligibleNodes) returns()
func (_DataMarketContract *DataMarketContractSession) UpdateEligibleNodesForDay(day *big.Int, eligibleNodes *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateEligibleNodesForDay(&_DataMarketContract.TransactOpts, day, eligibleNodes)
}

// UpdateEligibleNodesForDay is a paid mutator transaction binding the contract method 0x335e8a75.
//
// Solidity: function updateEligibleNodesForDay(uint256 day, uint256 eligibleNodes) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateEligibleNodesForDay(day *big.Int, eligibleNodes *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateEligibleNodesForDay(&_DataMarketContract.TransactOpts, day, eligibleNodes)
}

// UpdateEpochManager is a paid mutator transaction binding the contract method 0x04afb4d2.
//
// Solidity: function updateEpochManager(address _address) returns()
func (_DataMarketContract *DataMarketContractTransactor) UpdateEpochManager(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateEpochManager", _address)
}

// UpdateEpochManager is a paid mutator transaction binding the contract method 0x04afb4d2.
//
// Solidity: function updateEpochManager(address _address) returns()
func (_DataMarketContract *DataMarketContractSession) UpdateEpochManager(_address common.Address) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateEpochManager(&_DataMarketContract.TransactOpts, _address)
}

// UpdateEpochManager is a paid mutator transaction binding the contract method 0x04afb4d2.
//
// Solidity: function updateEpochManager(address _address) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateEpochManager(_address common.Address) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateEpochManager(&_DataMarketContract.TransactOpts, _address)
}

// UpdateMinAttestationsForConsensus is a paid mutator transaction binding the contract method 0x882cd698.
//
// Solidity: function updateMinAttestationsForConsensus(uint256 _minAttestationsForConsensus) returns()
func (_DataMarketContract *DataMarketContractTransactor) UpdateMinAttestationsForConsensus(opts *bind.TransactOpts, _minAttestationsForConsensus *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateMinAttestationsForConsensus", _minAttestationsForConsensus)
}

// UpdateMinAttestationsForConsensus is a paid mutator transaction binding the contract method 0x882cd698.
//
// Solidity: function updateMinAttestationsForConsensus(uint256 _minAttestationsForConsensus) returns()
func (_DataMarketContract *DataMarketContractSession) UpdateMinAttestationsForConsensus(_minAttestationsForConsensus *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateMinAttestationsForConsensus(&_DataMarketContract.TransactOpts, _minAttestationsForConsensus)
}

// UpdateMinAttestationsForConsensus is a paid mutator transaction binding the contract method 0x882cd698.
//
// Solidity: function updateMinAttestationsForConsensus(uint256 _minAttestationsForConsensus) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateMinAttestationsForConsensus(_minAttestationsForConsensus *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateMinAttestationsForConsensus(&_DataMarketContract.TransactOpts, _minAttestationsForConsensus)
}

// UpdateProtocolState is a paid mutator transaction binding the contract method 0x565cda85.
//
// Solidity: function updateProtocolState(address _protocolState) returns()
func (_DataMarketContract *DataMarketContractTransactor) UpdateProtocolState(opts *bind.TransactOpts, _protocolState common.Address) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateProtocolState", _protocolState)
}

// UpdateProtocolState is a paid mutator transaction binding the contract method 0x565cda85.
//
// Solidity: function updateProtocolState(address _protocolState) returns()
func (_DataMarketContract *DataMarketContractSession) UpdateProtocolState(_protocolState common.Address) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateProtocolState(&_DataMarketContract.TransactOpts, _protocolState)
}

// UpdateProtocolState is a paid mutator transaction binding the contract method 0x565cda85.
//
// Solidity: function updateProtocolState(address _protocolState) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateProtocolState(_protocolState common.Address) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateProtocolState(&_DataMarketContract.TransactOpts, _protocolState)
}

// UpdateRewardPoolSize is a paid mutator transaction binding the contract method 0x43c93f35.
//
// Solidity: function updateRewardPoolSize(uint256 newRewardPoolSize) returns()
func (_DataMarketContract *DataMarketContractTransactor) UpdateRewardPoolSize(opts *bind.TransactOpts, newRewardPoolSize *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateRewardPoolSize", newRewardPoolSize)
}

// UpdateRewardPoolSize is a paid mutator transaction binding the contract method 0x43c93f35.
//
// Solidity: function updateRewardPoolSize(uint256 newRewardPoolSize) returns()
func (_DataMarketContract *DataMarketContractSession) UpdateRewardPoolSize(newRewardPoolSize *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateRewardPoolSize(&_DataMarketContract.TransactOpts, newRewardPoolSize)
}

// UpdateRewardPoolSize is a paid mutator transaction binding the contract method 0x43c93f35.
//
// Solidity: function updateRewardPoolSize(uint256 newRewardPoolSize) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateRewardPoolSize(newRewardPoolSize *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateRewardPoolSize(&_DataMarketContract.TransactOpts, newRewardPoolSize)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0xb37f2590.
//
// Solidity: function updateRewards(uint256 slotId, uint256 submissions, uint256 day) returns(bool)
func (_DataMarketContract *DataMarketContractTransactor) UpdateRewards(opts *bind.TransactOpts, slotId *big.Int, submissions *big.Int, day *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateRewards", slotId, submissions, day)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0xb37f2590.
//
// Solidity: function updateRewards(uint256 slotId, uint256 submissions, uint256 day) returns(bool)
func (_DataMarketContract *DataMarketContractSession) UpdateRewards(slotId *big.Int, submissions *big.Int, day *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateRewards(&_DataMarketContract.TransactOpts, slotId, submissions, day)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0xb37f2590.
//
// Solidity: function updateRewards(uint256 slotId, uint256 submissions, uint256 day) returns(bool)
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateRewards(slotId *big.Int, submissions *big.Int, day *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateRewards(&_DataMarketContract.TransactOpts, slotId, submissions, day)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0x9b2f89ce.
//
// Solidity: function updateSnapshotSubmissionWindow(uint256 newsnapshotSubmissionWindow) returns()
func (_DataMarketContract *DataMarketContractTransactor) UpdateSnapshotSubmissionWindow(opts *bind.TransactOpts, newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.contract.Transact(opts, "updateSnapshotSubmissionWindow", newsnapshotSubmissionWindow)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0x9b2f89ce.
//
// Solidity: function updateSnapshotSubmissionWindow(uint256 newsnapshotSubmissionWindow) returns()
func (_DataMarketContract *DataMarketContractSession) UpdateSnapshotSubmissionWindow(newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateSnapshotSubmissionWindow(&_DataMarketContract.TransactOpts, newsnapshotSubmissionWindow)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0x9b2f89ce.
//
// Solidity: function updateSnapshotSubmissionWindow(uint256 newsnapshotSubmissionWindow) returns()
func (_DataMarketContract *DataMarketContractTransactorSession) UpdateSnapshotSubmissionWindow(newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _DataMarketContract.Contract.UpdateSnapshotSubmissionWindow(&_DataMarketContract.TransactOpts, newsnapshotSubmissionWindow)
}

// DataMarketContractAdminsUpdatedIterator is returned from FilterAdminsUpdated and is used to iterate over the raw logs and unpacked data for AdminsUpdated events raised by the DataMarketContract contract.
type DataMarketContractAdminsUpdatedIterator struct {
	Event *DataMarketContractAdminsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractAdminsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractAdminsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractAdminsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractAdminsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractAdminsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractAdminsUpdated represents a AdminsUpdated event raised by the DataMarketContract contract.
type DataMarketContractAdminsUpdated struct {
	AdminAddress common.Address
	Allowed      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAdminsUpdated is a free log retrieval operation binding the contract event 0x915a9250b000555737056ef1e5c2447ba962b934fb8e16b5e3a24db239c2dcf1.
//
// Solidity: event AdminsUpdated(address adminAddress, bool allowed)
func (_DataMarketContract *DataMarketContractFilterer) FilterAdminsUpdated(opts *bind.FilterOpts) (*DataMarketContractAdminsUpdatedIterator, error) {

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "AdminsUpdated")
	if err != nil {
		return nil, err
	}
	return &DataMarketContractAdminsUpdatedIterator{contract: _DataMarketContract.contract, event: "AdminsUpdated", logs: logs, sub: sub}, nil
}

// WatchAdminsUpdated is a free log subscription operation binding the contract event 0x915a9250b000555737056ef1e5c2447ba962b934fb8e16b5e3a24db239c2dcf1.
//
// Solidity: event AdminsUpdated(address adminAddress, bool allowed)
func (_DataMarketContract *DataMarketContractFilterer) WatchAdminsUpdated(opts *bind.WatchOpts, sink chan<- *DataMarketContractAdminsUpdated) (event.Subscription, error) {

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "AdminsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractAdminsUpdated)
				if err := _DataMarketContract.contract.UnpackLog(event, "AdminsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminsUpdated is a log parse operation binding the contract event 0x915a9250b000555737056ef1e5c2447ba962b934fb8e16b5e3a24db239c2dcf1.
//
// Solidity: event AdminsUpdated(address adminAddress, bool allowed)
func (_DataMarketContract *DataMarketContractFilterer) ParseAdminsUpdated(log types.Log) (*DataMarketContractAdminsUpdated, error) {
	event := new(DataMarketContractAdminsUpdated)
	if err := _DataMarketContract.contract.UnpackLog(event, "AdminsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractBatchSubmissionsCompletedIterator is returned from FilterBatchSubmissionsCompleted and is used to iterate over the raw logs and unpacked data for BatchSubmissionsCompleted events raised by the DataMarketContract contract.
type DataMarketContractBatchSubmissionsCompletedIterator struct {
	Event *DataMarketContractBatchSubmissionsCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractBatchSubmissionsCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractBatchSubmissionsCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractBatchSubmissionsCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractBatchSubmissionsCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractBatchSubmissionsCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractBatchSubmissionsCompleted represents a BatchSubmissionsCompleted event raised by the DataMarketContract contract.
type DataMarketContractBatchSubmissionsCompleted struct {
	EpochId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchSubmissionsCompleted is a free log retrieval operation binding the contract event 0x2b2a61c393ac413d7f99c0a6447e48e95915d95e379bfc8e26330c9611c904ce.
//
// Solidity: event BatchSubmissionsCompleted(uint256 indexed epochId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterBatchSubmissionsCompleted(opts *bind.FilterOpts, epochId []*big.Int) (*DataMarketContractBatchSubmissionsCompletedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "BatchSubmissionsCompleted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractBatchSubmissionsCompletedIterator{contract: _DataMarketContract.contract, event: "BatchSubmissionsCompleted", logs: logs, sub: sub}, nil
}

// WatchBatchSubmissionsCompleted is a free log subscription operation binding the contract event 0x2b2a61c393ac413d7f99c0a6447e48e95915d95e379bfc8e26330c9611c904ce.
//
// Solidity: event BatchSubmissionsCompleted(uint256 indexed epochId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchBatchSubmissionsCompleted(opts *bind.WatchOpts, sink chan<- *DataMarketContractBatchSubmissionsCompleted, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "BatchSubmissionsCompleted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractBatchSubmissionsCompleted)
				if err := _DataMarketContract.contract.UnpackLog(event, "BatchSubmissionsCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchSubmissionsCompleted is a log parse operation binding the contract event 0x2b2a61c393ac413d7f99c0a6447e48e95915d95e379bfc8e26330c9611c904ce.
//
// Solidity: event BatchSubmissionsCompleted(uint256 indexed epochId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseBatchSubmissionsCompleted(log types.Log) (*DataMarketContractBatchSubmissionsCompleted, error) {
	event := new(DataMarketContractBatchSubmissionsCompleted)
	if err := _DataMarketContract.contract.UnpackLog(event, "BatchSubmissionsCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractDailyTaskCompletedEventIterator is returned from FilterDailyTaskCompletedEvent and is used to iterate over the raw logs and unpacked data for DailyTaskCompletedEvent events raised by the DataMarketContract contract.
type DataMarketContractDailyTaskCompletedEventIterator struct {
	Event *DataMarketContractDailyTaskCompletedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractDailyTaskCompletedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractDailyTaskCompletedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractDailyTaskCompletedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractDailyTaskCompletedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractDailyTaskCompletedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractDailyTaskCompletedEvent represents a DailyTaskCompletedEvent event raised by the DataMarketContract contract.
type DataMarketContractDailyTaskCompletedEvent struct {
	SnapshotterAddress common.Address
	SlotId             *big.Int
	DayId              *big.Int
	Timestamp          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDailyTaskCompletedEvent is a free log retrieval operation binding the contract event 0x34c900c4105cef3bd58c4b7d2b6fe54f1f64845d5bd5ed2e2e92b52aed2d58ae.
//
// Solidity: event DailyTaskCompletedEvent(address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterDailyTaskCompletedEvent(opts *bind.FilterOpts) (*DataMarketContractDailyTaskCompletedEventIterator, error) {

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "DailyTaskCompletedEvent")
	if err != nil {
		return nil, err
	}
	return &DataMarketContractDailyTaskCompletedEventIterator{contract: _DataMarketContract.contract, event: "DailyTaskCompletedEvent", logs: logs, sub: sub}, nil
}

// WatchDailyTaskCompletedEvent is a free log subscription operation binding the contract event 0x34c900c4105cef3bd58c4b7d2b6fe54f1f64845d5bd5ed2e2e92b52aed2d58ae.
//
// Solidity: event DailyTaskCompletedEvent(address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchDailyTaskCompletedEvent(opts *bind.WatchOpts, sink chan<- *DataMarketContractDailyTaskCompletedEvent) (event.Subscription, error) {

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "DailyTaskCompletedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractDailyTaskCompletedEvent)
				if err := _DataMarketContract.contract.UnpackLog(event, "DailyTaskCompletedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDailyTaskCompletedEvent is a log parse operation binding the contract event 0x34c900c4105cef3bd58c4b7d2b6fe54f1f64845d5bd5ed2e2e92b52aed2d58ae.
//
// Solidity: event DailyTaskCompletedEvent(address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseDailyTaskCompletedEvent(log types.Log) (*DataMarketContractDailyTaskCompletedEvent, error) {
	event := new(DataMarketContractDailyTaskCompletedEvent)
	if err := _DataMarketContract.contract.UnpackLog(event, "DailyTaskCompletedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractDayStartedEventIterator is returned from FilterDayStartedEvent and is used to iterate over the raw logs and unpacked data for DayStartedEvent events raised by the DataMarketContract contract.
type DataMarketContractDayStartedEventIterator struct {
	Event *DataMarketContractDayStartedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractDayStartedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractDayStartedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractDayStartedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractDayStartedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractDayStartedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractDayStartedEvent represents a DayStartedEvent event raised by the DataMarketContract contract.
type DataMarketContractDayStartedEvent struct {
	DayId     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDayStartedEvent is a free log retrieval operation binding the contract event 0xf391963fbbcec4cbb1f4a6c915c531364db26c103d31434223c3bddb703c94fe.
//
// Solidity: event DayStartedEvent(uint256 dayId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterDayStartedEvent(opts *bind.FilterOpts) (*DataMarketContractDayStartedEventIterator, error) {

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "DayStartedEvent")
	if err != nil {
		return nil, err
	}
	return &DataMarketContractDayStartedEventIterator{contract: _DataMarketContract.contract, event: "DayStartedEvent", logs: logs, sub: sub}, nil
}

// WatchDayStartedEvent is a free log subscription operation binding the contract event 0xf391963fbbcec4cbb1f4a6c915c531364db26c103d31434223c3bddb703c94fe.
//
// Solidity: event DayStartedEvent(uint256 dayId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchDayStartedEvent(opts *bind.WatchOpts, sink chan<- *DataMarketContractDayStartedEvent) (event.Subscription, error) {

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "DayStartedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractDayStartedEvent)
				if err := _DataMarketContract.contract.UnpackLog(event, "DayStartedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDayStartedEvent is a log parse operation binding the contract event 0xf391963fbbcec4cbb1f4a6c915c531364db26c103d31434223c3bddb703c94fe.
//
// Solidity: event DayStartedEvent(uint256 dayId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseDayStartedEvent(log types.Log) (*DataMarketContractDayStartedEvent, error) {
	event := new(DataMarketContractDayStartedEvent)
	if err := _DataMarketContract.contract.UnpackLog(event, "DayStartedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractDelayedAttestationSubmittedIterator is returned from FilterDelayedAttestationSubmitted and is used to iterate over the raw logs and unpacked data for DelayedAttestationSubmitted events raised by the DataMarketContract contract.
type DataMarketContractDelayedAttestationSubmittedIterator struct {
	Event *DataMarketContractDelayedAttestationSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractDelayedAttestationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractDelayedAttestationSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractDelayedAttestationSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractDelayedAttestationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractDelayedAttestationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractDelayedAttestationSubmitted represents a DelayedAttestationSubmitted event raised by the DataMarketContract contract.
type DataMarketContractDelayedAttestationSubmitted struct {
	BatchCid      string
	EpochId       *big.Int
	Timestamp     *big.Int
	ValidatorAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelayedAttestationSubmitted is a free log retrieval operation binding the contract event 0x67440e4fec48a8284d31238f6c33109a5c9bafc1cebbf916fd31ef6ec9fa95b0.
//
// Solidity: event DelayedAttestationSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_DataMarketContract *DataMarketContractFilterer) FilterDelayedAttestationSubmitted(opts *bind.FilterOpts, epochId []*big.Int, validatorAddr []common.Address) (*DataMarketContractDelayedAttestationSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "DelayedAttestationSubmitted", epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractDelayedAttestationSubmittedIterator{contract: _DataMarketContract.contract, event: "DelayedAttestationSubmitted", logs: logs, sub: sub}, nil
}

// WatchDelayedAttestationSubmitted is a free log subscription operation binding the contract event 0x67440e4fec48a8284d31238f6c33109a5c9bafc1cebbf916fd31ef6ec9fa95b0.
//
// Solidity: event DelayedAttestationSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_DataMarketContract *DataMarketContractFilterer) WatchDelayedAttestationSubmitted(opts *bind.WatchOpts, sink chan<- *DataMarketContractDelayedAttestationSubmitted, epochId []*big.Int, validatorAddr []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "DelayedAttestationSubmitted", epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractDelayedAttestationSubmitted)
				if err := _DataMarketContract.contract.UnpackLog(event, "DelayedAttestationSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedAttestationSubmitted is a log parse operation binding the contract event 0x67440e4fec48a8284d31238f6c33109a5c9bafc1cebbf916fd31ef6ec9fa95b0.
//
// Solidity: event DelayedAttestationSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_DataMarketContract *DataMarketContractFilterer) ParseDelayedAttestationSubmitted(log types.Log) (*DataMarketContractDelayedAttestationSubmitted, error) {
	event := new(DataMarketContractDelayedAttestationSubmitted)
	if err := _DataMarketContract.contract.UnpackLog(event, "DelayedAttestationSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractDelayedBatchSubmittedIterator is returned from FilterDelayedBatchSubmitted and is used to iterate over the raw logs and unpacked data for DelayedBatchSubmitted events raised by the DataMarketContract contract.
type DataMarketContractDelayedBatchSubmittedIterator struct {
	Event *DataMarketContractDelayedBatchSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractDelayedBatchSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractDelayedBatchSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractDelayedBatchSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractDelayedBatchSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractDelayedBatchSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractDelayedBatchSubmitted represents a DelayedBatchSubmitted event raised by the DataMarketContract contract.
type DataMarketContractDelayedBatchSubmitted struct {
	BatchCid  string
	EpochId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayedBatchSubmitted is a free log retrieval operation binding the contract event 0xc98d86a4d6dd4592fae76e7bb009e979bd41889243da56718fb1a14ed56be06b.
//
// Solidity: event DelayedBatchSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterDelayedBatchSubmitted(opts *bind.FilterOpts, epochId []*big.Int) (*DataMarketContractDelayedBatchSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "DelayedBatchSubmitted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractDelayedBatchSubmittedIterator{contract: _DataMarketContract.contract, event: "DelayedBatchSubmitted", logs: logs, sub: sub}, nil
}

// WatchDelayedBatchSubmitted is a free log subscription operation binding the contract event 0xc98d86a4d6dd4592fae76e7bb009e979bd41889243da56718fb1a14ed56be06b.
//
// Solidity: event DelayedBatchSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchDelayedBatchSubmitted(opts *bind.WatchOpts, sink chan<- *DataMarketContractDelayedBatchSubmitted, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "DelayedBatchSubmitted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractDelayedBatchSubmitted)
				if err := _DataMarketContract.contract.UnpackLog(event, "DelayedBatchSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedBatchSubmitted is a log parse operation binding the contract event 0xc98d86a4d6dd4592fae76e7bb009e979bd41889243da56718fb1a14ed56be06b.
//
// Solidity: event DelayedBatchSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseDelayedBatchSubmitted(log types.Log) (*DataMarketContractDelayedBatchSubmitted, error) {
	event := new(DataMarketContractDelayedBatchSubmitted)
	if err := _DataMarketContract.contract.UnpackLog(event, "DelayedBatchSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractEpochReleasedIterator is returned from FilterEpochReleased and is used to iterate over the raw logs and unpacked data for EpochReleased events raised by the DataMarketContract contract.
type DataMarketContractEpochReleasedIterator struct {
	Event *DataMarketContractEpochReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractEpochReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractEpochReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractEpochReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractEpochReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractEpochReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractEpochReleased represents a EpochReleased event raised by the DataMarketContract contract.
type DataMarketContractEpochReleased struct {
	EpochId   *big.Int
	Begin     *big.Int
	End       *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEpochReleased is a free log retrieval operation binding the contract event 0x108f87075a74f81fa2271fdf9fc0883a1811431182601fc65d24513970336640.
//
// Solidity: event EpochReleased(uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterEpochReleased(opts *bind.FilterOpts, epochId []*big.Int) (*DataMarketContractEpochReleasedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "EpochReleased", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractEpochReleasedIterator{contract: _DataMarketContract.contract, event: "EpochReleased", logs: logs, sub: sub}, nil
}

// WatchEpochReleased is a free log subscription operation binding the contract event 0x108f87075a74f81fa2271fdf9fc0883a1811431182601fc65d24513970336640.
//
// Solidity: event EpochReleased(uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchEpochReleased(opts *bind.WatchOpts, sink chan<- *DataMarketContractEpochReleased, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "EpochReleased", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractEpochReleased)
				if err := _DataMarketContract.contract.UnpackLog(event, "EpochReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEpochReleased is a log parse operation binding the contract event 0x108f87075a74f81fa2271fdf9fc0883a1811431182601fc65d24513970336640.
//
// Solidity: event EpochReleased(uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseEpochReleased(log types.Log) (*DataMarketContractEpochReleased, error) {
	event := new(DataMarketContractEpochReleased)
	if err := _DataMarketContract.contract.UnpackLog(event, "EpochReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DataMarketContract contract.
type DataMarketContractOwnershipTransferredIterator struct {
	Event *DataMarketContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractOwnershipTransferred represents a OwnershipTransferred event raised by the DataMarketContract contract.
type DataMarketContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DataMarketContract *DataMarketContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DataMarketContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractOwnershipTransferredIterator{contract: _DataMarketContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DataMarketContract *DataMarketContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DataMarketContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractOwnershipTransferred)
				if err := _DataMarketContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DataMarketContract *DataMarketContractFilterer) ParseOwnershipTransferred(log types.Log) (*DataMarketContractOwnershipTransferred, error) {
	event := new(DataMarketContractOwnershipTransferred)
	if err := _DataMarketContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractRewardsDistributedEventIterator is returned from FilterRewardsDistributedEvent and is used to iterate over the raw logs and unpacked data for RewardsDistributedEvent events raised by the DataMarketContract contract.
type DataMarketContractRewardsDistributedEventIterator struct {
	Event *DataMarketContractRewardsDistributedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractRewardsDistributedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractRewardsDistributedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractRewardsDistributedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractRewardsDistributedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractRewardsDistributedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractRewardsDistributedEvent represents a RewardsDistributedEvent event raised by the DataMarketContract contract.
type DataMarketContractRewardsDistributedEvent struct {
	SnapshotterAddress common.Address
	SlotId             *big.Int
	DayId              *big.Int
	RewardPoints       *big.Int
	Timestamp          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributedEvent is a free log retrieval operation binding the contract event 0x18ea31eebfbd96f0fbddb48c5c9beca4b3952c503de5436b956c4e6168178766.
//
// Solidity: event RewardsDistributedEvent(address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 rewardPoints, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterRewardsDistributedEvent(opts *bind.FilterOpts) (*DataMarketContractRewardsDistributedEventIterator, error) {

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "RewardsDistributedEvent")
	if err != nil {
		return nil, err
	}
	return &DataMarketContractRewardsDistributedEventIterator{contract: _DataMarketContract.contract, event: "RewardsDistributedEvent", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributedEvent is a free log subscription operation binding the contract event 0x18ea31eebfbd96f0fbddb48c5c9beca4b3952c503de5436b956c4e6168178766.
//
// Solidity: event RewardsDistributedEvent(address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 rewardPoints, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchRewardsDistributedEvent(opts *bind.WatchOpts, sink chan<- *DataMarketContractRewardsDistributedEvent) (event.Subscription, error) {

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "RewardsDistributedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractRewardsDistributedEvent)
				if err := _DataMarketContract.contract.UnpackLog(event, "RewardsDistributedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsDistributedEvent is a log parse operation binding the contract event 0x18ea31eebfbd96f0fbddb48c5c9beca4b3952c503de5436b956c4e6168178766.
//
// Solidity: event RewardsDistributedEvent(address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 rewardPoints, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseRewardsDistributedEvent(log types.Log) (*DataMarketContractRewardsDistributedEvent, error) {
	event := new(DataMarketContractRewardsDistributedEvent)
	if err := _DataMarketContract.contract.UnpackLog(event, "RewardsDistributedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractSequencersUpdatedIterator is returned from FilterSequencersUpdated and is used to iterate over the raw logs and unpacked data for SequencersUpdated events raised by the DataMarketContract contract.
type DataMarketContractSequencersUpdatedIterator struct {
	Event *DataMarketContractSequencersUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractSequencersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractSequencersUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractSequencersUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractSequencersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractSequencersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractSequencersUpdated represents a SequencersUpdated event raised by the DataMarketContract contract.
type DataMarketContractSequencersUpdated struct {
	SequencerAddress common.Address
	Allowed          bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSequencersUpdated is a free log retrieval operation binding the contract event 0xe8706f0696b5c674b870163744013b1f5a0e18dfbf77e57997e1ab22148beae0.
//
// Solidity: event SequencersUpdated(address sequencerAddress, bool allowed)
func (_DataMarketContract *DataMarketContractFilterer) FilterSequencersUpdated(opts *bind.FilterOpts) (*DataMarketContractSequencersUpdatedIterator, error) {

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "SequencersUpdated")
	if err != nil {
		return nil, err
	}
	return &DataMarketContractSequencersUpdatedIterator{contract: _DataMarketContract.contract, event: "SequencersUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencersUpdated is a free log subscription operation binding the contract event 0xe8706f0696b5c674b870163744013b1f5a0e18dfbf77e57997e1ab22148beae0.
//
// Solidity: event SequencersUpdated(address sequencerAddress, bool allowed)
func (_DataMarketContract *DataMarketContractFilterer) WatchSequencersUpdated(opts *bind.WatchOpts, sink chan<- *DataMarketContractSequencersUpdated) (event.Subscription, error) {

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "SequencersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractSequencersUpdated)
				if err := _DataMarketContract.contract.UnpackLog(event, "SequencersUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencersUpdated is a log parse operation binding the contract event 0xe8706f0696b5c674b870163744013b1f5a0e18dfbf77e57997e1ab22148beae0.
//
// Solidity: event SequencersUpdated(address sequencerAddress, bool allowed)
func (_DataMarketContract *DataMarketContractFilterer) ParseSequencersUpdated(log types.Log) (*DataMarketContractSequencersUpdated, error) {
	event := new(DataMarketContractSequencersUpdated)
	if err := _DataMarketContract.contract.UnpackLog(event, "SequencersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractSnapshotBatchAttestationSubmittedIterator is returned from FilterSnapshotBatchAttestationSubmitted and is used to iterate over the raw logs and unpacked data for SnapshotBatchAttestationSubmitted events raised by the DataMarketContract contract.
type DataMarketContractSnapshotBatchAttestationSubmittedIterator struct {
	Event *DataMarketContractSnapshotBatchAttestationSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractSnapshotBatchAttestationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractSnapshotBatchAttestationSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractSnapshotBatchAttestationSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractSnapshotBatchAttestationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractSnapshotBatchAttestationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractSnapshotBatchAttestationSubmitted represents a SnapshotBatchAttestationSubmitted event raised by the DataMarketContract contract.
type DataMarketContractSnapshotBatchAttestationSubmitted struct {
	BatchCid      string
	EpochId       *big.Int
	Timestamp     *big.Int
	ValidatorAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchAttestationSubmitted is a free log retrieval operation binding the contract event 0xa5a3bc567e22fff65d0ba9bd5b4f331a885e8335b9f9b5d3d06e72c090f31c72.
//
// Solidity: event SnapshotBatchAttestationSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_DataMarketContract *DataMarketContractFilterer) FilterSnapshotBatchAttestationSubmitted(opts *bind.FilterOpts, epochId []*big.Int, validatorAddr []common.Address) (*DataMarketContractSnapshotBatchAttestationSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "SnapshotBatchAttestationSubmitted", epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractSnapshotBatchAttestationSubmittedIterator{contract: _DataMarketContract.contract, event: "SnapshotBatchAttestationSubmitted", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchAttestationSubmitted is a free log subscription operation binding the contract event 0xa5a3bc567e22fff65d0ba9bd5b4f331a885e8335b9f9b5d3d06e72c090f31c72.
//
// Solidity: event SnapshotBatchAttestationSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_DataMarketContract *DataMarketContractFilterer) WatchSnapshotBatchAttestationSubmitted(opts *bind.WatchOpts, sink chan<- *DataMarketContractSnapshotBatchAttestationSubmitted, epochId []*big.Int, validatorAddr []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "SnapshotBatchAttestationSubmitted", epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractSnapshotBatchAttestationSubmitted)
				if err := _DataMarketContract.contract.UnpackLog(event, "SnapshotBatchAttestationSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchAttestationSubmitted is a log parse operation binding the contract event 0xa5a3bc567e22fff65d0ba9bd5b4f331a885e8335b9f9b5d3d06e72c090f31c72.
//
// Solidity: event SnapshotBatchAttestationSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_DataMarketContract *DataMarketContractFilterer) ParseSnapshotBatchAttestationSubmitted(log types.Log) (*DataMarketContractSnapshotBatchAttestationSubmitted, error) {
	event := new(DataMarketContractSnapshotBatchAttestationSubmitted)
	if err := _DataMarketContract.contract.UnpackLog(event, "SnapshotBatchAttestationSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractSnapshotBatchFinalizedIterator is returned from FilterSnapshotBatchFinalized and is used to iterate over the raw logs and unpacked data for SnapshotBatchFinalized events raised by the DataMarketContract contract.
type DataMarketContractSnapshotBatchFinalizedIterator struct {
	Event *DataMarketContractSnapshotBatchFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractSnapshotBatchFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractSnapshotBatchFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractSnapshotBatchFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractSnapshotBatchFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractSnapshotBatchFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractSnapshotBatchFinalized represents a SnapshotBatchFinalized event raised by the DataMarketContract contract.
type DataMarketContractSnapshotBatchFinalized struct {
	EpochId   *big.Int
	BatchCid  string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchFinalized is a free log retrieval operation binding the contract event 0xc9446eca8a1bc65d8ed6c807688d74a76dee9a7e98ae1d7372c6216e56ae1d74.
//
// Solidity: event SnapshotBatchFinalized(uint256 indexed epochId, string batchCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterSnapshotBatchFinalized(opts *bind.FilterOpts, epochId []*big.Int) (*DataMarketContractSnapshotBatchFinalizedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "SnapshotBatchFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractSnapshotBatchFinalizedIterator{contract: _DataMarketContract.contract, event: "SnapshotBatchFinalized", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchFinalized is a free log subscription operation binding the contract event 0xc9446eca8a1bc65d8ed6c807688d74a76dee9a7e98ae1d7372c6216e56ae1d74.
//
// Solidity: event SnapshotBatchFinalized(uint256 indexed epochId, string batchCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchSnapshotBatchFinalized(opts *bind.WatchOpts, sink chan<- *DataMarketContractSnapshotBatchFinalized, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "SnapshotBatchFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractSnapshotBatchFinalized)
				if err := _DataMarketContract.contract.UnpackLog(event, "SnapshotBatchFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchFinalized is a log parse operation binding the contract event 0xc9446eca8a1bc65d8ed6c807688d74a76dee9a7e98ae1d7372c6216e56ae1d74.
//
// Solidity: event SnapshotBatchFinalized(uint256 indexed epochId, string batchCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseSnapshotBatchFinalized(log types.Log) (*DataMarketContractSnapshotBatchFinalized, error) {
	event := new(DataMarketContractSnapshotBatchFinalized)
	if err := _DataMarketContract.contract.UnpackLog(event, "SnapshotBatchFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractSnapshotBatchSubmittedIterator is returned from FilterSnapshotBatchSubmitted and is used to iterate over the raw logs and unpacked data for SnapshotBatchSubmitted events raised by the DataMarketContract contract.
type DataMarketContractSnapshotBatchSubmittedIterator struct {
	Event *DataMarketContractSnapshotBatchSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractSnapshotBatchSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractSnapshotBatchSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractSnapshotBatchSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractSnapshotBatchSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractSnapshotBatchSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractSnapshotBatchSubmitted represents a SnapshotBatchSubmitted event raised by the DataMarketContract contract.
type DataMarketContractSnapshotBatchSubmitted struct {
	BatchCid  string
	EpochId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchSubmitted is a free log retrieval operation binding the contract event 0x3151fabdd1bacd216584de61e446f964f73b7864a351548f60ce030cb6242a0d.
//
// Solidity: event SnapshotBatchSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterSnapshotBatchSubmitted(opts *bind.FilterOpts, epochId []*big.Int) (*DataMarketContractSnapshotBatchSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "SnapshotBatchSubmitted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractSnapshotBatchSubmittedIterator{contract: _DataMarketContract.contract, event: "SnapshotBatchSubmitted", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchSubmitted is a free log subscription operation binding the contract event 0x3151fabdd1bacd216584de61e446f964f73b7864a351548f60ce030cb6242a0d.
//
// Solidity: event SnapshotBatchSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchSnapshotBatchSubmitted(opts *bind.WatchOpts, sink chan<- *DataMarketContractSnapshotBatchSubmitted, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "SnapshotBatchSubmitted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractSnapshotBatchSubmitted)
				if err := _DataMarketContract.contract.UnpackLog(event, "SnapshotBatchSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchSubmitted is a log parse operation binding the contract event 0x3151fabdd1bacd216584de61e446f964f73b7864a351548f60ce030cb6242a0d.
//
// Solidity: event SnapshotBatchSubmitted(string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseSnapshotBatchSubmitted(log types.Log) (*DataMarketContractSnapshotBatchSubmitted, error) {
	event := new(DataMarketContractSnapshotBatchSubmitted)
	if err := _DataMarketContract.contract.UnpackLog(event, "SnapshotBatchSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractSnapshotFinalizedIterator is returned from FilterSnapshotFinalized and is used to iterate over the raw logs and unpacked data for SnapshotFinalized events raised by the DataMarketContract contract.
type DataMarketContractSnapshotFinalizedIterator struct {
	Event *DataMarketContractSnapshotFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractSnapshotFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractSnapshotFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractSnapshotFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractSnapshotFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractSnapshotFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractSnapshotFinalized represents a SnapshotFinalized event raised by the DataMarketContract contract.
type DataMarketContractSnapshotFinalized struct {
	EpochId     *big.Int
	EpochEnd    *big.Int
	ProjectId   string
	SnapshotCid string
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSnapshotFinalized is a free log retrieval operation binding the contract event 0xe5231a68c59ef23c90b7da4209eae4c795477f0d5dcfa14a612ea96f69a18e15.
//
// Solidity: event SnapshotFinalized(uint256 indexed epochId, uint256 epochEnd, string projectId, string snapshotCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterSnapshotFinalized(opts *bind.FilterOpts, epochId []*big.Int) (*DataMarketContractSnapshotFinalizedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "SnapshotFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractSnapshotFinalizedIterator{contract: _DataMarketContract.contract, event: "SnapshotFinalized", logs: logs, sub: sub}, nil
}

// WatchSnapshotFinalized is a free log subscription operation binding the contract event 0xe5231a68c59ef23c90b7da4209eae4c795477f0d5dcfa14a612ea96f69a18e15.
//
// Solidity: event SnapshotFinalized(uint256 indexed epochId, uint256 epochEnd, string projectId, string snapshotCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchSnapshotFinalized(opts *bind.WatchOpts, sink chan<- *DataMarketContractSnapshotFinalized, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "SnapshotFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractSnapshotFinalized)
				if err := _DataMarketContract.contract.UnpackLog(event, "SnapshotFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotFinalized is a log parse operation binding the contract event 0xe5231a68c59ef23c90b7da4209eae4c795477f0d5dcfa14a612ea96f69a18e15.
//
// Solidity: event SnapshotFinalized(uint256 indexed epochId, uint256 epochEnd, string projectId, string snapshotCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseSnapshotFinalized(log types.Log) (*DataMarketContractSnapshotFinalized, error) {
	event := new(DataMarketContractSnapshotFinalized)
	if err := _DataMarketContract.contract.UnpackLog(event, "SnapshotFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractTriggerBatchResubmissionIterator is returned from FilterTriggerBatchResubmission and is used to iterate over the raw logs and unpacked data for TriggerBatchResubmission events raised by the DataMarketContract contract.
type DataMarketContractTriggerBatchResubmissionIterator struct {
	Event *DataMarketContractTriggerBatchResubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractTriggerBatchResubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractTriggerBatchResubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractTriggerBatchResubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractTriggerBatchResubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractTriggerBatchResubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractTriggerBatchResubmission represents a TriggerBatchResubmission event raised by the DataMarketContract contract.
type DataMarketContractTriggerBatchResubmission struct {
	EpochId   *big.Int
	BatchCid  string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerBatchResubmission is a free log retrieval operation binding the contract event 0xc64a781b261d854c7e1e20ea4b58d112ca1a78fcc2e01d5c107b7131b3849455.
//
// Solidity: event TriggerBatchResubmission(uint256 indexed epochId, string batchCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterTriggerBatchResubmission(opts *bind.FilterOpts, epochId []*big.Int) (*DataMarketContractTriggerBatchResubmissionIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "TriggerBatchResubmission", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractTriggerBatchResubmissionIterator{contract: _DataMarketContract.contract, event: "TriggerBatchResubmission", logs: logs, sub: sub}, nil
}

// WatchTriggerBatchResubmission is a free log subscription operation binding the contract event 0xc64a781b261d854c7e1e20ea4b58d112ca1a78fcc2e01d5c107b7131b3849455.
//
// Solidity: event TriggerBatchResubmission(uint256 indexed epochId, string batchCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchTriggerBatchResubmission(opts *bind.WatchOpts, sink chan<- *DataMarketContractTriggerBatchResubmission, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "TriggerBatchResubmission", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractTriggerBatchResubmission)
				if err := _DataMarketContract.contract.UnpackLog(event, "TriggerBatchResubmission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTriggerBatchResubmission is a log parse operation binding the contract event 0xc64a781b261d854c7e1e20ea4b58d112ca1a78fcc2e01d5c107b7131b3849455.
//
// Solidity: event TriggerBatchResubmission(uint256 indexed epochId, string batchCid, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseTriggerBatchResubmission(log types.Log) (*DataMarketContractTriggerBatchResubmission, error) {
	event := new(DataMarketContractTriggerBatchResubmission)
	if err := _DataMarketContract.contract.UnpackLog(event, "TriggerBatchResubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractValidatorAttestationsInvalidatedIterator is returned from FilterValidatorAttestationsInvalidated and is used to iterate over the raw logs and unpacked data for ValidatorAttestationsInvalidated events raised by the DataMarketContract contract.
type DataMarketContractValidatorAttestationsInvalidatedIterator struct {
	Event *DataMarketContractValidatorAttestationsInvalidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractValidatorAttestationsInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractValidatorAttestationsInvalidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractValidatorAttestationsInvalidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractValidatorAttestationsInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractValidatorAttestationsInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractValidatorAttestationsInvalidated represents a ValidatorAttestationsInvalidated event raised by the DataMarketContract contract.
type DataMarketContractValidatorAttestationsInvalidated struct {
	EpochId   *big.Int
	BatchCid  string
	Validator common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAttestationsInvalidated is a free log retrieval operation binding the contract event 0x81aaea7f722a0d1ab2ab037c155b0e3b9ccd0fe0c099edaaa325ffd0c3fa928b.
//
// Solidity: event ValidatorAttestationsInvalidated(uint256 indexed epochId, string batchCid, address validator, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) FilterValidatorAttestationsInvalidated(opts *bind.FilterOpts, epochId []*big.Int) (*DataMarketContractValidatorAttestationsInvalidatedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "ValidatorAttestationsInvalidated", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DataMarketContractValidatorAttestationsInvalidatedIterator{contract: _DataMarketContract.contract, event: "ValidatorAttestationsInvalidated", logs: logs, sub: sub}, nil
}

// WatchValidatorAttestationsInvalidated is a free log subscription operation binding the contract event 0x81aaea7f722a0d1ab2ab037c155b0e3b9ccd0fe0c099edaaa325ffd0c3fa928b.
//
// Solidity: event ValidatorAttestationsInvalidated(uint256 indexed epochId, string batchCid, address validator, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) WatchValidatorAttestationsInvalidated(opts *bind.WatchOpts, sink chan<- *DataMarketContractValidatorAttestationsInvalidated, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "ValidatorAttestationsInvalidated", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractValidatorAttestationsInvalidated)
				if err := _DataMarketContract.contract.UnpackLog(event, "ValidatorAttestationsInvalidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorAttestationsInvalidated is a log parse operation binding the contract event 0x81aaea7f722a0d1ab2ab037c155b0e3b9ccd0fe0c099edaaa325ffd0c3fa928b.
//
// Solidity: event ValidatorAttestationsInvalidated(uint256 indexed epochId, string batchCid, address validator, uint256 timestamp)
func (_DataMarketContract *DataMarketContractFilterer) ParseValidatorAttestationsInvalidated(log types.Log) (*DataMarketContractValidatorAttestationsInvalidated, error) {
	event := new(DataMarketContractValidatorAttestationsInvalidated)
	if err := _DataMarketContract.contract.UnpackLog(event, "ValidatorAttestationsInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataMarketContractValidatorsUpdatedIterator is returned from FilterValidatorsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsUpdated events raised by the DataMarketContract contract.
type DataMarketContractValidatorsUpdatedIterator struct {
	Event *DataMarketContractValidatorsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataMarketContractValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataMarketContractValidatorsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataMarketContractValidatorsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataMarketContractValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataMarketContractValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataMarketContractValidatorsUpdated represents a ValidatorsUpdated event raised by the DataMarketContract contract.
type DataMarketContractValidatorsUpdated struct {
	ValidatorAddress common.Address
	Allowed          bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorsUpdated is a free log retrieval operation binding the contract event 0x7f3079c058f3e3dee87048158309898b46e9741ff53b6c7a3afac7c370649afc.
//
// Solidity: event ValidatorsUpdated(address validatorAddress, bool allowed)
func (_DataMarketContract *DataMarketContractFilterer) FilterValidatorsUpdated(opts *bind.FilterOpts) (*DataMarketContractValidatorsUpdatedIterator, error) {

	logs, sub, err := _DataMarketContract.contract.FilterLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return &DataMarketContractValidatorsUpdatedIterator{contract: _DataMarketContract.contract, event: "ValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsUpdated is a free log subscription operation binding the contract event 0x7f3079c058f3e3dee87048158309898b46e9741ff53b6c7a3afac7c370649afc.
//
// Solidity: event ValidatorsUpdated(address validatorAddress, bool allowed)
func (_DataMarketContract *DataMarketContractFilterer) WatchValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *DataMarketContractValidatorsUpdated) (event.Subscription, error) {

	logs, sub, err := _DataMarketContract.contract.WatchLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataMarketContractValidatorsUpdated)
				if err := _DataMarketContract.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorsUpdated is a log parse operation binding the contract event 0x7f3079c058f3e3dee87048158309898b46e9741ff53b6c7a3afac7c370649afc.
//
// Solidity: event ValidatorsUpdated(address validatorAddress, bool allowed)
func (_DataMarketContract *DataMarketContractFilterer) ParseValidatorsUpdated(log types.Log) (*DataMarketContractValidatorsUpdated, error) {
	event := new(DataMarketContractValidatorsUpdated)
	if err := _DataMarketContract.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
