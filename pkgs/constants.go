package pkgs

// Process Name Constants
// process : identifier
const (
	StartFetchingBlocks        = "SequencerEventCollector: StartFetchingBlocks"
	ProcessEvents              = "SequencerEventCollector: ProcessEvents"
	ConstructProjectMap        = "SequencerEventCollector: ConstructProjectMap"
	SendSubmissionBatchSize    = "SequencerEventCollector: SendSubmissionBatchSize"
	UpdateSlotSubmissionCount  = "SequencerEventCollector: UpdateSlotSubmissionCount"
	SendEligibleNodesCount     = "SequencerEventCollector: SendEligibleNodesCount"
	SendUpdateRewardsToRelayer = "SequencerEventCollector: SendUpdateRewardsToRelayer"
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
	LastKnownDayKey                = "LastKnownDayKey"
	DailySnapshotQuotaTableKey     = "DailySnapshotQuotaTableKey"
)
