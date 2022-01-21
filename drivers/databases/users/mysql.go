package users

import (
	"fmt"
	"genVoice/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(username string, password string) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "username = ?", username)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return toDomain(user), nil
}

func (rep *MysqlUserRepository) Register(domain *users.Domain) (users.Domain, error) {
	fmt.Print(domain)
	user := fromDomain(*domain)
	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return toDomain(user), nil
}

func (rep *MysqlUserRepository) Update(domain *users.UpdateDomain) (users.UpdateDomain, error) {
	user := fromUpdateDomain(*domain)
	rep.Conn.Model(&Users{}).Where("id = ?", user.ID).Updates(user)
	return *domain, nil
}
