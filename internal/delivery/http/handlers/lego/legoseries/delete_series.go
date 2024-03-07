package legoseries

import (
	"github.com/gin-gonic/gin"
	_ "github.com/legocy-co/legocy/docs"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"net/http"
	"strconv"
)

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

	e := lsh.service.DeleteSeries(c.Request.Context(), seriesID)
	if e != nil {
		httpErr := errors.FromAppError(*e)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": seriesID})
}
