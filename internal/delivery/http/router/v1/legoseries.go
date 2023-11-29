package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/delivery/http/handlers/lego/legoseries"
	s "legocy-go/internal/domain/lego/service"
	m "legocy-go/pkg/auth/jwt/middleware"
)

func (r V1router) addLegoSeries(rg *gin.RouterGroup, service s.LegoSeriesService) {

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
	}
}
