package v1

import (
	res "legocy-go/api/v1/resources"
	s "legocy-go/api/v1/usecase"
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
		res.ErrorRespond(c.Writer, "Error extracting LEGO Series List")
		return
	}

	seriesResponse := res.DataMetaResponse{
		Data: seriesList,
		Meta: map[string]interface{}{
			"status": 200,
			"msg":    "OK",
		},
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

	seriesResponse := res.GetLegoSeriesResponse(seriesObj)
	res.Respond(c.Writer, res.DataMetaResponse{
		Data: seriesResponse,
		Meta: map[string]interface{}{
			"status": 200,
			"msg":    res.MSG_SUCCESS,
		},
	})
}

func (lsh *LegoSeriesHandler) SeriesCreate(c *gin.Context) {
	var seriesRequest res.LegoSeriesRequest

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
		Data: res.GetLegoSeriesResponse(seriesObj),
		Meta: map[string]interface{}{
			"status": 200,
			"msg":    res.MSG_SUCCESS,
		},
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
		Meta: map[string]interface{}{
			"status": 200,
			"msg":    res.MSG_SUCCESS,
		},
	})
}
