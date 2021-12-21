package invoices

import (
	"fmt"
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

func (ctrl *InvoiceController) CreateInvoice(c echo.Context) error {
	req := request.Invoice{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	fmt.Print(req)
	data, err := ctrl.InvoiceService.CreateInvoice(req.ToInvoiceDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainInvoice(data))
}

func (ctrl *InvoiceController) CreateInvoiceDetail(c echo.Context) error {
	req := request.InvoiceDetail{}
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
