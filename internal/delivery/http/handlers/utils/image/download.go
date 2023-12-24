package image

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/utils/image"
	"net/http"
)

func (h Handler) DownloadImage(ctx *gin.Context) {
	imagePath := ctx.Query("fp")
	if imagePath == "" {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "No fp query argument given"},
		)
		return
	}

	downloadRequest, err := image.NewDownloadImageRequest(imagePath)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity, gin.H{"error": err.Error()},
		)
		return
	}

	bucketName, imageName, err := downloadRequest.ToBucketNameImageName()
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()},
		)
		return
	}

	imageData, err := h.storage.DownloadImage(bucketName, imageName)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusFailedDependency, gin.H{"error": err.Error()},
		)
		return
	}

	ctx.DataFromReader(
		http.StatusOK,
		imageData.PayloadSize,
		imageData.PayloadType,
		imageData.Payload,
		map[string]string{},
	)
}
