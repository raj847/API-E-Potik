package news

import (
	"minpro_arya/features/news"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"
)

type NewsResponse struct {
	Article interface{} `json:"news"`
}

func FromDomain(domain news.Domain) NewsResponse {
	return NewsResponse{
		Article: domain.Article,
	}
}

func FromListDomain(domain []news.Domain) []NewsResponse {
	var response []NewsResponse
	for _, value := range domain {
		response = append(response, FromDomain(value))
	}
	return response
}

type CustomerRegisterResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
