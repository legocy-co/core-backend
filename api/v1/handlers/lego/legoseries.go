package lego

import (
	r "legocy-go/api/v1/resources"
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
		r.ErrorRespond(c.Writer, err.Error())
		return
	}

	seriesResponses := make([]lego.LegoSeriesResponse, 0, len(seriesList))
	for _, series := range seriesList {
		seriesResponses = append(seriesResponses, lego.GetLegoSeriesResponse(series))
	}

	seriesResponse := r.DataMetaResponse{
		Data: seriesResponses,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, seriesResponse)
}

func (lsh *LegoSeriesHandler) DetailSeries(c *gin.Context) {
	seriesID, err := strconv.Atoi(c.Param("seriesID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	seriesObj, err := lsh.service.DetailSeries(c.Request.Context(), seriesID)
	if err != nil || seriesObj.ID == 0 {
		c.JSON(http.StatusBadRequest, "Error extracting LegoSeries object with given ID")
		c.Abort()
		return
	}

	seriesResponse := lego.GetLegoSeriesResponse(seriesObj)
	r.Respond(c.Writer, r.DataMetaResponse{
		Data: seriesResponse,
		Meta: r.SuccessMetaResponse,
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

	r.Respond(c.Writer, r.DataMetaResponse{
		Data: true,
		Meta: r.SuccessMetaResponse,
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

	r.Respond(c.Writer, r.DataMetaResponse{
		Data: seriesID,
		Meta: r.SuccessMetaResponse,
	})
}
