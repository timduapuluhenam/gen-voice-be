package users

import (
	"fmt"
	"genVoice/business/invoices"
	"genVoice/business/notifications"
	invRepo "genVoice/drivers/databases/invoices"
	"strconv"
	"time"

	emailSucces "genVoice/helper/mail/mailSuccess"

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

	thn, bln, dy := time.Now().Date()
	jm, mnt, dtk := time.Now().Clock()
	start := time.Now().UnixNano()
	sts := strconv.Itoa(int(start))
	mls := sts[len(sts)-6:]
	time := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%s+07:00", thn, bln, dy, jm, mnt, dtk, mls)
	if status == "settlement" {
		rep.Conn.Find(&invoiceDetail, "signature_key = ? ", signature_key)
		rep.Conn.Find(&invoice, "id = ? ", invoiceDetail.EventID)
		rep.Conn.Model(&InvoiceDetail{}).Where("signature_key = ? ", signature_key).Update("status", "Telah Dibayar").Update("updated_at", time)
		if invoiceDetail.Email != "" {
			emailSucces.Email(invoiceDetail.Email, invoiceDetail.Name, strconv.Itoa(invoiceDetail.Amount), invoice.Name)
		}
	}
	return nil

}

func (rep *MysqlNotifRepository) GetUserBySignature(signature_key string) (invoices.InvoiceDetailDomain, int, error) {
	invoiceDetail := InvoiceDetail{}
	invoice := invRepo.Invoices{}
	result := rep.Conn.Find(&invoiceDetail, "signature_key = ? ", signature_key)
	rep.Conn.Find(&invoice, "id = ? ", invoiceDetail.EventID)
	if result.Error != nil {
		return invoices.InvoiceDetailDomain{}, 0, result.Error
	}
	return invoices.InvoiceDetailDomain{Name: invoiceDetail.Name, Amount: invoiceDetail.Amount, EventID: invoiceDetail.EventID}, invoice.UserID, nil
}
