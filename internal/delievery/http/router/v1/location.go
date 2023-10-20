package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/delievery/http/handlers/marketplace"
	m "legocy-go/internal/delievery/http/middleware"
	s "legocy-go/internal/domain/marketplace/service"
)

func (r V1router) addLocations(rg *gin.RouterGroup, ser s.LocationUseCase) {
	handler := h.NewLocationHandler(ser)

	locations := rg.Group("/locations").Use(m.Auth())
	{
		locations.GET("/", handler.ListLocations)
		locations.GET("/:country", handler.CountryLocations)
	}
	locationsAdmin := rg.Group("/admin/locations").Use(m.AdminUserOnly())
	{
		locationsAdmin.POST("/", handler.CreateLocation)
		locationsAdmin.DELETE("/:locID", handler.LocationDelete)
	}
}
