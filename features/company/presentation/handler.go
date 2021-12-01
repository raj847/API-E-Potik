package presentation

import (
	"minpro_arya/features/company"
	"minpro_arya/features/company/presentation/request"
	"minpro_arya/features/company/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CompanyHandler struct {
	companyHand company.Service
}

func NewHandlerCompany(handler company.Service) *CompanyHandler {
	return &CompanyHandler{
		companyHand: handler,
	}
}

func (ctrl *CompanyHandler) Register(c echo.Context) error {

	registerReq := request.Company{}

	if err := c.Bind(&registerReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.companyHand.Register(registerReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *CompanyHandler) Login(c echo.Context) error {

	loginReq := request.CompanyLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.companyHand.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainLogin(result))
}
