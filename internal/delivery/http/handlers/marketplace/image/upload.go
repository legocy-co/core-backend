package image

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/utils/image"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/pkg/storage"
	"net/http"
	"strconv"
)

// UploadImage
//
//		@Summary	Upload Image
//		@Tags		market_item_images
//		@ID			upload_market_item_image
//		@Accept		multipart/form-data
//		@Produce	json
//		@Param		file	formData  file	true	"filepath"
//	 	@Param		marketItemID path int true "market item id"
//		@Success	200		{object}	marketplace.ImageUploadResponse
//		@Failure	400		{object}	map[string]interface{}
//		@Router		/market-items/images/{marketItemID} [post]
//
//		@Security JWT
func (h Handler) UploadImage(ctx *gin.Context) {

	// Call generic upload handler
	uploadHandler := image.NewUploadHandler(
		h.storage,
		storage.MarketItemsBucketName,
		"marketItemID",
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
	marketItemId, err := strconv.Atoi(ctx.Param("marketItemID"))

	isMain := ctx.Param("isMain") == "true"

	vo, err := models.NewMarketItemImageValueObject(
		marketItemId, imgUrl, isMain,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError, gin.H{"error": err.Error()},
		)
		return
	}

	createdImage, e := h.service.StoreMarketItemImage(*vo)
	if e != nil {
		httpErr := errors.FromAppError(*e)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	// Return response
	response := marketplace.NewImageUploadResponse(imgUrl, createdImage)
	ctx.JSON(http.StatusOK, response)
}
