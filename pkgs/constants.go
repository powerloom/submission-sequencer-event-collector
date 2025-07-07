package pkgs

// Process Name Constants
// process : identifier
const (
	StartFetchingBlocks             = "SequencerEventCollector: StartFetchingBlocks"
	ProcessEvents                   = "SequencerEventCollector: ProcessEvents"
	ProcessMarketData               = "SequencerEventCollector: ProcessMarketData"
	ConstructProjectMap             = "SequencerEventCollector: ConstructProjectMap"
	SendSubmissionBatchSize         = "SequencerEventCollector: SendSubmissionBatchSize"
	UpdateSlotSubmissionCount       = "SequencerEventCollector: UpdateSlotSubmissionCount"
	SendEligibleNodesCount          = "SequencerEventCollector: SendEligibleNodesCount"
	SendUpdateRewardsToRelayer      = "SequencerEventCollector: SendUpdateRewardsToRelayer"
	HandleDayTransition             = "SequencerEventCollector: HandleDayTransition"
	TriggerBatchPreparation         = "SequencerEventCollector: TriggerBatchPreparation"
	CheckAndTriggerBatchPreparation = "SequencerEventCollector: CheckAndTriggerBatchPreparation"
)

// General Key Constants
const (
	CurrentEpoch                     = "CurrentEpochID"
	CurrentDayKey                    = "CurrentDayKey"
	TotalNodesCount                  = "TotalNodesCount"
	DaySizeTableKey                  = "DaySizeTableKey"
	BatchCountKey                    = "BatchCountKey"
	CollectorKey                     = "SnapshotCollector"
	EpochMarkerSetKey                = "EpochMarkerSetKey"
	DayRolloverEpochMarkerSetKey     = "DayRolloverEpochMarkerSetKey"
	DayRolloverEpochMarkerDetailsKey = "DayRolloverEpochMarkerDetailsKey"
	EpochMarkerDetailsKey            = "EpochMarkerDetailsKey"
	BlockHashByNumberKey             = "BlockHashByNumberKey"
	SubmissionLimitTableKey          = "SubmissionLimitTable"
	BatchSubmissionsKey              = "BatchSubmissionsKey"
	SlotSubmissionsKey               = "SlotSubmissionsKey"
	EligibleSlotSubmissionsKey       = "EligibleSlotSubmissionsKey"
	EligibleNodeByDayKey             = "EligibleNodeByDayKey"
	LastKnownDayKey                  = "LastKnownDayKey"
	DailySnapshotQuotaTableKey       = "DailySnapshotQuotaTableKey"
	EpochSubmissionsKey              = "EpochSubmissionsKey"
	EligibleSlotSubmissionByEpochKey = "EligibleSlotSubmissionByEpochKey"
	DiscardedSubmissionKey           = "DiscardedSubmissionKey"
	EpochsInADay                     = "EpochsInADay"
	ZeroCountKey                     = "ZeroCountKey"
	LastSimulatedSubmissionKey       = "LastSimulatedSubmissionKey"
	LastSnapshotSubmissionKey        = "LastSnapshotSubmissionKey"
	ActiveSnapshottersForEpoch       = "ActiveSnapshottersForEpoch"
)
