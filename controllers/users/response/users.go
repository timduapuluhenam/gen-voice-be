package response

import (
	"genVoice/business/users"
	"time"
)

type UserRegisterResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainRegister(domain users.Domain) UserRegisterResponse {
	return UserRegisterResponse{
		Message:   "Registration Success",
		ID:        domain.ID,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type UserLoginResponse struct {
	Message  string `json:"message"`
	ID       int    `json:"id"`
	Token    string `json:"token"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
type UpdateUserResponse struct {
	Message  string `json:"message"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FromDomainLogin(domain users.Domain) UserLoginResponse {
	return UserLoginResponse{
		Message:  "Login Success",
		Name:     domain.Name,
		Email:    domain.Email,
		Address:  domain.Address,
		Username: domain.Username,
		ID:       domain.ID,
		Token:    domain.Token,
	}
}

func FromDomainUpdate(domain users.UpdateDomain) UpdateUserResponse {

	// fmt.Println("DOMAIN : ", domain)
	return UpdateUserResponse{
		Message:  "Update Succesz",
		Name:     domain.Name,
		ID:       domain.ID,
		Address:  domain.Address,
		Email:    domain.Email,
		Password: domain.Password,
	}
}
