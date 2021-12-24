package response

import (
	"genVoice/business/invoices"
	"time"
)

type InvoiceResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Name      string    `json:"name"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainInvoice(domain invoices.InvoiceDomain) InvoiceResponse {
	return InvoiceResponse{
		Message:   "Invoice Success",
		ID:        domain.ID,
		Name:      domain.Name,
		UserID:    domain.UserID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type InvoiceDetailResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Name      string    `json:"name"`
	Email     string    `json:"emial"`
	Amount    int       `json:"amount"`
	EventID   string    `json:"event_id"`
	Link      string    `json:"link"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainInvoiceDetail(domain []invoices.InvoiceDetailDomain) []InvoiceDetailResponse {
	result := []InvoiceDetailResponse{}

	for _, e := range domain {
		result = append(result, InvoiceDetailResponse{Message: "Invoice Success", ID: e.ID, Name: e.Name, Email: e.Email, Amount: e.Amount, EventID: e.EventID, Link: e.Link, Status: e.Status, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt})
	}
	return result
}
