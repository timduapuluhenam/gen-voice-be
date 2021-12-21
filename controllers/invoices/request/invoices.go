package request

import (
	"genVoice/business/invoices"
)

type Invoice struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

func (req *Invoice) ToInvoiceDomain() *invoices.InvoiceDomain {
	return &invoices.InvoiceDomain{
		Name:   req.Name,
		UserID: req.UserID,
	}
}

type InvoiceDetail []struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Amount  int    `json:"amount"`
	EventID string `json:"event_id"`
	Link    string `json:"link"`
	Status  string `json:"status"`
}

func (req *InvoiceDetail) ToInvoiceDetailDomain() []*invoices.InvoiceDetailDomain {

	result := []*invoices.InvoiceDetailDomain{}

	for _, e := range *req {
		result = append(result, &invoices.InvoiceDetailDomain{Name: e.Name, Email: e.Email, Amount: e.Amount, EventID: e.EventID, Link: e.Link, Status: e.Status})
	}
	return result
}
