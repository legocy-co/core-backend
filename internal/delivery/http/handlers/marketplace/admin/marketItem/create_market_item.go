package marketItem

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace/admin"
	"net/http"
)

// CreateMarketItem
//
//	@Summary	Create Market Item (Admin)
//	@Tags		market_items_admin
//	@ID			create_market_item_admin
//	@Param		data	body	admin.MarketItemAdminCreateRequest	true	"data"
//	@Produce	json
//	@Success	200	{object}	admin.MarketItemAdminResponse
//	@Failure	409	{object}	map[string]interface{}
//	@Failure	422	{object}	map[string]interface{}
//	@Router		/admin/market-items/ [post]
//
//	@Security	JWT
func (h Handler) CreateMarketItem(c *gin.Context) {
	var itemRequest *admin.MarketItemAdminCreateRequest
	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vo, err := itemRequest.ToMarketItemAdminValueObject()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	appErr := h.service.CreateMarketItem(c, vo)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, itemRequest)

}
