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
type UsersUpdate struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
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

func (req *UsersUpdate) ToUpdateDomain() *users.UpdateDomain {
	return &users.UpdateDomain{
		ID:       req.ID,
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
		Address:  req.Address,
	}
}
