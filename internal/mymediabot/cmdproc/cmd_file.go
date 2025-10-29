package cmdproc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	m "github.com/devldavydov/mymedia/internal/common/messages"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v4"
)

func (r *CmdProcessor) fileListCommand(userID int64) []CmdResponse {
	entries, err := os.ReadDir(r.storageDir)
	if err != nil {
		r.logger.Error(
			"failed to get files list",
			zap.Int64("userID", userID),
			zap.Error(err))
		return NewSingleCmdResponse(m.MsgErrInternal)
	}

	if len(entries) == 0 {
		return NewSingleCmdResponse(m.MsgErrEmptyResult)
	}

	lstFiles := make([]string, 0, len(entries))
	for _, entry := range entries {
		lstFiles = append(lstFiles, entry.Name())
	}
	lstFiles = append(lstFiles, "", fmt.Sprintf(m.MsgFileTotal, len(lstFiles)))

	return NewSingleCmdResponse(strings.Join(lstFiles, "\n"))
}

func (r *CmdProcessor) fileRmCommand(userID int64, pattern string) []CmdResponse {
	entries, err := os.ReadDir(r.storageDir)
	if err != nil {
		r.logger.Error(
			"failed to get files list",
			zap.Int64("userID", userID),
			zap.Error(err))
		return NewSingleCmdResponse(m.MsgErrInternal)
	}

	var lstErr []string
	var totalDeleted int64
	for _, entry := range entries {
		fileName := entry.Name()

		if pattern != "" && !strings.Contains(fileName, pattern) {
			continue
		}

		if err := os.Remove(filepath.Join(r.storageDir, fileName)); err != nil {
			r.logger.Error(
				"failed to delete file",
				zap.String("fileName", fileName),
				zap.Int64("userID", userID),
				zap.Error(err),
			)
			lstErr = append(lstErr, fmt.Sprintf(m.MsgFileDeleteErr, fileName, err))
			continue
		}

		totalDeleted += 1
	}

	if len(lstErr) != 0 {
		resp := make([]CmdResponse, 0, len(lstErr))
		for _, fErr := range lstErr {
			resp = append(resp, NewCmdResponse(fErr))
			return resp
		}
	}

	return NewSingleCmdResponse(fmt.Sprintf(m.MsgFileDeleted, totalDeleted))
}

func (r *CmdProcessor) fileDownloadCommand(userID int64, pattern string) []CmdResponse {
	entries, err := os.ReadDir(r.storageDir)
	if err != nil {
		r.logger.Error(
			"failed to get files list",
			zap.Int64("userID", userID),
			zap.Error(err))
		return NewSingleCmdResponse(m.MsgErrInternal)
	}

	type downloadFile struct {
		name string
		path string
	}
	var toDownload []downloadFile

	for _, entry := range entries {
		fileName := entry.Name()

		if pattern != "" && !strings.Contains(fileName, pattern) {
			continue
		}

		toDownload = append(toDownload, downloadFile{
			name: fileName,
			path: filepath.Join(r.storageDir, fileName),
		})
	}

	if len(toDownload) == 0 {
		return NewSingleCmdResponse(m.MsgErrEmptyResult)
	}

	resp := make([]CmdResponse, 0, len(toDownload))
	for _, f := range toDownload {
		resp = append(resp, NewCmdResponse(&tele.Document{
			File:     tele.FromDisk(f.path),
			FileName: f.name,
		}))
	}

	return resp
}
