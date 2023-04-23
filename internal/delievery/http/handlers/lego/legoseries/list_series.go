package legoseries

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	"legocy-go/internal/delievery/http/resources/lego"
	"net/http"
)

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
