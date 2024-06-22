package like

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware/auth"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
)

// GetLikedItems godoc
// @Summary Get Liked Items
// @Tags market_items
// @ID get_liked_items
// @Produce json
// @Success 200 {array} schemas.LikeResponse
// @Failure 401 {object} map[string]interface{}
// @Router /market-items/likes/ [get]
// @Security JWT
func (h *Handler) GetLikedItems(c *gin.Context) {

	tokenPayload, err := auth.GetUserPayload(c)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	likedItems, appErr := h.r.GetLikesByUserID(tokenPayload.ID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	var likedItemsResponse = make([]schemas.LikeResponse, 0, len(likedItems))
	for _, likedItem := range likedItems {
		likedItemsResponse = append(likedItemsResponse, schemas.FromLikeDomain(likedItem))
	}

	c.JSON(200, likedItemsResponse)
}
