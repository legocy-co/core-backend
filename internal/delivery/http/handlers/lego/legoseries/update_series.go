package legoseries

import (
	"github.com/gin-gonic/gin"
	_ "github.com/legocy-co/legocy/docs"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/lego"
	"net/http"
	"strconv"
)

// UpdateSeries
//
//	@Summary	Update LegoSeries object
//	@Tags		lego_series_admin
//	@ID			put_series
//	@Param		data	body	lego.LegoSeriesRequest	true	"create data"
//	@Param		seriesID	path	int	true	"Lego Series ID"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/series/{seriesID} [put]
//
//	@Security	JWT
func (lsh *LegoSeriesHandler) UpdateSeries(c *gin.Context) {
	var seriesRequest lego.LegoSeriesRequest
	if _err := c.ShouldBindJSON(&seriesRequest); _err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": _err.Error()})
		return
	}

	seriesID, _err := strconv.Atoi(c.Param("seriesID"))
	if _err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	legoSeriesValueObject := seriesRequest.ToLegoSeriesValueObject()
	err := lsh.service.UpdateSeries(seriesID, legoSeriesValueObject)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
