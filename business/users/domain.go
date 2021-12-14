package us

import (
	"time"
)

type Domain struct{
	ID 				int
	Username 		string
	Password 		string
	Email 			string
	Token 			string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

type Service interface {
	Register(domain *Domain) (Domain, error)
	Login(username string, password string) (Domain, error)
}

type Repository interface {
	Register(domain *Domain) (Domain, error)
	Login(username string, password string) (Domain, error)
}