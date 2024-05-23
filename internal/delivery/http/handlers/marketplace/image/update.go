package image

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	"net/http"
	"strconv"
)

// Update
// @Summary	Update Image
// @Tags		market_item_images
// @ID			update_market_item_image
// @Accept		json
// @Produce	json
// @Param		imageID path int true "image id"
// @Param		marketItemID path int true "market item id"
// @Param		data body schemas.ImageUpdateRequest true "image update request"
// @Success	200		{object}	map[string]interface{}
// @Failure	400		{object}	map[string]interface{}
// @Router		/market-items/images/{marketItemID}/{imageID} [patch]
//
// @Security JWT
func (h Handler) Update(ctx *gin.Context) {

	imageId, err := strconv.Atoi(ctx.Param("imageID"))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": err.Error()},
		)
		return
	}

	var req *schemas.ImageUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": err.Error()},
		)
		return
	}

	if _, appErr := h.service.UpdateImageByID(imageId, req.ToVO()); appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Image updated successfully"})
}
