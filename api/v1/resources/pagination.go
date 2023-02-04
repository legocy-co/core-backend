package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/pkg/helpers"
	u "net/url"
	"strconv"
)

type paginationUrls struct {
	Prev string `json:"previous""`
	Curr string `json:"current"`
	Next string `json:"next"`
}

type PaginatedMetaResponse struct {
	Status   string
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	Links    paginationUrls `json:"links"`
}

func getPrevPageUrl(url *u.URL, page int) string {

	var urlCopy *u.URL
	err := helpers.DeepCopy(url, urlCopy)
	if err != nil {
		return ""
	}

	prevPage := page - 1
	switch prevPage {
	case 0:
		return ""
	default:
		urlCopy.Query().Set("page", strconv.Itoa(prevPage))
		return urlCopy.Path
	}
}

func getNextPageUrl(url *u.URL, page int) string {
	var urlCopy *u.URL
	err := helpers.DeepCopy(url, urlCopy)
	if err != nil {
		return ""
	}

	nextPage := page + 1
	urlCopy.Query().Set("page", strconv.Itoa(nextPage))
	return urlCopy.Path
}

func generatePrevNextUrls(url *u.URL, page int) paginationUrls {
	return paginationUrls{
		Prev: getPrevPageUrl(url, page),
		Curr: url.Path,
		Next: getNextPageUrl(url, page),
	}
}

func GetPaginatedMetaResponse(ctx *gin.Context, message string) *PaginatedMetaResponse {
	// TODO:
	return nil
}
