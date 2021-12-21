package invoices

import (
	"genVoice/app/middlewares"
	"time"
)

type InvoiceService struct {
	repository     Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewInvoiceService(repo Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &InvoiceService{
		repository:     repo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (servUser *InvoiceService) CreateInvoice(invoicedomain *InvoiceDomain) (InvoiceDomain, error) {
	invoice, err := servUser.repository.CreateInvoice(invoicedomain)

	if err != nil {
		return InvoiceDomain{}, err
	}
	return invoice, nil
}

func (servUser *InvoiceService) CreateInvoiceDetail(invoiceDetailDomain []*InvoiceDetailDomain) ([]InvoiceDetailDomain, error) {
	invoice, err := servUser.repository.CreateInvoiceDetail(invoiceDetailDomain)

	if err != nil {
		return []InvoiceDetailDomain{}, err
	}
	return invoice, nil
}
