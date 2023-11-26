package legoseries

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	"legocy-go/internal/delivery/http/resources/lego"
	"net/http"
	"strconv"
)

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
