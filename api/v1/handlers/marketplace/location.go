package marketplace

import (
	"github.com/gin-gonic/gin"
	r "legocy-go/api/v1/resources"
	res "legocy-go/api/v1/resources/marketplace"
	s "legocy-go/api/v1/usecase/marketplace"
	"net/http"
)

type LocationHandler struct {
	service s.LocationUseCase
}

func NewLocationHandler(service s.LocationUseCase) LocationHandler {
	return LocationHandler{service: service}
}

func (h *LocationHandler) ListLocations(c *gin.Context) {
	locationsList, err := h.service.ListLocations(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request body")
		c.Abort()
		return
	}

	var locationsResponse []res.LocationResponse
	for _, location := range locationsList {
		locationsResponse = append(locationsResponse, res.GetLocationResponse(location))
	}

	response := r.DataMetaResponse{
		Data: locationsResponse,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
}
