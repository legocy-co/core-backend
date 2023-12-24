package market_item_image

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/pkg/storage"
	image "github.com/legocy-co/legocy/pkg/storage/models"
	"net/http"
	"strconv"
)

func (h Handler) UploadImage(ctx *gin.Context) {
	// Get market item id
	marketItemId, err := strconv.Atoi(ctx.Param("marketItemID"))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity, gin.H{"error": err.Error()},
		)
		return
	}

	// Get file
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()},
		)
		return
	}

	// Open file
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity, gin.H{"error": err.Error()},
		)
		return
	}

	// Domain Image
	img := image.ImageUnitFromFile(src, marketItemId, file.Filename, file.Size)

	// Save image to storage
	imgUrl, err := h.storage.UploadImage(img, storage.MarketItemsBucketName)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError, gin.H{"error": err.Error()},
		)
		return
	}

	// Save img reference to database

	isMain := ctx.Param("isMain") == "true"

	vo, e := models.NewMarketItemImageValueObject(
		marketItemId, imgUrl, isMain,
	)
	if e != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError, gin.H{"error": e.Error()},
		)
		return
	}

	_, e = h.service.StoreMarketItemImage(*vo)
	if e != nil {
		httpErr := errors.FromAppError(*e)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	// Return response
	response := marketplace.NewImageUploadResponse(imgUrl)
	ctx.JSON(http.StatusOK, response)
}
