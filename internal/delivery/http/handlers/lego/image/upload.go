package image

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/utils/image"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/lego"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/pkg/storage"
	"net/http"
	"strconv"
)

// Upload
//
//		@Summary	Upload Image
//		@Tags		lego_set_images
//		@ID			upload_lego_set_image
//		@Accept		multipart/form-data
//		@Produce	json
//		@Param		file	formData  file	true	"filepath"
//	 	@Param		legoSetID path int true "lego set id"
//		@Success	200		{object}	schemas.LegoSetImageResponse
//		@Failure	400		{object}	map[string]interface{}
//		@Router		/admin/sets/images/{legoSetID} [post]
//
//		@Security JWT
func (h LegoSetImageHandler) Upload(ctx *gin.Context) {
	// Call generic upload handler
	uploadHandler := image.NewUploadHandler(
		h.storage,
		storage.LegoSetsBucketName,
		"legoSetID",
	)

	imgUrl, err := uploadHandler(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Assuming err == nil always since upload handler didn't return error
	legoSetId, _ := strconv.Atoi(ctx.Param("legoSetID"))

	vo, e := models.NewLegoSetImageValueObject(legoSetId, imgUrl, false)
	if e != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": e.Error()})
		return
	}

	// Call service
	img, appErr := h.service.StoreLegoSetImage(*vo)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	ctx.JSON(http.StatusOK, schemas.GetLegoSetImageResponse(img))
}
