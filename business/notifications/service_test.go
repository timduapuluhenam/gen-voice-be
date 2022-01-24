package notifications_test

import (
	"genVoice/app/middlewares"
	"genVoice/business/activities"
	activitiesMock "genVoice/business/activities/mocks"
	"genVoice/business/invoices"
	"genVoice/business/notifications"
	notificationsMock "genVoice/business/notifications/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	notificationsMockRepo notificationsMock.Repository
	jwtAuth       *middlewares.ConfigJWT
	notificationsService notifications.Service

	mockStatus string
	mockSignatureKey string
	userID int

	activitiesMockRepo activitiesMock.Repository
	activitiesData activities.Domain

	invoicesDetailDomainData invoices.InvoiceDetailDomain
)

func TestMain(m *testing.M){
	jwtAuth = &middlewares.ConfigJWT{
		SecretJWT:       "rahasiamock123",
		ExpiresDuration: 1,
	}
	mockStatus = "settlement"
	mockSignatureKey = "mocksignaturekey"

	invoicesDetailDomainData = invoices.InvoiceDetailDomain{
		
		InvoiceName : "Mockname",

		ID           :"mockidstring",
		Name         :"Mockname",
		Email        :"mockemail@mock.com",
		Amount       :10000,
		EventID      :1,
		SignatureKey :"signatureMock",
		Link         :"mocklink.com",
		Status       :"MockStatus",
		CreatedAt   :time.Now(),
		UpdatedAt   :time.Now(),
	}

	notificationsService = notifications.NewNotifService(&notificationsMockRepo,&activitiesMockRepo,2,jwtAuth)
	m.Run()
}

func TestGetNotif(t *testing.T){
	t.Run("Test Case 1 | Valid Get Notification", func(t *testing.T){
		notificationsMockRepo.On("GetNotif",mockStatus,mock.AnythingOfType("string")).Return(nil).Once()
		notificationsMockRepo.On("GetUserBySignature",mock.AnythingOfType("string")).Return(invoicesDetailDomainData,userID,nil).Once()
		activitiesMockRepo.On("CreateActivity",mock.Anything).Return(activitiesData, nil).Once()
		notificationsService.GetNotif(mockStatus,mockSignatureKey)
		assert.NotNil(t, activitiesData)
	})
	t.Run("Test Case 2 | Invalid Get Notification", func(t *testing.T){
		notificationsMockRepo.On("GetNotif",mockStatus,mock.AnythingOfType("string")).Return(nil).Once()
		notificationsMockRepo.On("GetUserBySignature",mock.AnythingOfType("string")).Return(invoicesDetailDomainData,userID,nil).Once()
		err:=activitiesMockRepo.On("CreateActivity",mock.Anything).Return(activitiesData, assert.AnError).Once()
		notificationsService.GetNotif(mockStatus,mockSignatureKey)
		assert.NotNil(t, err)
	})
}

func TestGetUserBySignature(t *testing.T){
	t.Run("Test Case 1 | Valid Get User By Signature", func(t *testing.T){
		notificationsMockRepo.On("GetUserBySignature",mock.AnythingOfType("string")).Return(invoicesDetailDomainData,userID,nil).Once()
		notificationsService.GetUserBySignature(mockSignatureKey)
		assert.NotNil(t, invoicesDetailDomainData)
	})
	t.Run("Test Case 2 | Invalid Get User By Signature", func(t *testing.T){
		notificationsMockRepo.On("GetUserBySignature",mock.AnythingOfType("string")).Return(invoicesDetailDomainData,userID,assert.AnError).Once()
		_,_,err := notificationsService.GetUserBySignature(mockSignatureKey)
		assert.NotNil(t, err)
	})
}