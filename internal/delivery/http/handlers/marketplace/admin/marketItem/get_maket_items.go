package marketItem

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace/admin"
	"github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"net/http"
)

// GetMarketItemsAdmin
//
//	@Summary	Get Market Items (Admin)
//	@Tags		market_items_admin
//	@ID			list_market_items_admin
//	@Produce	json
//	@Success	200	{object}	utils.DataMetaResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/market-items/authorized/ [get]
//
//	@Security	JWT
func (h Handler) GetMarketItemsAdmin(c *gin.Context) {
	var marketItems []*marketplace.MarketItemAdmin

	marketItems, err := h.service.GetMarketItems(c)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	var response = make([]admin.MarketItemAdminResponse, 0, len(marketItems))
	for _, marketItem := range marketItems {
		response = append(response, admin.GetMarketItemAdminResponse(marketItem))
	}

	// TODO: add pagination
	c.JSON(http.StatusOK, response)
}
