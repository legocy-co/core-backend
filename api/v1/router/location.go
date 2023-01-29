package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/api/v1/handlers/marketplace"
	m "legocy-go/api/v1/middleware"
	s "legocy-go/api/v1/usecase/marketplace"
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
