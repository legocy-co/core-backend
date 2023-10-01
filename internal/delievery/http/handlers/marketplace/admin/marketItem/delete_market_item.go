package marketItem

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
