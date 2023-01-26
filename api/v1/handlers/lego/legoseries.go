package lego

import (
	res "legocy-go/api/v1/resources"
	"legocy-go/api/v1/resources/lego"
	s "legocy-go/api/v1/usecase/lego"
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
		res.ErrorRespond(c.Writer, err.Error())
		return
	}

	var seriesResponses []lego.LegoSeriesResponse

	for _, series := range seriesList {
		seriesResponses = append(seriesResponses, lego.GetLegoSeriesResponse(series))
	}

	seriesResponse := res.DataMetaResponse{
		Data: seriesResponses,
		Meta: res.SuccessMetaResponse,
	}

	res.Respond(c.Writer, seriesResponse)
}

func (lsh *LegoSeriesHandler) DetailSeries(c *gin.Context) {
	seriesID, err := strconv.Atoi(c.Param("seriesID"))
	if err != nil {
		res.ErrorRespond(c.Writer, "Error extracting ID from url path")
		return
	}

	seriesObj, err := lsh.service.DetailSeries(c.Request.Context(), seriesID)
	if err != nil || seriesObj.ID == 0 {
		res.ErrorRespond(c.Writer, "Error extracting LegoSeries object with given ID")
		return
	}

	seriesResponse := lego.GetLegoSeriesResponse(seriesObj)
	res.Respond(c.Writer, res.DataMetaResponse{
		Data: seriesResponse,
		Meta: res.SuccessMetaResponse,
	})
}

func (lsh *LegoSeriesHandler) SeriesCreate(c *gin.Context) {
	var seriesRequest lego.LegoSeriesRequest

	if err := c.ShouldBindJSON(&seriesRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	seriesObj := seriesRequest.ToLegoSeries()
	err := lsh.service.CreateLegoSeries(c.Request.Context(), seriesObj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	res.Respond(c.Writer, res.DataMetaResponse{
		Data: lego.GetLegoSeriesResponse(seriesObj),
		Meta: res.SuccessMetaResponse,
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

	res.Respond(c.Writer, res.DataMetaResponse{
		Data: seriesID,
		Meta: res.SuccessMetaResponse,
	})
}
