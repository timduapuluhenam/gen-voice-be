package activities

import (
	"time"
)

type Domain struct {
	ID        int
	UserID    int
	Activity  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	CreateActivity(domain *Domain) (Domain, error)
}

type Repository interface {
	CreateActivity(domain *Domain) (Domain, error)
}
