package like

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware/auth"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"strconv"
)

// LikeMarketItem godoc
// @Summary Like Market Item
// @Tags market_items
// @ID like_market_item
// @Produce json
// @Param marketItemID path int true "Market Item ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /market-items/likes/{marketItemID} [post]
// @Security JWT
func (h *Handler) LikeMarketItem(c *gin.Context) {

	marketItemID, err := strconv.Atoi(c.Param("marketItemID"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid market item ID"})
		return
	}

	tokenPayload, err := auth.GetUserPayload(c)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if appErr := h.r.AddLike(
		models.LikeValueObject{
			MarketItemID: marketItemID,
			UserID:       tokenPayload.ID,
		},
	); appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
	}

	c.JSON(200, gin.H{"message": "Market item liked"})
}
