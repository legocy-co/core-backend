package marketplace

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/middleware"
	resources "legocy-go/internal/delievery/http/resources"
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/delievery/http/resources/pagination"
	"legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
	s "legocy-go/internal/domain/marketplace/service"
	"legocy-go/internal/domain/users/middleware"
	"net/http"
	"strconv"
)

type MarketItemHandler struct {
	service s.MarketItemService
}

func NewMarketItemHandler(
	service s.MarketItemService) MarketItemHandler {

	return MarketItemHandler{
		service: service,
	}
}

// ListMarketItems
//
//	@Summary	Get Market Items
//	@Tags		market_items
//	@ID			list_market_items
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/ [get]
//
//	@Security	JWT
func (h *MarketItemHandler) ListMarketItems(c *gin.Context) {

	ctx := pagination.GetPaginationContext(c)

	var marketItems []*models.MarketItem
	marketItems, err := h.service.ListMarketItems(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
	}

	if len(marketItems) == 0 {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"error": errors.ErrMarketItemsNotFound.Error()})
		return
	}

	marketItemResponse := make([]marketplace.MarketItemResponse, 0, len(marketItems))
	for _, m := range marketItems {
		marketItemResponse = append(marketItemResponse, marketplace.GetMarketItemResponse(m))
	}

	response := resources.DataMetaResponse{
		Data: marketItemResponse,
		Meta: pagination.GetPaginatedMetaResponse(
			c.Request.URL.Path, resources.MsgSuccess, ctx),
	}
	resources.Respond(c.Writer, response)
}

// ListMarketItemsAuthorized
//
//	@Summary	Get Market Items Authorized
//	@Tags		market_items
//	@ID			list_market_items_authorized
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/authorized/ [get]
//
//	@Security	JWT
func (h *MarketItemHandler) ListMarketItemsAuthorized(c *gin.Context) {

	ctx := pagination.GetPaginationContext(c)

	var marketItems []*models.MarketItem

	tokenPayload, err := v1.GetUserPayload(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := tokenPayload.ID

	marketItems, err = h.service.ListMarketItemsAuthorized(ctx, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
	}

	if len(marketItems) == 0 {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"error": errors.ErrMarketItemsNotFound.Error()})
		return
	}

	marketItemResponse := make([]marketplace.MarketItemResponse, 0, len(marketItems))
	for _, m := range marketItems {
		marketItemResponse = append(marketItemResponse, marketplace.GetMarketItemResponse(m))
	}

	response := resources.DataMetaResponse{
		Data: marketItemResponse,
		Meta: pagination.GetPaginatedMetaResponse(
			c.Request.URL.Path, resources.MsgSuccess, ctx),
	}
	resources.Respond(c.Writer, response)
}

// MarketItemDetail
//
//	@Summary	Get Market Item
//	@Tags		market_items
//	@ID			detail_market_item
//	@Param		itemID	path	int	true	"item ID"
//	@Produce	json
//	@Success	200	{object}	marketplace.MarketItemResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/{itemID} [get]
//
//	@Security	JWT
func (h *MarketItemHandler) MarketItemDetail(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("itemID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	marketItem, err := h.service.MarketItemDetail(c, itemID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	marketItemResponse := marketplace.GetMarketItemResponse(marketItem)
	c.JSON(http.StatusOK, marketItemResponse)
}

// CreateMarketItem
//
//	@Summary	Create Market Item
//	@Tags		market_items
//	@ID			create_market_item
//	@Param		data	body	marketplace.MarketItemRequest	true	"data"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/ [post]
//
//	@Security	JWT
func (h *MarketItemHandler) CreateMarketItem(c *gin.Context) {
	// If we get here, then token payload is valid
	tokenString := v1.GetAuthTokenHeader(c)
	userPayload, ok := auth.ParseTokenClaims(tokenString)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "invalid token credentials"})
		return
	}

	var itemRequest *marketplace.MarketItemRequest
	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Payload ID as SellerID
	err := h.service.CreateMarketItem(c, itemRequest.ToMarketItemValueObject(userPayload.ID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := resources.DataMetaResponse{
		Data: itemRequest,
		Meta: resources.SuccessMetaResponse,
	}
	resources.Respond(c.Writer, response)
}

// DeleteMarketItem
//
//	@Summary	Delete Market Item
//	@Tags		market_items
//	@ID			delete_market_item
//	@Param		itemId	path	int	true	"item ID"
//	@Produce	json
//	@Success	200	{object}	map[string]bool
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/{itemId} [delete]
//
//	@Security	JWT
func (h *MarketItemHandler) DeleteMarketItem(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	err = h.service.DeleteMarketItem(c, itemID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			"Error deleting MarketItem object")
	}

	c.JSON(http.StatusOK, map[string]bool{"status": true})
}

// UpdateMarketItemByID
//
//	@Summary	Update Market Item
//	@Tags		market_items
//	@ID			update_market_item
//	@Param		itemID	path	int	true	"item ID"
//	@Param		data	body	marketplace.MarketItemRequest	true	"data"
//	@Produce	json
//	@Success	200	{object}	marketplace.MarketItemResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/{itemID} [put]
//
//	@Security	JWT
func (h *MarketItemHandler) UpdateMarketItemByID(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("itemID"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	tokenString := v1.GetAuthTokenHeader(c)
	userPayload, ok := auth.ParseTokenClaims(tokenString)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": "invalid token credentials"})
		return
	}

	var itemRequest *marketplace.MarketItemRequest
	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	marketItem, err := h.service.UpdateMarketItemByID(
		c, userPayload.ID, itemID, itemRequest.ToMarketItemValueObject(userPayload.ID))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	marketItemResponse := marketplace.GetMarketItemResponse(marketItem)
	c.JSON(http.StatusOK, marketItemResponse)
}
