package marketItem

import (
	"github.com/gin-gonic/gin"
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

	marketItemDomain, err := h.service.GetMarketItemByID(c, itemID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	marketItemResponse := admin.GetMarketItemAdminResponse(marketItemDomain)
	c.JSON(http.StatusOK, marketItemResponse)
	return
}
