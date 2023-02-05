package pagination

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"legocy-go/pkg/filter"
	"log"
	"strconv"
)

func GetPaginationContext(c *gin.Context) context.Context {
	params, err := filter.GetQueryParams(c)
	if err != nil {
		log.Println(err.Error())
		return c
	}
	return context.WithValue(context.Background(), "pagination", params)
}

type paginationUrls struct {
	Prev string `json:"previous""`
	Curr string `json:"current"`
	Next string `json:"next"`
}

type PaginatedMetaResponse struct {
	Message  string         `json:"message"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	Links    paginationUrls `json:"links"`
}

func GetPaginatedMetaResponse(
	url string, message string, pagination Pagination) PaginatedMetaResponse {

	page, err := strconv.Atoi(pagination.Page)
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(pagination.Limit)
	if err != nil {
		limit = 0
	}

	return PaginatedMetaResponse{
		Message:  message,
		Page:     page,
		PageSize: limit,
		Links:    generateMetaUrls(url, page),
	}

}

func getPrevPageUrl(url string, page int) string {
	prevPage := page - 1
	var pageSymbol string

	if prevPage <= 0 {
		pageSymbol = ""
	} else {
		pageSymbol = strconv.Itoa(prevPage)
	}

	return url + fmt.Sprintf("?page=%v", pageSymbol)

}

func getNextPageUrl(url string, page int) string {
	return url + fmt.Sprintf("?page=%v", strconv.Itoa(page+1))
}

func generateMetaUrls(url string, page int) paginationUrls {
	return paginationUrls{
		Prev: getPrevPageUrl(url, page),
		Curr: url,
		Next: getNextPageUrl(url, page),
	}
}
