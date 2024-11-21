package pkgs

// Process Name Constants
// process : identifier
const (
	StartFetchingBlocks       = "StartFetchingBlocks"
	ProcessEvents             = "ProcessEvents"
	ConstructProjectMap       = "ConstructProjectMap"
	SendSubmissionBatchSize   = "SendSubmissionBatchSize"
	UpdateSlotSubmissionCount = "UpdateSlotSubmissionCount"
)

// General Key Constants
const (
	CurrentDayKey              = "CurrentDayKey"
	DaySizeTableKey            = "DaySizeTableKey"
	DailySnapshotQuotaTableKey = "DailySnapshotQuotaTableKey"
	CollectorKey               = "SnapshotCollector"
	EpochMarkerSetKey          = "EpochMarkerSetKey"
	EpochMarkerDetailsKey      = "EpochMarkerDetailsKey"
	BlockHashByNumberKey       = "BlockHashByNumberKey"
	SubmissionLimitTableKey    = "SubmissionLimitTable"
	BatchSubmissionsKey        = "BatchSubmissionsKey"
	SlotSubmissionsKey         = "SlotSubmissionsKey"
	DaySubmissionsKey          = "DaySubmissionsKey"
	EligibleSlotSubmissionsKey = "EligibleSlotSubmissionsKey"
)
