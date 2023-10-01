package marketItem

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/resources/marketplace/admin"
	"net/http"
	"strconv"
)

func (h Handler) UpdateMarketItemByID(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("itemID"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	var itemRequest *admin.MarketItemAdminUpdateRequest
	if err := c.ShouldBindJSON(itemRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	vo, err := itemRequest.ToMarketItemAdminValueObject()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	marketItemDomain, err := h.service.UpdateMarketItem(c, itemID, vo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	marketItemResponse := admin.GetMarketItemAdminResponse(marketItemDomain)
	c.JSON(http.StatusOK, marketItemResponse)
}
