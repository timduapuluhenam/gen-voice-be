package notifications

import (
	"time"

	"genVoice/app/middlewares"
	"genVoice/business/activities"
)

type NotifService struct {
	repository     Repository
	activityRepo   activities.Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewNotifService(repo Repository, activityRepo activities.Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &NotifService{
		repository:     repo,
		activityRepo:   activityRepo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (servUser *NotifService) GetNotif(status, signature_key string) error {
	servUser.repository.GetNotif(status, signature_key)
	_, errActivity := servUser.activityRepo.CreateActivity(&activities.Domain{
		UserID:   1,
		Activity: "tesmasuk"})

	if errActivity != nil {
		return nil
	}

	return nil
}
