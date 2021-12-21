package invoices

import (
	"fmt"
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

func (rep *MysqlInvoiceRepository) CreateInvoice(invoicedomain *invoices.InvoiceDomain) (invoices.InvoiceDomain, error) {
	fmt.Print(invoicedomain)
	invoice := fromInvoiceDomain(*invoicedomain)
	result := rep.Conn.Create(&invoice)

	if result.Error != nil {
		return invoices.InvoiceDomain{}, result.Error
	}

	return toInvoiceDomain(invoice), nil
}

func (rep *MysqlInvoiceRepository) CreateInvoiceDetail(invoiceDetailDomain []*invoices.InvoiceDetailDomain) ([]invoices.InvoiceDetailDomain, error) {
	// fmt.Print(invoiceDetailDomain)
	invoiceDetail := fromInvoiceDetailDomain(invoiceDetailDomain)
	result := rep.Conn.Create(&invoiceDetail)

	if result.Error != nil {
		return []invoices.InvoiceDetailDomain{}, result.Error
	}

	return toListInvoiceDetailDomain(invoiceDetail), nil
}
