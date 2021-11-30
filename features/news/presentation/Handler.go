package news

import (
	"fmt"
	"minpro_arya/features/news"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NewsHandler struct {
	NewsRepo news.Repository
}

func NewNewsHandler(newsRepo news.Repository) *NewsHandler {
	return &NewsHandler{
		NewsRepo: newsRepo,
	}
}

func (newsHandler NewsHandler) GetNewsByCategory(c echo.Context) error {
	category := "health"
	fmt.Println(category)
	data, error := newsHandler.NewsRepo.GetNewsByCategory(category)

	if error != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return NewSuccessResponse(c, FromDomain(data))
}
