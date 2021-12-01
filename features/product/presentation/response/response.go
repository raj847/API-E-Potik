package response

import (
	"minpro_arya/features/product"
	"time"

	"net/http"

	echo "github.com/labstack/echo/v4"
)

type CreateProductResponse struct {
	Message     string    `json:"message"`
	ID          int       `json:"id:"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	NetWeight   int       `json:"net-weight"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something not right"
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(status, response)
}

func FromDomainCreate(domain product.Domain) CreateProductResponse {
	return CreateProductResponse{
		Message:     "Create Product Success",
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		NetWeight:   domain.NetWeight,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

type ProductResponse struct {
	ID          int       `json:"id:"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	NetWeight   int       `json:"net-weight"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomainAllProduct(domain product.Domain) ProductResponse {
	return ProductResponse{
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		NetWeight:   domain.NetWeight,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromDomainUpdateProduct(domain product.Domain) CreateProductResponse {
	return CreateProductResponse{
		Message:     "Update Product Success",
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		NetWeight:   domain.NetWeight,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromProductListDomain(domain []product.Domain) []ProductResponse {
	var response []ProductResponse
	for _, value := range domain {
		response = append(response, FromDomainAllProduct(value))
	}
	return response
}
