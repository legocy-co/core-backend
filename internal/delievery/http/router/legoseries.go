package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/handlers/lego"
	m "legocy-go/internal/delievery/http/middleware"
	s "legocy-go/internal/delievery/http/usecase/lego"
)

func (r V1router) addLegoSeries(rg *gin.RouterGroup, service s.LegoSeriesService) {

	handler := lego.NewLegoSeriesHandler(service)
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
