package pagination

import (
	"github.com/legocy-co/legocy/pkg/pagination"
)

type PageResponse[T any] struct {
	Data []T              `json:"data"`
	Meta PageMetaResponse `json:"meta"`
}

type PageMetaResponse struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// GetPageResponse
// TODO: domain Page[T] -> PageResponse[T] without making a new Page[T] object for response structs inside handler
func GetPageResponse[T any](page pagination.Page[T]) PageResponse[T] {

	return PageResponse[T]{
		Data: page.GetObjects(),
		Meta: GetPageMetaResponse(page),
	}
}

func GetPageMetaResponse[T any](page pagination.Page[T]) PageMetaResponse {
	return PageMetaResponse{
		Total:  page.GetTotal(),
		Limit:  page.GetLimit(),
		Offset: page.GetOffset(),
	}
}
