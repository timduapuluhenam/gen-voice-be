package notifications

import (
	"time"

	"genVoice/app/middlewares"
)

type NotifService struct {
	repository     Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewNotifService(repo Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &NotifService{
		repository:     repo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (servUser *NotifService) GetNotif(status, signature_key string) error {
	servUser.repository.GetNotif(status, signature_key)

	return nil
}
