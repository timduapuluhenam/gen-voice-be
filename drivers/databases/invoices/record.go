package invoices

import (
	"genVoice/business/invoices"
	"time"
)

type Invoices struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toInvoiceDomain(invoice Invoices) invoices.InvoiceDomain {
	return invoices.InvoiceDomain{
		ID:        invoice.ID,
		Name:      invoice.Name,
		UserID:    invoice.UserID,
		CreatedAt: invoice.CreatedAt,
		UpdatedAt: invoice.UpdatedAt,
	}
}

func fromInvoiceDomain(domain invoices.InvoiceDomain) Invoices {
	return Invoices{
		ID:        domain.ID,
		Name:      domain.Name,
		UserID:    domain.UserID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type InvoiceDetail struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Amount    int
	EventID   string
	Link      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toInvoiceDetailDomain(invoice InvoiceDetail) invoices.InvoiceDetailDomain {

	return invoices.InvoiceDetailDomain{
		ID:        invoice.ID,
		Name:      invoice.Name,
		Email:     invoice.Email,
		Amount:    invoice.Amount,
		EventID:   invoice.EventID,
		Link:      invoice.Link,
		Status:    invoice.Status,
		CreatedAt: invoice.CreatedAt,
		UpdatedAt: invoice.UpdatedAt,
	}
}

func toListInvoiceDetailDomain(invoice []InvoiceDetail) []invoices.InvoiceDetailDomain {
	result := []invoices.InvoiceDetailDomain{}

	for _, e := range invoice {
		result = append(result, toInvoiceDetailDomain(InvoiceDetail{ID: e.ID, Name: e.Name, Email: e.Email, Amount: e.Amount, EventID: e.EventID, Link: e.Link, Status: e.Status, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt}))
	}
	return result
}

func fromInvoiceDetailDomain(domain []*invoices.InvoiceDetailDomain) []InvoiceDetail {
	result := []InvoiceDetail{}

	for _, e := range domain {
		result = append(result, InvoiceDetail{ID: e.ID, Name: e.Name, Email: e.Email, Amount: e.Amount, EventID: e.EventID, Link: e.Link, Status: e.Status, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt})
	}

	return result
}
