package marketItem

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/errors"
	"legocy-go/internal/delievery/http/resources/marketplace/admin"
	"net/http"
	"strconv"
)

// GetMarketItemByID
//
//	@Summary	Get Market Item by ID (Admin)
//	@Tags		market_items_admin
//	@ID			detail_market_item_admin
//	@Param		itemId	path	int	true	"item ID"
//	@Produce	json
//	@Success	200	{object}	admin.MarketItemAdminResponse
//	@Failure	404	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/market-items/{itemId} [get]
//
//	@Security	JWT
func (h Handler) GetMarketItemByID(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	marketItemDomain, appErr := h.service.GetMarketItemByID(c, itemID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	marketItemResponse := admin.GetMarketItemAdminResponse(marketItemDomain)
	c.JSON(http.StatusOK, marketItemResponse)
	return
}
