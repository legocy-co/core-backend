package marketItem

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delivery/http/errors"
	resources "legocy-go/internal/delivery/http/resources"
	"legocy-go/internal/delivery/http/resources/marketplace/admin"
	"legocy-go/internal/delivery/http/resources/pagination"
	"legocy-go/internal/domain/marketplace/models"
)

// GetMarketItemsAdmin
//
//	@Summary	Get Market Items (Admin)
//	@Tags		market_items_admin
//	@ID			list_market_items_admin
//	@Produce	json
//	@Success	200	{object}	resources.DataMetaResponse
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

	dataMetaResponse := resources.DataMetaResponse{
		Data: response,
		Meta: pagination.GetPaginatedMetaResponse(
			c.Request.URL.Path,
			resources.MsgSuccess,
			c),
	}
	resources.Respond(c.Writer, dataMetaResponse)

}
