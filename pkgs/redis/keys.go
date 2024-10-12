package redis

import (
	"fmt"
	"submission-sequencer-collector/pkgs"
)

func SubmissionSetByHeaderKey(epoch uint64, header string) string {
	return fmt.Sprintf("%s.%d.%s", pkgs.CollectorKey, epoch, header)
}
