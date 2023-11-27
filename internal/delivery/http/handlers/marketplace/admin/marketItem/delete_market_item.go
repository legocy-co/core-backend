package marketItem

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delivery/http/errors"
	"net/http"
	"strconv"
)

// DeleteMarketItemById
//
//	@Summary	Delete Market Item (Admin)
//	@Tags		market_items_admin
//	@ID			delete_market_item_admin
//	@Param		itemId	path	int	true	"item ID"
//	@Produce	json
//	@Success	200	{object}	map[string]bool
//	@Failure	404	{object}	map[string]interface{}
//	@Router		/market-items/{itemId} [delete]
//
//	@Security	JWT
func (h Handler) DeleteMarketItemById(c *gin.Context) {

	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	appErr := h.service.DeleteMarketItemById(c, itemID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": itemID, "status": "OK"})

}
