package marketplace

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/errors"
	"legocy-go/internal/delievery/http/resources"
	"legocy-go/internal/delievery/http/resources/marketplace"
	s "legocy-go/internal/domain/marketplace/service"
	"net/http"
	"strconv"
)

type LocationHandler struct {
	service s.LocationUseCase
}

func NewLocationHandler(service s.LocationUseCase) LocationHandler {
	return LocationHandler{service: service}
}

// ListLocations
//
//	@Summary	Get locations
//	@Tags		locations
//	@ID			locations_list
//	@Produce	json
//	@Success	200	{object}	[]marketplace.LocationResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/locations/ [get]
//
//	@Security	JWT
func (h *LocationHandler) ListLocations(c *gin.Context) {
	locationsList, err := h.service.ListLocations(c)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	var locationsResponse []marketplace.LocationResponse
	for _, location := range locationsList {
		locationsResponse = append(locationsResponse, marketplace.GetLocationResponse(location))
	}

	c.JSON(http.StatusOK, locationsResponse)
}

// CreateLocation
//
//	@Summary	Create Location
//	@Tags		locations_admin
//	@ID			locations_create
//	@Param		data	body	marketplace.LocationRequest	true	"location data"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/locations/ [post]
//
//	@Security	JWT
func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var locationRequest marketplace.LocationRequest
	if _err := c.ShouldBindJSON(&locationRequest); _err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	locationValueObject := locationRequest.ToLocationValueObject()
	err := h.service.CreateLocation(c, locationValueObject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := v1.DataMetaResponse{Data: true, Meta: v1.SuccessMetaResponse}
	v1.Respond(c.Writer, response)
}

// CountryLocations
//
//	@Summary	Get locations by country
//	@Tags		locations
//	@ID			locations_country
//	@Param		country	path	string	true	"country"
//	@Produce	json
//	@Success	200	{object}	[]marketplace.LocationResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/locations/{country} [get]
//
//	@Security	JWT
func (h *LocationHandler) CountryLocations(c *gin.Context) {
	country := c.Param("country")
	if country == "" {
		c.JSON(http.StatusBadRequest, "Invalid request")
		c.Abort()
		return
	}

	locations, err := h.service.CountryLocations(c.Request.Context(), country)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	var locationsResponse []marketplace.LocationResponse
	for _, location := range locations {
		locationsResponse = append(locationsResponse, marketplace.GetLocationResponse(location))
	}

	response := v1.DataMetaResponse{
		Data: locationsResponse,
		Meta: v1.SuccessMetaResponse,
	}

	v1.Respond(c.Writer, response)
}

// LocationDelete
//
//	@Summary	Get locations by country
//	@Tags		locations_admin
//	@ID			locations_delete
//	@Param		locID	path	int	true	"location ID"
//	@Produce	json
//	@Success	200	{object}	[]marketplace.LocationResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/locations/{locID} [delete]
//
//	@Security	JWT
func (h *LocationHandler) LocationDelete(c *gin.Context) {
	locID, _err := strconv.Atoi(c.Param("locID"))
	if _err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	err := h.service.DeleteLocation(c.Request.Context(), locID)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	response := v1.DataMetaResponse{
		Data: true,
		Meta: v1.SuccessMetaResponse,
	}

	v1.Respond(c.Writer, response)
}
