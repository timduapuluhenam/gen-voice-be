package invoices

import (
	"time"
)

type InvoiceDomain struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InvoiceDetailDomain struct {
	ID        int
	Name      string
	Email     string
	Amount    int
	EventID   string
	Link      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	CreateInvoice(invoicedomain *InvoiceDomain) (InvoiceDomain, error)
	CreateInvoiceDetail(invoiceDetailDomain []*InvoiceDetailDomain) ([]InvoiceDetailDomain, error)
}

type Repository interface {
	CreateInvoice(invoicedomain *InvoiceDomain) (InvoiceDomain, error)
	CreateInvoiceDetail(invoiceDetailDomain []*InvoiceDetailDomain) ([]InvoiceDetailDomain, error)
}
