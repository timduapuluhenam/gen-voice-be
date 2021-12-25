package invoices

import (
	"genVoice/business/invoices"

	"gorm.io/gorm"
)

type MysqlInvoiceRepository struct {
	Conn *gorm.DB
}

func NewMysqlInvoiceRepository(conn *gorm.DB) invoices.Repository {
	return &MysqlInvoiceRepository{
		Conn: conn,
	}
}

func (rep *MysqlInvoiceRepository) CreateInvoiceDetail(invoiceDetailDomain *invoices.DatasDomain) (invoices.DatasDomain, error) {

	invoic := fromInvoiceDomain(invoiceDetailDomain)
	// fmt.Print(invoic)
	resultInvoice := rep.Conn.Create(&invoic)

	id := invoic.ID
	invoiceDetail := fromInvoiceDetailDomain(invoiceDetailDomain, id)
	// fmt.Print(invoiceDetail)
	resultInvoiceDetail := rep.Conn.Create(&invoiceDetail)

	if resultInvoice.Error != nil {
		return invoices.DatasDomain{}, resultInvoice.Error
	}
	if resultInvoiceDetail.Error != nil {
		return invoices.DatasDomain{}, resultInvoiceDetail.Error
	}
	// fmt.Print("mysql go invoice nih", invoic, invoiceDetail)
	return toListDomain(invoic, invoiceDetail), nil

}
