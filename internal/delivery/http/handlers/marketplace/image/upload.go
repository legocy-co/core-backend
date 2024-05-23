package image

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/utils/image"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/pkg/s3"
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
//		@Param		sortIndex	formData  int	false	"sort index"
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
		s3.MarketItemsBucketName,
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

	// Get Sort Index from request
	sortIndex, err := strconv.Atoi(ctx.PostForm("sortIndex"))
	if err != nil {
		sortIndex = 0
	}

	vo, e := models.NewMarketItemImageValueObject(
		marketItemId, imgUrl, sortIndex,
	)
	if e != nil {
		httpErr := errors.FromAppError(*e)
		ctx.AbortWithStatusJSON(
			httpErr.Status, httpErr.Message,
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
