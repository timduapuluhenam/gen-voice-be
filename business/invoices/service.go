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
	if err != nil {
		return DatasDomain{}, err
	}
	_, errActivity := servUser.activityRepo.CreateActivity(&activities.Domain{
		UserID:    invoiceDetailDomain.DataInvoice.UserID,
		Activity:  fmt.Sprintf("Invoice %s successfully added.", invoice.DataInvoice.Name),
		CreatedAt: invoice.DataInvoice.CreatedAt,
		UpdatedAt: invoice.DataInvoice.UpdatedAt})

	if errActivity != nil {
		return DatasDomain{}, nil
	}
	return invoice, nil
}

func (servUser *InvoiceService) GetAllByUserID(userID int) ([]InvoiceDetailDomain, error) {
	invoice, err := servUser.repository.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

func (servUser *InvoiceService) DeleteInvoice(invoiceID int) (InvoiceDomain, error) {
	invoice, err := servUser.repository.DeleteInvoice(invoiceID)
	if err != nil {
		return InvoiceDomain{}, err
	}
	return invoice, nil
}

func (servUser *InvoiceService) GetInvoiceDetailByID(id string) (InvoiceDetailDomain, error){
	invoice, err := servUser.repository.GetInvoiceDetailByID(id)
	if err != nil {
		return InvoiceDetailDomain{}, err
	}
	return invoice, nil
}

func (servUser *InvoiceService) GetAllEventByUserID(userID int) ([]InvoiceDomain, error){
	invoice, err := servUser.repository.GetAllEventByUserID(userID)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}