package invoices

import (
	"fmt"
	"genVoice/app/middlewares"
	"genVoice/business/activities"
	"time"
)

type InvoiceService struct {
	repository     Repository
	activityRepo   activities.Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewInvoiceService(repo Repository, activityRepo activities.Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &InvoiceService{
		repository:     repo,
		activityRepo:   activityRepo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (servUser *InvoiceService) CreateInvoiceDetail(invoiceDetailDomain *DatasDomain) (DatasDomain, error) {
	invoice, err := servUser.repository.CreateInvoiceDetail(invoiceDetailDomain)
	_, errActivity := servUser.activityRepo.CreateActivity(&activities.Domain{
		UserID:    invoiceDetailDomain.DataInvoice.UserID,
		Activity:  fmt.Sprintf("Invoice %s berhasil ditambahkan.", invoice.DataInvoice.Name),
		CreatedAt: invoice.DataInvoice.CreatedAt,
		UpdatedAt: invoice.DataInvoice.UpdatedAt})

	if err != nil {
		return DatasDomain{}, err
	}
	if errActivity != nil {
		return DatasDomain{}, nil
	}
	return invoice, nil
}
