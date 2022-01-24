package invoices_test

import (
	"genVoice/app/middlewares"
	"genVoice/business/activities"
	activitiesMock "genVoice/business/activities/mocks"
	"genVoice/business/invoices"
	invoicesMock "genVoice/business/invoices/mocks"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	invoicesMockRepo invoicesMock.Repository
	jwtAuth       *middlewares.ConfigJWT
	invoicesService invoices.Service
	invoicesData invoices.DatasDomain
	invoicesDomainData invoices.InvoiceDomain
	invoicesDomainDatas []invoices.InvoiceDomain
	invoicesDetailDomainData []invoices.InvoiceDetailDomain

	activitiesMockRepo activitiesMock.Repository
	activitiesData activities.Domain
)

func TestMain(m *testing.M){
	jwtAuth = &middlewares.ConfigJWT{
		SecretJWT:       "rahasiamock123",
		ExpiresDuration: 1,
	}
	invoicesDomainData = invoices.InvoiceDomain{
		ID          :1,
		UserID      :1,
		Name        :"Mockname",
		TimeExpired :1,
		CreatedAt   :time.Now(),
		UpdatedAt   :time.Now(),
	}
	invoicesDomainDatas = []invoices.InvoiceDomain{
		{
			ID          :1,
			UserID      :1,
			Name        :"Mockname",
			TimeExpired :1,
			CreatedAt   :time.Now(),
			UpdatedAt   :time.Now(),
		},
	}
	invoicesDetailDomainData = []invoices.InvoiceDetailDomain{
		{
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
	},

	}
	invoicesData = invoices.DatasDomain{
		DataInvoice: invoicesDomainData,
		InvoiceDetail: invoicesDetailDomainData,
	}
	invoicesService = invoices.NewInvoiceService(&invoicesMockRepo,&activitiesMockRepo,2,jwtAuth)
	m.Run()
}

func TestCreateInvoiceDetail(t *testing.T){
	t.Run("Test Case 1 | Valid Create Invoice Detail", func(t *testing.T){
		invoicesMockRepo.On("CreateInvoiceDetail",mock.Anything).Return(invoicesData, nil).Once()
		activitiesMockRepo.On("CreateActivity",mock.Anything).Return(activitiesData, nil).Once()

		resp, err := invoicesService.CreateInvoiceDetail(&invoicesData)
		assert.Nil(t, err)
		assert.Equal(t, invoicesData, resp)
	})
	t.Run("Test Case 2 | Invalid Create Invoice Detail", func(t *testing.T){
		invoicesMockRepo.On("CreateInvoiceDetail",mock.Anything).Return(invoicesData, assert.AnError).Once()
		// activitiesMockRepo.On("CreateActivity",mock.Anything).Return(activitiesData, assert.AnError).Once()

		resp, err := invoicesService.CreateInvoiceDetail(&invoicesData)
		assert.NotNil(t, err)
		assert.IsType(t, invoicesData, resp)
		// assert.Equal(t, invoicesData, resp)
	})
	t.Run("Test Case 3 | Invalid Create Activity", func(t *testing.T){
		invoicesMockRepo.On("CreateInvoiceDetail",mock.Anything).Return(invoicesData, nil).Once()
		activitiesMockRepo.On("CreateActivity",mock.Anything).Return(activitiesData, assert.AnError).Once()

		resp, _ := invoicesService.CreateInvoiceDetail(&invoicesData)
		assert.IsType(t, invoicesData, resp)
	})
}

func TestGetAllByUserID(t *testing.T){
	t.Run("Test Case 1 | Valid Get All By User ID",func(t *testing.T){
		invoicesMockRepo.On("GetAllByUserID",mock.Anything).Return(invoicesDetailDomainData, nil).Once()
		resp, err := invoicesService.GetAllByUserID(1)
		assert.Nil(t, err)
		assert.IsType(t, invoicesDetailDomainData, resp)
	})
	t.Run("Test Case 2 | Invalid Get All By User ID",func(t *testing.T){
		invoicesMockRepo.On("GetAllByUserID",mock.Anything).Return(invoicesDetailDomainData, assert.AnError).Once()
		resp, err := invoicesService.GetAllByUserID(1)
		assert.NotNil(t, err)
		assert.IsType(t, invoicesDetailDomainData, resp)
	})
}

func TestDeleteInvoice(t *testing.T){
	t.Run("Test Case 1 | Valid Delete Invoice by ID",func(t *testing.T){
		invoicesMockRepo.On("DeleteInvoice",mock.Anything).Return(invoicesDomainData, nil).Once()
		resp, err := invoicesService.DeleteInvoice(1)
		assert.Nil(t, err)
		assert.Equal(t, invoicesDomainData, resp)
	})
	t.Run("Test Case 2 | Invalid Delete Invoice by ID",func(t *testing.T){
		invoicesMockRepo.On("DeleteInvoice",mock.Anything).Return(invoicesDomainData, assert.AnError).Once()
		resp, err := invoicesService.DeleteInvoice(1)
		assert.NotNil(t, err)
		assert.IsType(t, invoicesDomainData, resp)
	})
}

func TestGetInvoiceDetailByID(t *testing.T){
	t.Run("Test Case 1 | Valid Get Invoice Detail by ID",func(t *testing.T){
		invoicesMockRepo.On("GetInvoiceDetailByID",mock.Anything).Return(invoicesDetailDomainData[0], nil).Once()
		resp, err := invoicesService.GetInvoiceDetailByID("1")
		assert.Nil(t, err)
		assert.Equal(t, invoicesDetailDomainData[0], resp)
	})
	t.Run("Test Case 2 | Invalid Get Invoice Detail by ID",func(t *testing.T){
		invoicesMockRepo.On("GetInvoiceDetailByID",mock.Anything).Return(invoicesDetailDomainData[0], assert.AnError).Once()
		resp, err := invoicesService.GetInvoiceDetailByID("1")
		assert.NotNil(t, err)
		assert.IsType(t, invoicesDetailDomainData[0], resp)
	})
}

func TestGetAllEventByUserID(t *testing.T){
	t.Run("Test Case 1 | Valid Get Invoice by ID",func(t *testing.T){
		invoicesMockRepo.On("GetAllEventByUserID",mock.AnythingOfType("int")).Return(invoicesDomainDatas, nil).Once()
		resp, err := invoicesService.GetAllEventByUserID(1)
		assert.Nil(t, err)
		assert.Equal(t, invoicesDomainDatas[0], resp[0])
	})
	t.Run("Test Case 2 | Invalid Get Invoice by ID",func(t *testing.T){
		invoicesMockRepo.On("GetAllEventByUserID",mock.AnythingOfType("int")).Return(invoicesDomainDatas, assert.AnError).Once()
		resp, err := invoicesService.GetAllEventByUserID(1)
		assert.NotNil(t, err)
		assert.IsType(t, invoicesDomainDatas, resp)
	})
}