package presentation

import (
	"minpro_arya/features/customer"
	"minpro_arya/features/customer/presentation/request"
	"minpro_arya/features/customer/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerService customer.Service
}

func NewHandlerCustomer(serv customer.Service) *CustomerHandler {
	return &CustomerHandler{
		customerService: serv,
	}
}

func (ctrl *CustomerHandler) Register(c echo.Context) error {

	registerReq := request.Customer{}

	if err := c.Bind(&registerReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.customerService.Register(registerReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *CustomerHandler) Login(c echo.Context) error {

	loginReq := request.CustomerLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.customerService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainLogin(result))
}
