package users

import (
	"genVoice/business/users"
	"time"
)

type Users struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Username  string `gorm:"unique"`
	Password  string
	Email     string `gorm:"unique"`
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(user Users) users.Domain {
	return users.Domain{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) Users {
	return Users{
		ID:        domain.ID,
		Name:      domain.Name,
		Username:  domain.Username,
		Password:  domain.Password,
		Email:     domain.Email,
		Address:   domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
