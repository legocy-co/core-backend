package marketItem

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/resources/marketplace/admin"
	"net/http"
)

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

	err = h.service.CreateMarketItem(c, vo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, itemRequest)

}
