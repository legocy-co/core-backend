package marketItem

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/errors"
	"legocy-go/internal/delievery/http/resources/marketplace/admin"
	"net/http"
	"strconv"
)

// UpdateMarketItemByID
//
//	@Summary	Update Market Item (Admin)
//	@Tags		market_items_admin
//	@ID			update_market_item_admin
//	@Param		itemId	path	int	true  "item ID"
//	@Param		data	body	admin.MarketItemAdminUpdateRequest	true	"data"
//	@Produce	json
//	@Success	200	{object}	admin.MarketItemAdminResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/market-items/{itemId} [put]
//
//	@Security	JWT
func (h Handler) UpdateMarketItemByID(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	var itemRequest *admin.MarketItemAdminUpdateRequest
	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	vo, err := itemRequest.ToMarketItemAdminValueObject()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	marketItemDomain, appErr := h.service.UpdateMarketItem(c, itemID, vo)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	marketItemResponse := admin.GetMarketItemAdminResponse(marketItemDomain)
	c.JSON(http.StatusOK, marketItemResponse)
}
