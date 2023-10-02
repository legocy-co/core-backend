package marketItem

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeleteMarketItemById
//
//	@Summary	Delete Market Item (Admin)
//	@Tags		market_items_admin
//	@ID			delete_market_item
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

	err = h.service.DeleteMarketItemById(c, itemID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": itemID, "status": "OK"})

}
