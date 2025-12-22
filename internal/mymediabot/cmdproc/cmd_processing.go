package cmdproc

import (
	"fmt"
	"os"
	"strings"

	"github.com/devldavydov/mymedia/internal/common/exif"
	m "github.com/devldavydov/mymedia/internal/common/messages"
	"go.uber.org/zap"
)

func (r *CmdProcessor) prcExifRename(userID int64, pattern string) []CmdResponse {
	entries, err := os.ReadDir(r.storageDir)
	if err != nil {
		r.logger.Error(
			"failed to get files list",
			zap.Int64("userID", userID),
			zap.Error(err))
		return NewSingleCmdResponse(m.MsgErrInternal)
	}

	var lstErr []string
	var totalRenamed int64
	for _, entry := range entries {
		fileName := entry.Name()

		if pattern != "" && !strings.Contains(fileName, pattern) {
			continue
		}

		if err := exif.RenameFile(r.storageDir, fileName); err != nil {
			lstErr = append(lstErr, fmt.Sprintf(m.MsgFileRenameErr, fileName, err))
			r.logger.Error(
				"failed to rename file",
				zap.Int64("userID", userID),
				zap.String("fileName", fileName),
				zap.Error(err),
			)
			continue
		}

		totalRenamed += 1
	}

	if len(lstErr) != 0 {
		resp := make([]CmdResponse, 0, len(lstErr))
		for _, fErr := range lstErr {
			resp = append(resp, NewCmdResponse(fErr))
			return resp
		}
	}

	return NewSingleCmdResponse(fmt.Sprintf(m.MsgFileRenamed, totalRenamed))
}
