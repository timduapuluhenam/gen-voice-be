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
	ID        int       `json:"id:"`
	Name      string    `json:"name"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InvoiceDetailResponse struct {
	// Message   string    `json:"message"`
	ID    int    `json:"id:"`
	Name  string `json:"name"`
	Email string `json:"email"`

	Amount    int       `json:"amount"`
	EventID   int       `json:"event_id"`
	Link      string    `json:"link"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainInvoiceDetail(domain invoices.DatasDomain) DatasResponse {

	result := DatasResponse{}

	result.DataInvoice.ID = domain.DataInvoice.ID
	result.DataInvoice.Name = domain.DataInvoice.Name
	result.DataInvoice.UserID = domain.DataInvoice.UserID
	result.DataInvoice.CreatedAt = domain.DataInvoice.CreatedAt
	result.DataInvoice.UpdatedAt = domain.DataInvoice.UpdatedAt

	fmt.Print("domain invoice detail adasdas  ", domain.InvoiceDetail)
	for _, e := range domain.InvoiceDetail {
		result.InvoiceDetail = append(result.InvoiceDetail, InvoiceDetailResponse{ID: e.ID, Name: e.Name, Email: e.Email, Amount: e.Amount, EventID: e.EventID, Link: e.Link, Status: "Belum Dibayar", CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt})
	}
	return result
}

func GenerateReportFromListDomain(domlist []invoices.InvoiceDetailDomain) []InvoiceDetailResponse {
	result := []InvoiceDetailResponse{}
	for _, e := range domlist {
		result = append(result, InvoiceDetailResponse{ID: e.ID, Name: e.Name, Email: e.Email, Amount: e.Amount, EventID: e.EventID, Link: e.Link, Status: "Belum Dibayar", CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt})
	}
	return result
}
