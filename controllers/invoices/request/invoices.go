package request

import (
	"genVoice/business/invoices"
	"genVoice/helper/idgenerator"
	"strconv"
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
type DeleteInvoice struct {
	InvoiceID int
}

func (req *Datas) ToInvoiceDetailDomain() *invoices.DatasDomain {

	result := &invoices.DatasDomain{}
	result.DataInvoice.Name = req.DataInvoice.Name
	result.DataInvoice.UserID = req.DataInvoice.UserID
	result.DataInvoice.TimeExpired = req.DataInvoice.TimeExpired

	for _, e := range req.InvoiceDetail {
		id, _ := idgenerator.SF.NextID()
		IDstr := strconv.FormatUint(id, 10)
		result.InvoiceDetail = append(result.InvoiceDetail, invoices.InvoiceDetailDomain{ID: IDstr, Name: e.Name, Email: e.Email, Amount: e.Amount, EventID: e.EventID})
	}
	return result
}

func (req *DeleteInvoice) ToDeleteInvoiceDomain() *invoices.InvoiceDomain {
	// middlewareApp.GetIdUser(echo.Context)
	return &invoices.InvoiceDomain{
		ID: req.InvoiceID,
	}

}
