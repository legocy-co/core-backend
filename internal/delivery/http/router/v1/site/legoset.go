package site

import (
	"github.com/gin-gonic/gin"
	h "github.com/legocy-co/legocy/internal/delivery/http/handlers/lego/legoset"
	s "github.com/legocy-co/legocy/internal/domain/lego/service"
	m "github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
)

func AddLegoSets(rg *gin.RouterGroup, service s.LegoSetService) {
	handler := h.NewLegoSetHandler(service)

	sets := rg.Group("/sets").Use(m.IsAuthenticated())
	{
		sets.GET("/", handler.ListSets)
		sets.GET("/:setID", handler.SetDetail)
	}
	setsAdmin := rg.Group("/admin/sets").Use(m.IsAdmin())
	{
		setsAdmin.POST("/", handler.SetCreate)
		setsAdmin.DELETE("/:setID", handler.SetDelete)
	}
}
