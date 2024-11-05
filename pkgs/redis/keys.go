package redis

import (
	"fmt"
	"strings"
	"submission-sequencer-collector/pkgs"
)

func GetDaySizeTableKey() string {
	return pkgs.DaySizeTableKey
}

func GetSubmissionLimitTableKey() string {
	return pkgs.SubmissionLimitTableKey
}

func GetCurrentDayKey(dataMarketAddress string) string {
	return fmt.Sprintf("%s.%s", pkgs.CurrentDayKey, strings.ToLower(dataMarketAddress))
}

func EpochMarkerSet(dataMarketAddress string) string {
	return fmt.Sprintf("%s.%s", pkgs.EpochMarkerSetKey, strings.ToLower(dataMarketAddress))
}

func EpochMarkerDetails(dataMarketAddress, epochID string) string {
	return fmt.Sprintf("%s.%s.%s", pkgs.EpochMarkerDetailsKey, strings.ToLower(dataMarketAddress), epochID)
}

func SubmissionSetByHeaderKey(dataMarketAddress string, epoch uint64, header string) string {
	return fmt.Sprintf("%s.%s.%d.%s", pkgs.CollectorKey, strings.ToLower(dataMarketAddress), epoch, header)
}

func SlotSubmissionKey(dataMarketAddress string, slotID, currentDay string) string {
	return fmt.Sprintf("%s.%s.%s.%s", pkgs.SlotSubmissionsKey, strings.ToLower(dataMarketAddress), currentDay, slotID)
}

func BlockHashByNumber(blockNum int64) string {
	return fmt.Sprintf("%s.%d", pkgs.BlockHashByNumberKey, blockNum)
}
