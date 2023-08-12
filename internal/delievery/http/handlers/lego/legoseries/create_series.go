package legoseries

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	v1 "legocy-go/internal/delievery/http/resources"
	"legocy-go/internal/delievery/http/resources/lego"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	seriesObj := seriesRequest.ToLegoSeriesValueObject()
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
