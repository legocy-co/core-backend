package market_item

import (
	"github.com/legocy-co/legocy/internal/delivery/http/middleware/auth"
	"github.com/legocy-co/legocy/lib/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace/filters"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/utils/pagination"
	s "github.com/legocy-co/legocy/internal/domain/marketplace/service"
	"github.com/legocy-co/legocy/internal/pkg/config"
)

type MarketItemHandler struct {
	service        s.MarketItemService
	reviewsService s.UserReviewService
}

func NewMarketItemHandler(service s.MarketItemService, reviewsService s.UserReviewService) MarketItemHandler {
	return MarketItemHandler{
		service:        service,
		reviewsService: reviewsService,
	}
}

// ListMarketItems
//
//	@Summary	Get Market Items
//	@Tags		market_items
//	@ID			list_market_items
//	@Produce	json
//	@Param		limit	query	int	false	"limit" 10
//	@Param		offset	query	int	false	"offset" 0
//
// @Param       filter  query  	filters.MarketItemFilterDTO false "filter"
//
//	@Success	200	{object}	pagination.PageResponse[marketplace.MarketItemResponse]
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/ [get]
func (h *MarketItemHandler) ListMarketItems(c *gin.Context) {

	ctx := pagination.GetPaginationContext(c)

	filter, e := filters.GetMarketItemFilterCritera(c)
	if e != nil {
		httpErr := errors.FromAppError(*e)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	marketItemsPage, err := h.service.ListMarketItems(ctx, filter)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	marketItems := marketItemsPage.GetObjects()

	marketItemResponse := make([]marketplace.MarketItemResponse, 0, len(marketItems))
	for _, m := range marketItems {
		marketItemResponse = append(marketItemResponse, marketplace.GetMarketItemResponse(m))
	}

	responsePage := pagination.GetPageResponse[marketplace.MarketItemResponse](
		marketItemResponse,
		marketItemsPage.GetTotal(),
		marketItemsPage.GetLimit(),
		marketItemsPage.GetOffset(),
	)
	c.JSON(http.StatusOK, responsePage)
}

// ListMarketItemsAuthorized
//
//	@Summary	Get Market Items Authorized
//	@Tags		market_items
//	@ID			list_market_items_authorized
//	@Produce	json
//	@Param		limit	query	int	false	"limit" 10
//	@Param		offset	query	int	false	"offset" 0
//
// @Param       filter  query  	filters.MarketItemFilterDTO false "filter"
//
//	@Success	200	{object}	pagination.PageResponse[marketplace.MarketItemResponse]
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/authorized/ [get]
//
//	@Security	JWT
func (h *MarketItemHandler) ListMarketItemsAuthorized(c *gin.Context) {

	ctx := pagination.GetPaginationContext(c)

	filter, e := filters.GetMarketItemFilterCritera(c)
	if e != nil {
		httpErr := errors.FromAppError(*e)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	tokenPayload, err := auth.GetUserPayload(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	marketItemsPage, appErr := h.service.ListMarketItemsAuthorized(ctx, filter, tokenPayload.ID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	marketItems := marketItemsPage.GetObjects()

	marketItemResponse := make([]marketplace.MarketItemResponse, 0, len(marketItems))
	for _, m := range marketItems {
		marketItemResponse = append(marketItemResponse, marketplace.GetMarketItemResponse(m))
	}

	responsePage := pagination.GetPageResponse[marketplace.MarketItemResponse](
		marketItemResponse,
		marketItemsPage.GetTotal(),
		marketItemsPage.GetLimit(),
		marketItemsPage.GetOffset(),
	)
	c.JSON(http.StatusOK, responsePage)
}

// GetFavorites
//
//	@Summary	Get Favorite Market Items
//	@Tags		market_items
//	@ID			list_market_items_favorites
//	@Produce	json
//	@Param		limit	query	int	false	"limit" 10
//	@Param		offset	query	int	false	"offset" 0
//
//	@Success	200	{object}	pagination.PageResponse[marketplace.MarketItemResponse]
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/favorites/ [get]
//
//	@Security	JWT
func (h *MarketItemHandler) GetFavorites(ctx *gin.Context) {
	tokenPayload, err := auth.GetUserPayload(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paginationCtx := pagination.GetPaginationContext(ctx)

	marketItems, appErr := h.service.GetLikedItems(paginationCtx, tokenPayload.ID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	marketItemResponse := make([]marketplace.MarketItemResponse, 0, len(marketItems.GetObjects()))
	for _, m := range marketItems.GetObjects() {
		marketItemResponse = append(marketItemResponse, marketplace.GetMarketItemResponse(m))
	}

	responsePage := pagination.GetPageResponse[marketplace.MarketItemResponse](
		marketItemResponse,
		marketItems.GetTotal(),
		marketItems.GetLimit(),
		marketItems.GetOffset(),
	)

	ctx.JSON(http.StatusOK, responsePage)
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	marketItem, appErr := h.service.GetMarketItemByID(c, itemID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	userReviewsTotals, _ := h.reviewsService.GetUserReviewsTotals(c, marketItem.Seller.ID)

	marketItemResponse := marketplace.GetMarketItemResponseWithReviewTotals(
		marketItem, userReviewsTotals,
	)

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
	tokenString := auth.GetAuthTokenHeader(c)
	userPayload, ok := jwt.ParseTokenClaims(tokenString, config.GetAppConfig().JwtConf.SecretKey)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized, gin.H{"error": "invalid token credentials"},
		)
		return
	}

	var itemRequest *marketplace.MarketItemRequest
	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": "Error binding MarketItemRequest " + err.Error()},
		)
		return
	}

	// Payload ID as SellerID
	vo, err := itemRequest.ToMarketItemValueObject(userPayload.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	marketItem, appErr := h.service.CreateMarketItem(c, vo)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	marketItemResponse := marketplace.GetMarketItemResponse(marketItem)
	c.JSON(http.StatusOK, marketItemResponse)
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

	e := h.service.DeleteMarketItem(c, itemID)
	if e != nil {
		httpErr := errors.FromAppError(*e)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, map[string]bool{"status": true})
}

// UpdateMarketItemByID
//
//	@Summary	Update Market Item
//	@Tags		market_items
//	@ID			update_market_item
//	@Param		itemId	path	int	true	"item ID"
//	@Param		data	body	marketplace.MarketItemRequest	true	"data"
//	@Produce	json
//	@Success	200	{object}	marketplace.MarketItemResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/{itemId} [put]
//
//	@Security	JWT
func (h *MarketItemHandler) UpdateMarketItemByID(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	tokenString := auth.GetAuthTokenHeader(c)
	userPayload, ok := jwt.ParseTokenClaims(tokenString, config.GetAppConfig().JwtConf.SecretKey)
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

	vo, err := itemRequest.ToMarketItemValueObject(userPayload.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	marketItem, appErr := h.service.UpdateMarketItemByID(c, userPayload.ID, itemID, vo)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	marketItemResponse := marketplace.GetMarketItemResponse(marketItem)
	c.JSON(http.StatusOK, marketItemResponse)
}
