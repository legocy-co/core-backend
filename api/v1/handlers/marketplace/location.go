package marketplace

import (
	"github.com/gin-gonic/gin"
	r "legocy-go/api/v1/resources"
	res "legocy-go/api/v1/resources/marketplace"
	s "legocy-go/api/v1/usecase/marketplace"
	"net/http"
	"strconv"
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
		c.JSON(http.StatusBadRequest, err.Error())
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

func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var locationRequest res.LocationRequest
	if err := c.ShouldBindJSON(&locationRequest); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request body")
		c.Abort()
		return
	}

	locationBasic := locationRequest.ToLocationBasic()
	err := h.service.CreateLocation(c, locationBasic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	response := r.DataMetaResponse{Data: true, Meta: r.SuccessMetaResponse}
	r.Respond(c.Writer, response)
}

func (h *LocationHandler) CountryLocations(c *gin.Context) {
	country := c.Param("country")
	if country == "" {
		c.JSON(http.StatusBadRequest, "Invalid request")
		c.Abort()
		return
	}

	locations, err := h.service.CountryLocations(c.Request.Context(), country)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	var locationsResponse []res.LocationResponse
	for _, location := range locations {
		locationsResponse = append(locationsResponse, res.GetLocationResponse(location))
	}

	response := r.DataMetaResponse{
		Data: locationsResponse,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
}

func (h *LocationHandler) LocationDelete(c *gin.Context) {
	locID, err := strconv.Atoi(c.Param("locID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	err = h.service.DeleteLocation(c.Request.Context(), locID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	response := r.DataMetaResponse{
		Data: true,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
}
