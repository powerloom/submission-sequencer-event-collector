package prost

import (
	"math/big"
)

// EpochMarkerDetails contains information about epoch release and submission limit blocks
type EpochMarkerDetails struct {
	EpochReleaseBlockNumber    int64
	SubmissionLimitBlockNumber int64
}

// BatchDetails contains information about a submitted batch
type BatchDetails struct {
	DataMarketAddress string
	BatchCID          string
	EpochID           *big.Int
}

// SubmissionDetails contains information about submissions for a batch
type SubmissionDetails struct {
	DataMarketAddress string
	EpochID           *big.Int
	BatchID           int
	Batch             map[string][]string // ProjectID -> SubmissionKeys
}

// DayTransitionEpochInfo contains information about day transitions
type DayTransitionEpochInfo struct {
	LastKnownDay string
	CurrentEpoch int64
	BufferEpoch  int64
}

// RelayerRequestBody contains information for relayer requests
type RelayerRequestBody struct {
	DataMarketAddress  string     `json:"dataMarketAddress"`
	SlotIDs            []*big.Int `json:"slotIDs"`
	Submissions        []*big.Int `json:"submissions"`
	CurrentDay         *big.Int   `json:"currentDay"`
	EligibleNodesCount int64      `json:"eligibleNodesCount"`
}
