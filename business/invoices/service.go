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

func (servUser *InvoiceService) CreateInvoiceDetail(invoiceDetailDomain *DatasDomain) (DatasDomain, error) {
	invoice, err := servUser.repository.CreateInvoiceDetail(invoiceDetailDomain)

	if err != nil {
		return DatasDomain{}, err
	}
	return invoice, nil
}
