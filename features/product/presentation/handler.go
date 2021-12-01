package presentation

import (
	"minpro_arya/features/product"
	"minpro_arya/features/product/presentation/request"
	"minpro_arya/features/product/presentation/response"
	"minpro_arya/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.Service
}

func NewHandlerProduct(serv product.Service) *ProductHandler {
	return &ProductHandler{
		productService: serv,
	}
}

func (ctrl *ProductHandler) Create(c echo.Context) error {

	createReq := request.Product{}

	if err := c.Bind(&createReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.productService.Create(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainCreate(result))

}

func (ctrl *ProductHandler) AllProduct(c echo.Context) error {

	result, err := ctrl.productService.AllProduct()

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromProductListDomain(result))

}

func (ctrl *ProductHandler) Update(c echo.Context) error {

	updateReq := request.Product{}

	if err := c.Bind(&updateReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.productService.Update(jwtGetID.ID, id, updateReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainUpdateProduct(result))

}

func (ctrl *ProductHandler) Delete(c echo.Context) error {

	orgzID := middleware.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.productService.Delete(orgzID.ID, deletedId)

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, result)

}

func (ctrl *ProductHandler) MyProductByCompany(c echo.Context) error {
	orgzID := middleware.GetUser(c)

	result, err := ctrl.productService.MyProductByCompany(orgzID.ID)

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromProductListDomain(result))
}

func (ctrl *ProductHandler) ProductByID(c echo.Context) error {

	productID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.productService.ProductByID(productID)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainAllProduct(result))
}

func (ctrl *ProductHandler) ProductByIdCompany(c echo.Context) error {

	orgzID, _ := strconv.Atoi(c.Param("organizerID"))

	result, err := ctrl.productService.ProductByIdCompany(orgzID)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromProductListDomain(result))
}
