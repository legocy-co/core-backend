package site

import (
	"github.com/gin-gonic/gin"
	h "github.com/legocy-co/legocy/internal/delivery/http/handlers/lego/legoseries"
	m "github.com/legocy-co/legocy/internal/delivery/http/middleware/auth"
	s "github.com/legocy-co/legocy/internal/domain/lego/service"
)

func AddLegoSeries(rg *gin.RouterGroup, service s.LegoSeriesService) {

	handler := h.NewLegoSeriesHandler(service)
	series := rg.Group("/series").Use(m.IsAuthenticated())
	{
		series.GET("/", handler.ListSeries)
		series.GET("/:seriesID", handler.DetailSeries)
	}
	seriesPrivate := rg.Group("/admin/series").Use(m.IsAdmin())
	{
		seriesPrivate.POST("/", handler.SeriesCreate)
		seriesPrivate.DELETE("/:seriesID", handler.DeleteSeries)
		seriesPrivate.PUT("/:seriesID", handler.UpdateSeries)
	}
}
