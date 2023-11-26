package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/delivery/http/handlers/lego/legoset"
	m "legocy-go/internal/delivery/http/middleware"
	s "legocy-go/internal/domain/lego/service"
)

func (r V1router) addLegoSets(rg *gin.RouterGroup, service s.LegoSetService) {
	handler := h.NewLegoSetHandler(service)

	sets := rg.Group("/sets").Use(m.Auth())
	{
		sets.GET("/", handler.ListSets)
		sets.GET("/:setID", handler.SetDetail)
	}
	setsAdmin := rg.Group("/admin/sets").Use(m.AdminUserOnly())
	{
		setsAdmin.POST("/", handler.SetCreate)
		setsAdmin.DELETE("/:setID", handler.SetDelete)
	}
}
