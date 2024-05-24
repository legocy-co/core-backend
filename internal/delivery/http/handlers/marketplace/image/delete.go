package image

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Delete
// @Summary	Delete Image
// @Tags		market_item_images
// @ID			delete_market_item_image
// @Accept		json
// @Produce	json
// @Param		imageId path int true "image id"
// @Param marketItemID path int true "market item id"
// @Success	200		{object}	map[string]interface{}
// @Failure	400		{object}	map[string]interface{}
// @Router		/market-items/images/{marketItemID}/{imageId} [delete]
//
// @Security JWT
func (h Handler) Delete(ctx *gin.Context) {

	imageId, err := strconv.Atoi(ctx.Param("imageId"))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": err.Error()},
		)
		return
	}

	e := h.service.DeleteImageByID(imageId)
	if e != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": e.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
