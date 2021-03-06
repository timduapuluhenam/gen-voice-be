package drivers

import (
	userDomain "genVoice/business/users"
	userDB "genVoice/drivers/databases/users"

	activityDomain "genVoice/business/activities"
	activityDB "genVoice/drivers/databases/activities"

	invoiceDomain "genVoice/business/invoices"
	invoiceDB "genVoice/drivers/databases/invoices"

	notifDomain "genVoice/business/notifications"
	notifDB "genVoice/drivers/databases/notifications"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}

func NewInvoiceRepository(conn *gorm.DB) invoiceDomain.Repository {
	return invoiceDB.NewMysqlInvoiceRepository(conn)
}

func NewNotifRepository(conn *gorm.DB) notifDomain.Repository {
	return notifDB.NewMysqlNotifRepository(conn)
}

func NewActivityRepository(conn *gorm.DB) activityDomain.Repository {
	return activityDB.NewMysqlActivityRepository(conn)
}
