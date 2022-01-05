package users

import (
	"fmt"
	"genVoice/business/invoices"
	"genVoice/business/notifications"

	// emailSucces "genVoice/helper/mail/mailSuccess"
	// "strconv"

	"gorm.io/gorm"
)

type MysqlNotifRepository struct {
	Conn *gorm.DB
}

func NewMysqlNotifRepository(conn *gorm.DB) notifications.Repository {
	return &MysqlNotifRepository{
		Conn: conn,
	}
}

func (rep *MysqlNotifRepository) GetNotif(status, signature_key string) error {
	invoiceDetail := InvoiceDetail{}
	invoice := Invoice{}
	if status == "settlement" {
		rep.Conn.Find(&invoiceDetail, "signature_key = ? ", signature_key)
		rep.Conn.Find(&invoice, "id = ? ", invoiceDetail.EventID)
		rep.Conn.Model(&InvoiceDetail{}).Where("signature_key = ? ", signature_key).Update("status", "Telah Dibayar")
		if invoiceDetail.Email != "" {
			// emailSucces.Email(invoiceDetail.Email, invoiceDetail.Name, strconv.Itoa(invoiceDetail.Amount), invoice.Name)
		}
	}
	fmt.Println(invoiceDetail)
	fmt.Println(invoice)

	// Email(to string, name string, amount string, event string)
	return nil

}

func (rep *MysqlNotifRepository) GetUserBySignature(signature_key string) (invoices.InvoiceDetailDomain, error) {
	invoiceDetail := InvoiceDetail{}
	result := rep.Conn.Find(&invoiceDetail, "signature_key = ? ", signature_key)
	if result.Error != nil {
		return invoices.InvoiceDetailDomain{}, result.Error
	}
	return invoices.InvoiceDetailDomain{Name: invoiceDetail.Name, EventID: invoiceDetail.EventID}, nil
}
