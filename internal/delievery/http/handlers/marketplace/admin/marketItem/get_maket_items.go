package marketItem

import (
	"github.com/gin-gonic/gin"
	resources "legocy-go/internal/delievery/http/resources"
	"legocy-go/internal/delievery/http/resources/marketplace/admin"
	"legocy-go/internal/delievery/http/resources/pagination"
	"legocy-go/internal/domain/marketplace/models"
	"net/http"
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var response = make([]admin.MarketItemAdminResponse, 0, len(marketItems))
	for _, marketItem := range marketItems {
		response = append(response, admin.GetMarketItemAdminResponse(marketItem))
	}

	dataMetaResponse := resources.DataMetaResponse{
		Data: response,
		Meta: pagination.GetPaginatedMetaResponse(
			c.Request.URL.Path, resources.MsgSuccess, c),
	}
	resources.Respond(c.Writer, dataMetaResponse)

}
