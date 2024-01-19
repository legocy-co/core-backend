package image

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"net/http"
	"strconv"
)

// DeleteImageById
//
//	@Summary	Delete Image By Id
//	@Tags		lego_set_images
//	@ID			delete_lego_set_image_by_id
//	@Accept		json
//	@Produce	json
//	@Param		imageId path int true "image id"
//	@Success	200		{object}	map[string]interface{}
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/admin/sets/images/{imageId} [delete]
func (h LegoSetImageHandler) DeleteImageById(ctx *gin.Context) {
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
		httpErr := errors.FromAppError(*e)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
