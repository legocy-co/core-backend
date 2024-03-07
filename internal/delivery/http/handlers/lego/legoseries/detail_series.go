package legoseries

import (
	"github.com/gin-gonic/gin"
	_ "github.com/legocy-co/legocy/docs"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/lego"
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

	seriesObj, e := lsh.service.DetailSeries(c.Request.Context(), seriesID)
	if e != nil {
		httpErr := errors.FromAppError(*e)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	seriesResponse := lego.GetLegoSeriesResponse(seriesObj)
	c.JSON(http.StatusOK, seriesResponse)
}
