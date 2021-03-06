package notifications

import (
	"fmt"
	"time"

	"genVoice/app/middlewares"
	"genVoice/business/activities"
	"genVoice/business/invoices"
)

type NotifService struct {
	repository     Repository
	activityRepo   activities.Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewNotifService(repo Repository, activityRepo activities.Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &NotifService{
		repository:     repo,
		activityRepo:   activityRepo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (servUser *NotifService) GetNotif(status, signature_key string) error {
	servUser.repository.GetNotif(status, signature_key)
	if status == "settlement" {
		customer, userID, _ := servUser.repository.GetUserBySignature(signature_key)

		_, errActivity := servUser.activityRepo.CreateActivity(&activities.Domain{
			UserID:   userID,
			Activity: fmt.Sprintf("Customer %s has made payment of Rp %d on invoice %d", customer.Name, customer.Amount, customer.EventID)})

		if errActivity != nil {
			return nil
		}
	}
	return nil
}

func (servUser *NotifService) GetUserBySignature(signature_key string) (invoices.InvoiceDetailDomain, int, error) {
	invoiceDetail, userID, err := servUser.repository.GetUserBySignature(signature_key)
	if err != nil {
		return invoices.InvoiceDetailDomain{}, userID, err
	}
	return invoiceDetail, userID, nil
}
