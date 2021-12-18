package request

import (
	"genVoice/business/users"
)

// request body for login
type UsersLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// request body for register
type Users struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

// turn register body to domain
func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Address:  req.Address,
	}
}
