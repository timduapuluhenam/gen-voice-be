package invoices

import (
	"fmt"
	middlewareApp "genVoice/app/middlewares"
	"genVoice/business/invoices"
	controller "genVoice/controllers"
	"genVoice/controllers/invoices/request"
	"genVoice/controllers/invoices/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InvoiceController struct {
	InvoiceService invoices.Service
}

func NewInvoiceController(service invoices.Service) *InvoiceController {
	return &InvoiceController{
		InvoiceService: service,
	}
}

func (ctrl *InvoiceController) CreateInvoiceDetail(c echo.Context) error {
	req := request.Datas{}

	req.DataInvoice.UserID = middlewareApp.GetIdUser(c)
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result := req.ToInvoiceDetailDomain()

	data, err := ctrl.InvoiceService.CreateInvoiceDetail(result)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainInvoiceDetail(data))
}

func (ctrl *InvoiceController) GetAllByUserID(c echo.Context) error {
	userID := middlewareApp.GetIdUser(c)
	data, err := ctrl.InvoiceService.GetAllByUserID(userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	fmt.Print(data)
	return controller.NewSuccessResponse(c, response.GenerateReportFromListDomain(data))
}

func (ctrl *InvoiceController) DeleteInvoice(c echo.Context) error {
	req := request.DeleteInvoice{}

	// req.DataInvoice.UserID = middlewareApp.GetIdUser(c)
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.InvoiceService.DeleteInvoice(req.InvoiceID)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	fmt.Print(" DATA : ", data)
	return controller.NewSuccessResponse(c, response.FromDomainDeleteInvoice(data))
}

func (ctrl *InvoiceController) GetInvoiceDetailByID(c echo.Context) error {
	id := c.Param("id")
	data, err := ctrl.InvoiceService.GetInvoiceDetailByID(id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, response.FromDomainGetInvoiceDetails(data))
}

func (ctrl *InvoiceController) GetAllEventByUserID(c echo.Context) error {
	userID := middlewareApp.GetIdUser(c)
	data, err := ctrl.InvoiceService.GetAllEventByUserID(userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	fmt.Print(data)
	return controller.NewSuccessResponse(c, response.FromDomainInvoices(data))
}