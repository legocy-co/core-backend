package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/api/v1/handlers/marketplace"
	m "legocy-go/internal/api/v1/middleware"
	s "legocy-go/internal/api/v1/usecase/marketplace"
)

func (r V1router) addLocations(rg *gin.RouterGroup, ser s.LocationUseCase) {
	handler := h.NewLocationHandler(ser)

	locations := rg.Group("/locations").Use(m.Auth())
	{
		locations.GET("/", handler.ListLocations)
	}
	locationsAdmin := rg.Group("/admin/locations").Use(m.AdminUserOnly())
	{
		locationsAdmin.POST("/", handler.CreateLocation)
	}
}
