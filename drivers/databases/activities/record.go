package activities

import (
	"genVoice/business/activities"
	"time"
)

type Activities struct {
	ID        int `gorm:"primaryKey"`
	Activity  string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(activity Activities) activities.Domain {
	return activities.Domain{
		ID:        activity.ID,
		Activity:  activity.Activity,
		UserID:    activity.UserID,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
	}
}

func fromDomain(domain activities.Domain) Activities {

	return Activities{
		ID:        domain.ID,
		Activity:  domain.Activity,
		UserID:    domain.UserID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
