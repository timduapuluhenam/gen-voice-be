package drivers

import (
	userDomain "genVoice/business/users"
	userDB "genVoice/drivers/databases/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}
