package lego

import (
	"legocy-go/internal/api/v1/resources"
	"legocy-go/internal/api/v1/resources/lego"
	s "legocy-go/internal/api/v1/usecase/lego"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LegoSeriesHandler struct {
	service s.LegoSeriesService
}

func NewLegoSeriesHandler(service s.LegoSeriesService) LegoSeriesHandler {
	return LegoSeriesHandler{service: service}
}

func (lsh *LegoSeriesHandler) ListSeries(c *gin.Context) {
	seriesList, err := lsh.service.ListSeries(c.Request.Context())
	if err != nil {
		v1.ErrorRespond(c.Writer, err.Error())
		return
	}

	var seriesResponses []lego.LegoSeriesResponse

	for _, series := range seriesList {
		seriesResponses = append(seriesResponses, lego.GetLegoSeriesResponse(series))
	}

	seriesResponse := v1.DataMetaResponse{
		Data: seriesResponses,
		Meta: v1.SuccessMetaResponse,
	}

	v1.Respond(c.Writer, seriesResponse)
}

func (lsh *LegoSeriesHandler) DetailSeries(c *gin.Context) {
	seriesID, err := strconv.Atoi(c.Param("seriesID"))
	if err != nil {
		v1.ErrorRespond(c.Writer, "Error extracting ID from url path")
		return
	}

	seriesObj, err := lsh.service.DetailSeries(c.Request.Context(), seriesID)
	if err != nil || seriesObj.ID == 0 {
		v1.ErrorRespond(c.Writer, "Error extracting LegoSeries object with given ID")
		return
	}

	seriesResponse := lego.GetLegoSeriesResponse(seriesObj)
	v1.Respond(c.Writer, v1.DataMetaResponse{
		Data: seriesResponse,
		Meta: v1.SuccessMetaResponse,
	})
}

func (lsh *LegoSeriesHandler) SeriesCreate(c *gin.Context) {
	var seriesRequest lego.LegoSeriesRequest

	if err := c.ShouldBindJSON(&seriesRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	seriesObj := seriesRequest.ToLegoSeriesBasic()
	err := lsh.service.CreateLegoSeries(c.Request.Context(), seriesObj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	v1.Respond(c.Writer, v1.DataMetaResponse{
		Data: true,
		Meta: v1.SuccessMetaResponse,
	})
}

func (lsh *LegoSeriesHandler) DeleteSeries(c *gin.Context) {
	seriesID, err := strconv.Atoi(c.Param("seriesID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	err = lsh.service.DeleteSeries(c.Request.Context(), seriesID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	v1.Respond(c.Writer, v1.DataMetaResponse{
		Data: seriesID,
		Meta: v1.SuccessMetaResponse,
	})
}
