package response

import (
	"fmt"
	"genVoice/business/invoices"
	"time"
)

type DatasResponse struct {
	DataInvoice   InvoiceResponse
	InvoiceDetail []InvoiceDetailResponse
}

type InvoiceResponse struct {
	// Message   string    `json:"message"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	UserID      int       `json:"user_id"`
	TimeExpired int       `json:"TimeExpired"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InvoiceDetailResponse struct {
	// Message   string    `json:"message"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	Amount      int       `json:"amount"`
	EventID     int       `json:"event_id"`
	InvoiceName string    `json:"invoice_name"`
	Link        string    `json:"link"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomainInvoiceDetail(domain invoices.DatasDomain) DatasResponse {

	result := DatasResponse{}

	result.DataInvoice.ID = domain.DataInvoice.ID
	result.DataInvoice.Name = domain.DataInvoice.Name
	result.DataInvoice.UserID = domain.DataInvoice.UserID
	result.DataInvoice.TimeExpired = domain.DataInvoice.TimeExpired
	result.DataInvoice.CreatedAt = domain.DataInvoice.CreatedAt
	result.DataInvoice.UpdatedAt = domain.DataInvoice.UpdatedAt

	fmt.Print("domain invoice detail adasdas  ", domain.InvoiceDetail)
	for _, e := range domain.InvoiceDetail {
		result.InvoiceDetail = append(result.InvoiceDetail, InvoiceDetailResponse{ID: e.ID, Name: e.Name, Email: e.Email, Amount: e.Amount, EventID: e.EventID, Link: e.Link, Status: "Not Paid", CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt, InvoiceName: result.DataInvoice.Name})
	}
	return result
}

func GenerateReportFromListDomain(domlist []invoices.InvoiceDetailDomain) []InvoiceDetailResponse {
	result := []InvoiceDetailResponse{}
	for _, e := range domlist {
		fmt.Print("cek status hereeee", e.Status)
		result = append(result, InvoiceDetailResponse{ID: e.ID, Name: e.Name, Email: e.Email, Amount: e.Amount, EventID: e.EventID, InvoiceName: e.InvoiceName, Link: e.Link, Status: e.Status, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt})
	}
	return result
}

func FromDomainDeleteInvoice(deleteInvoice invoices.InvoiceDomain) InvoiceResponse {
	result := InvoiceResponse{}
	result.ID = deleteInvoice.ID
	result.Name = deleteInvoice.Name
	result.UserID = deleteInvoice.ID
	result.TimeExpired = deleteInvoice.TimeExpired
	result.UpdatedAt = deleteInvoice.UpdatedAt
	result.CreatedAt = deleteInvoice.CreatedAt
	return result
}

func FromDomainGetInvoiceDetails(domain invoices.InvoiceDetailDomain) InvoiceDetailResponse {
	result := InvoiceDetailResponse{}

	result.ID = domain.ID
	result.Name = domain.Name
	result.Email = domain.Email
	result.Amount = domain.Amount
	result.EventID = domain.EventID
	result.InvoiceName = domain.InvoiceName
	result.Link = domain.Link
	result.Status = domain.Status
	result.CreatedAt = domain.CreatedAt
	result.UpdatedAt = domain.UpdatedAt

	return result
}

func FromDomainInvoices(domlist []invoices.InvoiceDomain) []InvoiceResponse {
	result := []InvoiceResponse{}
	for _, e := range domlist {
		result = append(result, InvoiceResponse{ID: e.ID, UserID: e.ID, Name: e.Name, TimeExpired: e.TimeExpired, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt})
	}
	return result
}
