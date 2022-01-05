package request

import (
	"genVoice/business/invoices"
)

type Datas struct {
	DataInvoice   Invoice         `json:"invoice"`
	InvoiceDetail []InvoiceDetail `json:"invoiceDetails"`
}

type Invoice struct {
	Name        string `json:"name"`
	UserID      int    `json:"UserID"`
	TimeExpired int    `json:"TimeExpired"`
}

type InvoiceDetail struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Amount  int    `json:"amount"`
	Status  string
	EventID int
}

func (req *Datas) ToInvoiceDetailDomain() *invoices.DatasDomain {

	result := &invoices.DatasDomain{}
	result.DataInvoice.Name = req.DataInvoice.Name
	result.DataInvoice.UserID = req.DataInvoice.UserID
	result.DataInvoice.TimeExpired = req.DataInvoice.TimeExpired

	for _, e := range req.InvoiceDetail {
		result.InvoiceDetail = append(result.InvoiceDetail, invoices.InvoiceDetailDomain{Name: e.Name, Email: e.Email, Amount: e.Amount, EventID: e.EventID})
	}
	return result
}
