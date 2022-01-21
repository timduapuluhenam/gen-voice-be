package users

import (
	"fmt"
	"time"

	"genVoice/app/middlewares"
	"genVoice/business"
	"genVoice/helper/encrypt"
)

type UserService struct {
	repository     Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserService(repo Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &UserService{
		repository:     repo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (servUser *UserService) Login(username, password string) (Domain, error) {
	fmt.Print("service bisnis", username, password)
	if username == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if password == "" {
		return Domain{}, business.ErrEmptyForm
	}

	user, err := servUser.repository.Login(username, password)
	if err != nil {
		return Domain{}, err
	}
	validPass := encrypt.CheckPasswordHash(password, user.Password)
	if !validPass {
		return Domain{}, business.ErrUser
	}
	user.Token = servUser.jwtAuth.GenerateToken(user.ID, "user")
	return user, nil
}

func (servUser *UserService) Register(domain *Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if domain.Username == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if domain.Email == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if domain.Address == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if domain.Password == "" {
		return Domain{}, business.ErrEmptyForm
	}
	encryptedPass, _ := encrypt.HashPassword(domain.Password)
	domain.Password = encryptedPass
	user, err := servUser.repository.Register(domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (servUser *UserService) Update(domain *UpdateDomain) (UpdateDomain, error) {
	if domain.Name == "" {
		return UpdateDomain{}, business.ErrEmptyForm
	}

	if domain.Email == "" {
		return UpdateDomain{}, business.ErrEmptyForm
	}
	if domain.Address == "" {
		return UpdateDomain{}, business.ErrEmptyForm
	}
	if domain.Password == "" {
		return UpdateDomain{}, business.ErrEmptyForm
	}
	encryptedPass, _ := encrypt.HashPassword(domain.Password)
	domain.Password = encryptedPass
	user, err := servUser.repository.Update(domain)

	if err != nil {
		return UpdateDomain{}, err
	}
	return user, nil
}
