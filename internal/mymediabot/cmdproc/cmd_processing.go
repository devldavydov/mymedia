package cmdproc

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"

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

		if err := r.exifRenameFile(fileName); err != nil {
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

func (r *CmdProcessor) exifRenameFile(fileName string) error {
	f, err := os.Open(filepath.Join(r.storageDir, fileName))
	if err != nil {
		return err
	}
	defer f.Close()

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		return err
	}

	tm, _ := x.DateTime()
	dstPath := filepath.Join(r.storageDir, tm.Format("20060102_150405.jpg"))
	duplNumber := 0
	for r.isFileExists(dstPath) {
		duplNumber += 1
		dstPath = filepath.Join(r.storageDir, fmt.Sprintf("%s_%d.jpg", tm.Format("20060102_150405"), duplNumber))
	}

	return os.Rename(filepath.Join(r.storageDir, fileName), dstPath)
}

func (r *CmdProcessor) isFileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}
