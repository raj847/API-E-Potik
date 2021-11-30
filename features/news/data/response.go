package data

import (
	"minpro_arya/features/news"
	"time"
)

type Response struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    struct {
		Link        string `json:"link"`
		Description string `json:"description"`
		Title       string `json:"title"`
		Image       string `json:"image"`
		Article     []struct {
			Link        string    `json:"link"`
			Title       string    `json:"title"`
			PubDate     time.Time `json:"pubDate"`
			Description string    `json:"description"`
			Thumbnail   string    `json:"thumbnail"`
		} `json:"posts"`
	} `json:"data"`
}

type Articles []struct {
	Link        string    `json:"link"`
	Title       string    `json:"title"`
	PubDate     time.Time `json:"pubDate"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
}

func (resp *Response) ToDomain() news.Domain {
	return news.Domain{
		Article: resp.Data,
	}
}

func ToListDomain(data []Response) []news.Domain {
	result := []news.Domain{}
	for _, domain := range data {
		result = append(result, domain.ToDomain())
	}
	return result
}
