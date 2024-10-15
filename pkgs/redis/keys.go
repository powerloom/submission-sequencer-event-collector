package redis

import (
	"fmt"
	"submission-sequencer-collector/pkgs"
)

func SubmissionSetByHeaderKey(epoch uint64, header string) string {
	return fmt.Sprintf("%s.%d.%s", pkgs.CollectorKey, epoch, header)
}

func EpochMarkerKey(epochID string) string {
	return fmt.Sprintf("%s.%s", pkgs.EpochMarkerKey, epochID)
}

func BlockNumber(blockNum int64) string {
	return fmt.Sprintf("%s.%d", pkgs.BlockNumberKey, blockNum)
}
