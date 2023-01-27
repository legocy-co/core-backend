package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/api/v1/handlers/lego"
	m "legocy-go/api/v1/middleware"
	s "legocy-go/api/v1/usecase/lego"
)

func (r V1router) addLegoSets(rg *gin.RouterGroup, service s.LegoSetUseCase) {
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
