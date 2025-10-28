package cmdproc

import (
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v4"
)

type CmdProcessor struct {
	logger *zap.Logger
}

func NewCmdProcessor(logger *zap.Logger) *CmdProcessor {
	return &CmdProcessor{logger: logger}
}

func (r *CmdProcessor) Stop() {
}

func (r *CmdProcessor) ProcessCmd(c tele.Context, cmd string, userID int64) error {
	return nil
}

func (r *CmdProcessor) ProcessDocument(c tele.Context, doc *tele.Document, userID int64) error {
	return nil
}
