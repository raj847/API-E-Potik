package response

import (
	"minpro_arya/features/transactions"

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

type CreateTransResponse struct {
	Message          string `json:"message"`
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	ProductID        int    `json:"product_id"`
	Transaction_code string `json:"transaction_code"`
}

func FromDomainCreate(domain transactions.Domain) CreateTransResponse {
	return CreateTransResponse{
		Message:          "Transactions Success, Silahkan ambil pesanan anda pada apotik terdekat, Salam Sehat",
		ID:               domain.ID,
		UserID:           domain.UserID,
		ProductID:        domain.ProductID,
		Transaction_code: domain.Transaction_code,
	}
}

type TransResponse struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	ProductID        int    `json:"product_id"`
	Transaction_code string `json:"transaction_code"`
}

func FromDomainAllTrans(domain transactions.Domain) TransResponse {
	return TransResponse{
		ID:               domain.ID,
		UserID:           domain.UserID,
		ProductID:        domain.ProductID,
		Transaction_code: domain.Transaction_code,
	}
}

func FromTransListDomain(domain []transactions.Domain) []TransResponse {
	var response []TransResponse
	for _, value := range domain {
		response = append(response, FromDomainAllTrans(value))
	}
	return response
}
