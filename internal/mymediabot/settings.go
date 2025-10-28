package mymediabot

import "time"

type ServiceSettings struct {
	Token          string
	PollTimeOut    time.Duration
	BuildCommit    string
	AllowedUserIDs []int64
	DebugMode      bool
}

func NewServiceSettings(
	token string,
	pollTimeout time.Duration,
	allowedUserIDs []int64,
	buildVersion string,
	debugMode bool) (*ServiceSettings, error) {

	return &ServiceSettings{
		Token:          token,
		PollTimeOut:    pollTimeout,
		BuildCommit:    buildVersion,
		AllowedUserIDs: allowedUserIDs,
		DebugMode:      debugMode,
	}, nil
}
