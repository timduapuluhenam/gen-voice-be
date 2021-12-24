package drivers

import (
	userDomain "genVoice/business/users"
	userDB "genVoice/drivers/databases/users"

	invoiceDomain "genVoice/business/invoices"
	invoiceDB "genVoice/drivers/databases/invoices"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}

func NewInvoiceRepository(conn *gorm.DB) invoiceDomain.Repository {
	return invoiceDB.NewMysqlInvoiceRepository(conn)
}
