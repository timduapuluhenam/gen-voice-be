package users

import (
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Username  string
	Password  string
	Email     string
	Address   string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateDomain struct {
	ID        int
	Name      string
	Password  string
	Email     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Register(domain *Domain) (Domain, error)
	Login(username string, password string) (Domain, error)
	Update(domain *UpdateDomain) (UpdateDomain, error)
}

type Repository interface {
	Register(domain *Domain) (Domain, error)
	Login(username string, password string) (Domain, error)
	Update(domain *UpdateDomain) (UpdateDomain, error)
}
