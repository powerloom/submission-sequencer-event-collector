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

func SubmissionSetByHeaderKey(epoch uint64, header, dataMarketAddress string) string {
	return fmt.Sprintf("%s.%d.%s.%s", pkgs.CollectorKey, epoch, header, dataMarketAddress)
}

func BlockHashByNumber(blockNum int64) string {
	return fmt.Sprintf("%s.%d", pkgs.BlockHashByNumberKey, blockNum)
}

func TotalSubmissionsCountKey(currentDay, DataMarketAddress string) string {
	return fmt.Sprintf("%s.%s.%s", pkgs.TotalSubmissionCount, currentDay, DataMarketAddress)
}
