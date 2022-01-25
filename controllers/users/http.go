package users

import (
	middlewareApp "genVoice/app/middlewares"
	"genVoice/business/users"
	controller "genVoice/controllers"
	"genVoice/controllers/users/request"
	"genVoice/controllers/users/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService users.Service
}

func NewUserController(service users.Service) *UserController {
	return &UserController{
		UserService: service,
	}
}

func (ctrl *UserController) Register(c echo.Context) error {
	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.UserService.Register(req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainRegister(data))
}

func (ctrl *UserController) Login(c echo.Context) error {
	req := request.UsersLogin{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.UserService.Login(req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controller.NewSuccessResponse(c, response.FromDomainLogin(data))
}

func (ctrl *UserController) Update(c echo.Context) error {
	req := request.UsersUpdate{}
	req.ID = middlewareApp.GetIdUser(c)

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.UserService.Update(req.ToUpdateDomain())

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	response.FromDomainUpdate(data)
	return controller.NewSuccessResponse(c, response.FromDomainUpdate(data))
}
