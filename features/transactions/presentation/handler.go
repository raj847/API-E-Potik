package transactions

import (
	"minpro_arya/features/transactions"
	"minpro_arya/features/transactions/presentation/request"
	"minpro_arya/features/transactions/presentation/response"
	"minpro_arya/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransHandler struct {
	transService transactions.Service
}

func NewHandlerProduct(serv transactions.Service) *TransHandler {
	return &TransHandler{
		transService: serv,
	}
}

func (ctrl *TransHandler) Create(c echo.Context) error {

	createReq := request.Transactions{}

	if err := c.Bind(&createReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.transService.Trans(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainCreate(result))

}

func (ctrl *TransHandler) GetAllTrans(c echo.Context) error {

	result, err := ctrl.transService.GetAllTrans()

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromTransListDomain(result))

}

func (ctrl *TransHandler) GetTransByID(c echo.Context) error {

	transID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.transService.GetTransByID(transID)

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainAllTrans(result))
}

func (ctrl *TransHandler) GetAllUserTrans(c echo.Context) error {

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.transService.GetAllUserTrans(jwtGetID.ID)

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromTransListDomain(result))
}
