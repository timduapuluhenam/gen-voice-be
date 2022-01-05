package activities

import (
	"genVoice/business/activities"

	"gorm.io/gorm"
)

type MySqlActivityRepository struct {
	Conn *gorm.DB
}

func NewMysqlActivityRepository(conn *gorm.DB) activities.Repository {
	return &MySqlActivityRepository{
		Conn: conn,
	}
}

func (rep *MySqlActivityRepository) CreateActivity(activityDomain *activities.Domain) (activities.Domain, error) {
	activity := fromDomain(*activityDomain)
	result := rep.Conn.Create(&activity)

	if result.Error != nil {
		return activities.Domain{}, result.Error
	}

	return toDomain(activity), nil
}
