package invoices

import (
	"time"
)

type DatasDomain struct {
	DataInvoice   InvoiceDomain
	InvoiceDetail []InvoiceDetailDomain
}
type InvoiceDomain struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InvoiceDetailDomain struct {
	ID           int
	Name         string
	Email        string
	Amount       int
	EventID      int
	SignatureKey string
	Link         string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type Service interface {
	CreateInvoiceDetail(invoiceDetailDomain *DatasDomain) (DatasDomain, error)
}

type Repository interface {
	CreateInvoiceDetail(invoiceDetailDomain *DatasDomain) (DatasDomain, error)
}
