package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/delievery/http/handlers/lego/legoseries"
	m "legocy-go/internal/delievery/http/middleware"
	s "legocy-go/internal/delievery/http/service/lego"
)

func (r V1router) addLegoSeries(rg *gin.RouterGroup, service s.LegoSeriesService) {

	handler := h.NewLegoSeriesHandler(service)
	series := rg.Group("/series").Use(m.Auth())
	{
		series.GET("/", handler.ListSeries)
		series.GET("/:seriesID", handler.DetailSeries)
	}
	seriesPrivate := rg.Group("/admin/series").Use(m.AdminUserOnly())
	{
		seriesPrivate.POST("/", handler.SeriesCreate)
		seriesPrivate.DELETE("/:seriesID", handler.DeleteSeries)
	}
}
