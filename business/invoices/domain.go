package invoices

import (
	"time"
)

type DatasDomain struct {
	DataInvoice   InvoiceDomain
	InvoiceDetail []InvoiceDetailDomain
}
type InvoiceDomain struct {
	ID          int
	UserID      int
	Name        string
	TimeExpired int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type InvoiceDetailDomain struct {
	InvoiceName string

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
	GetAllByUserID(userID int) ([]InvoiceDetailDomain, error)
	DeleteInvoice(invoiceID int) (InvoiceDomain, error)
}

type Repository interface {
	CreateInvoiceDetail(invoiceDetailDomain *DatasDomain) (DatasDomain, error)
	GetAllByUserID(userID int) ([]InvoiceDetailDomain, error)
	DeleteInvoice(invoiceID int) (InvoiceDomain, error)
}
