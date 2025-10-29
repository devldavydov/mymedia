//go:generate go run ./gen/gen.go -in commands.yaml -out cmdproc_generated.go
package cmdproc

import (
	"fmt"
	"path/filepath"
	"strings"

	m "github.com/devldavydov/mymedia/internal/common/messages"
	"github.com/google/uuid"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v4"
)

var optsHTML = &tele.SendOptions{ParseMode: tele.ModeHTML}

type CmdProcessor struct {
	storageDir string
	logger     *zap.Logger
	debugMode  bool
}

func NewCmdProcessor(storageDir string, debugMode bool, logger *zap.Logger) *CmdProcessor {
	return &CmdProcessor{storageDir: storageDir, debugMode: debugMode, logger: logger}
}

func (r *CmdProcessor) Stop() {
}

func (r *CmdProcessor) ProcessCmd(c tele.Context, cmd string, userID int64) error {
	return r.process(c, cmd, userID)
}

func (r *CmdProcessor) ProcessDocument(c tele.Context, doc *tele.Document, userID int64) error {
	if doc == nil {
		r.logger.Error("empty document", zap.Int64("userID", userID))
		return c.Send(m.MsgFileEmpty)
	}

	fi, err := c.Bot().FileByID(doc.FileID)
	if err != nil {
		r.logger.Error(
			"get file info error",
			zap.String("fileName", doc.FileName),
			zap.Int64("userID", userID),
			zap.Error(err))
		return c.Send(fmt.Sprintf(m.MsgFileInfoErr, doc.FileName, err))
	}

	localFilePath := filepath.Join(r.storageDir, getLocalFileName(doc.FileName))
	if err = c.Bot().Download(&fi, localFilePath); err != nil {
		r.logger.Error(
			"file upload error",
			zap.String("fileName", doc.FileName),
			zap.Int64("userID", userID),
			zap.Error(err))
		return c.Send(fmt.Sprintf(m.MsgFileUploadErr, doc.FileName, err))
	}

	return c.Send(fmt.Sprintf(m.MsgFileUploaded, doc.FileName))
}

func getLocalFileName(srcFileName string) string {
	localName := strings.ReplaceAll(srcFileName, " ", "_")
	localName = uuid.New().String()[:8] + "_" + localName
	return localName
}

type CmdResponse struct {
	what any
	opts []any
}

func NewCmdResponse(what any, opts ...any) CmdResponse {
	return CmdResponse{what: what, opts: opts}
}

func NewSingleCmdResponse(what any, opts ...any) []CmdResponse {
	return []CmdResponse{
		{what: what, opts: opts},
	}
}
