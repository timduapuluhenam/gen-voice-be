package activities

import (
	"time"
)

type ActivityService struct {
	repository     Repository
	contextTimeout time.Duration
}

func NewActivityService(repo Repository, timeout time.Duration) Service {
	return &ActivityService{
		repository:     repo,
		contextTimeout: timeout,
	}
}

func (servActivity *ActivityService) CreateActivity(domain *Domain) (Domain, error) {
	activity, err := servActivity.repository.CreateActivity(domain)

	if err != nil {
		return Domain{}, err
	}
	return activity, nil
}
