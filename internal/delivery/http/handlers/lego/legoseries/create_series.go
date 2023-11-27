package legoseries

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	"legocy-go/internal/delivery/http/resources/lego"
	"net/http"
)

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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seriesObj := seriesRequest.ToLegoSeriesValueObject()
	err := lsh.service.CreateLegoSeries(c.Request.Context(), seriesObj)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
