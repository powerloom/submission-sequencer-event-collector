package redis

import (
	"fmt"
	"submission-sequencer-collector/pkgs"
)

func EpochMarkerSet() string {
	return pkgs.EpochMarkerSetKey
}

func SubmissionSetByHeaderKey(epoch uint64, header string) string {
	return fmt.Sprintf("%s.%d.%s", pkgs.CollectorKey, epoch, header)
}

func BlockHashByNumber(blockNum int64) string {
	return fmt.Sprintf("%s.%d", pkgs.BlockHashByNumberKey, blockNum)
}
