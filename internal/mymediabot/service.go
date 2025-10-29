package mymediabot

import (
	"context"
	"fmt"
	"os"

	"github.com/devldavydov/mymedia/internal/mymediabot/cmdproc"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

type Service struct {
	settings *ServiceSettings
	cmdProc  *cmdproc.CmdProcessor
	logger   *zap.Logger
}

func NewService(settings *ServiceSettings, logger *zap.Logger) (*Service, error) {
	srv := &Service{
		settings: settings,
		cmdProc: cmdproc.NewCmdProcessor(
			settings.StorageDir,
			settings.DebugMode,
			logger),
		logger: logger,
	}

	return srv, nil
}

func (r *Service) Run(ctx context.Context) error {
	// Try create storage dir if not exists
	if err := os.MkdirAll(r.settings.StorageDir, 0755); err != nil {
		r.logger.Error(
			"failed to create storage dir",
			zap.String("storage dir", r.settings.StorageDir),
			zap.Error(err))
		return err
	}

	// Create and run telebot
	pref := tele.Settings{
		Token:  r.settings.Token,
		Poller: &tele.LongPoller{Timeout: r.settings.PollTimeOut},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return err
	}

	r.setupRouting(b, r.settings.AllowedUserIDs)
	go b.Start()

	<-ctx.Done()
	b.Stop()
	r.cmdProc.Stop()

	return nil
}

func (r *Service) setupRouting(b *tele.Bot, allowedUserIDs []int64) {
	b.Handle("/start", r.onStart)

	allowedGroup := b.Group()
	allowedGroup.Use(middleware.Whitelist(allowedUserIDs...))
	allowedGroup.Handle(tele.OnText, r.onText)
	allowedGroup.Handle(tele.OnDocument, r.onDocument)
}

func (r *Service) onStart(c tele.Context) error {
	return c.Send(
		fmt.Sprintf(
			"Привет, %s [%d]!\nДобро пожаловать в MyMediaBot!\nОтправь 'h' для помощи",
			c.Sender().Username,
			c.Sender().ID,
		),
	)
}

func (r *Service) onText(c tele.Context) error {
	return r.cmdProc.ProcessCmd(c, c.Text(), c.Sender().ID)
}

func (r *Service) onDocument(c tele.Context) error {
	return r.cmdProc.ProcessDocument(c, c.Message().Document, c.Sender().ID)
}
