package pkgs

// Process Name Constants
// process : identifier
const (
	StartFetchingBlocks        = "StartFetchingBlocks"
	ProcessEvents              = "ProcessEvents"
	ConstructProjectMap        = "ConstructProjectMap"
	SendSubmissionBatchSize    = "SendSubmissionBatchSize"
	UpdateSlotSubmissionCount  = "UpdateSlotSubmissionCount"
	SendEligibleNodesCount     = "SendEligibleNodesCount"
	SendUpdateRewardsToRelayer = "SendUpdateRewardsToRelayer"
)

// General Key Constants
const (
	CurrentDayKey                  = "CurrentDayKey"
	DaySizeTableKey                = "DaySizeTableKey"
	CollectorKey                   = "SnapshotCollector"
	EpochMarkerSetKey              = "EpochMarkerSetKey"
	EpochMarkerDetailsKey          = "EpochMarkerDetailsKey"
	BlockHashByNumberKey           = "BlockHashByNumberKey"
	SubmissionLimitTableKey        = "SubmissionLimitTable"
	BatchSubmissionsKey            = "BatchSubmissionsKey"
	SlotSubmissionsKey             = "SlotSubmissionsKey"
	EligibleSlotSubmissionsKey     = "EligibleSlotSubmissionsKey"
	EligibleSlotSubmissionByDayKey = "EligibleSlotSubmissionByDayKey"
)
