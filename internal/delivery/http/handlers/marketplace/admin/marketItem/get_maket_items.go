package marketItem

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace/admin"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/utils/pagination"
	"net/http"
)

// GetMarketItemsAdmin
//
//	@Summary	Get Market Items (Admin)
//	@Tags		market_items_admin
//	@ID			list_market_items_admin
//	@Produce	json
//	@Param		limit	query	int	false	"limit" 10
//	@Param		offset	query	int	false	"offset" 0
//	@Success	200	{object}	pagination.PageResponse[admin.MarketItemAdminResponse]
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/market-items/ [get]
//
//	@Security	JWT
func (h Handler) GetMarketItemsAdmin(c *gin.Context) {

	ctx := pagination.GetPaginationContext(c)

	marketItemsPage, err := h.service.GetMarketItems(ctx)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	marketItems := marketItemsPage.GetObjects()

	var response = make([]admin.MarketItemAdminResponse, 0, len(marketItems))
	for _, marketItem := range marketItems {
		response = append(response, admin.GetMarketItemAdminResponse(marketItem))
	}

	responsePage := pagination.GetPageResponse[admin.MarketItemAdminResponse](
		response,
		marketItemsPage.GetTotal(),
		marketItemsPage.GetLimit(),
		marketItemsPage.GetOffset(),
	)
	c.JSON(http.StatusOK, responsePage)
}
