package lego

import (
	"legocy-go/internal/delievery/http/resources"
	"legocy-go/internal/delievery/http/resources/lego"
	s "legocy-go/internal/delievery/http/usecase/lego"
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

// ListSeries
//
//	@Summary	List of LEGO Series objects
//	@Tags		lego_series
//	@ID			list_lego_series
//	@Produce	json
//	@Success	200	{object}	[]lego.LegoSeriesResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/series/ [get]
//
//	@Security	JWT
func (lsh *LegoSeriesHandler) ListSeries(c *gin.Context) {
	seriesList, err := lsh.service.ListSeries(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seriesResponses := make([]lego.LegoSeriesResponse, 0, len(seriesList))
	for _, series := range seriesList {
		seriesResponses = append(seriesResponses, lego.GetLegoSeriesResponse(series))
	}

	c.JSON(http.StatusOK, seriesResponses)
}

// DetailSeries
//
//	@Summary	LEGO Series by id
//	@Tags		lego_series
//	@ID			detail_lego_series
//	@Param		seriesID	path	int	true	"series ID"
//	@Produce	json
//	@Success	200	{object}	lego.LegoSeriesResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/series/{seriesID} [get]
//
//	@Security	JWT
func (lsh *LegoSeriesHandler) DetailSeries(c *gin.Context) {
	seriesID, err := strconv.Atoi(c.Param("seriesID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seriesObj, err := lsh.service.DetailSeries(c.Request.Context(), seriesID)
	if err != nil || seriesObj.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error extracting LegoSeries object with given ID"})
		c.Abort()
		return
	}

	seriesResponse := lego.GetLegoSeriesResponse(seriesObj)
	c.JSON(http.StatusOK, seriesResponse)
}

// SeriesCreate
//
//	@Summary	Create LEGO Series object
//	@Tags		lego_series_admin
//	@ID			create_series
//	@Param		data	body	lego.LegoSeriesRequest	true	"create data"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/series/ [post]
//
//	@Security	JWT
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

// DeleteSeries
//
//	@Summary	Delete LegoSeries object
//	@Tags		lego_series_admin
//	@ID			delete_series
//	@Param		seriesID	path	int	true	"series ID"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/series/{seriesID} [delete]
//
//	@Security	JWT
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
