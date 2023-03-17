package pagination

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	r "legocy-go/internal/delievery/http/resources"
	"legocy-go/pkg/filter"
	"log"
	"strconv"
)

type PaginatedMetaResponse struct {
	Message  string         `json:"message"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	Links    paginationUrls `json:"links"`
}

func GetPaginationContext(c *gin.Context) context.Context {
	params, err := filter.GetQueryParams(c)
	if err != nil {
		log.Println(err.Error())
		return c
	}
	return context.WithValue(context.Background(), "pagination", params)
}

func GetPaginatedMetaResponse(
	url string, message string, ctx context.Context) PaginatedMetaResponse {

	params := ctx.Value("pagination").(*filter.QueryParams)
	if params == nil {
		return PaginatedMetaResponse{
			Message:  r.MsgError,
			Page:     1,
			PageSize: 10,
			Links:    generateMetaUrls(url, 1),
		}
	}

	return PaginatedMetaResponse{
		Message:  message,
		Page:     params.Page,
		PageSize: params.PageSize,
		Links:    generateMetaUrls(url, params.Page),
	}

}

type paginationUrls struct {
	Prev string `json:"previous""`
	Curr string `json:"current"`
	Next string `json:"next"`
}

func generateMetaUrls(url string, page int) paginationUrls {
	return paginationUrls{
		Prev: getPrevPageUrl(url, page),
		Curr: url,
		Next: getNextPageUrl(url, page),
	}
}

func getPrevPageUrl(url string, page int) string {
	prevPage := page - 1

	if prevPage <= 0 {
		return url
	}

	pageSymbol := strconv.Itoa(prevPage)
	return url + fmt.Sprintf("?page=%v", pageSymbol)

}

func getNextPageUrl(url string, page int) string {
	return url + fmt.Sprintf("?page=%v", strconv.Itoa(page+1))
}
