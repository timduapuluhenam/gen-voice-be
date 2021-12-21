package routes

import (
	"errors"
	middlewareApp "genVoice/app/middlewares"
	controller "genVoice/controllers"
	"genVoice/controllers/invoices"
	"genVoice/controllers/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig     middleware.JWTConfig
	JWTMiddleware middleware.JWTConfig

	UserController    users.UserController
	InvoiceController invoices.InvoiceController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Register)
	users.POST("/login", cl.UserController.Login)

	invoices := e.Group("invoices")
	invoices.POST("/add", cl.InvoiceController.CreateInvoice)
	invoices.POST("/details", cl.InvoiceController.CreateInvoiceDetail)
}

func RoleValidationUser() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "user" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("unathorized"))
			}
		}
	}
}
