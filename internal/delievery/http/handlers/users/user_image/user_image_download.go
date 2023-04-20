package user_image

import (
	"github.com/gin-gonic/gin"
	resources "legocy-go/internal/delievery/http/resources/users"
	"log"
	"net/http"
)

func (h user_image.UserImageHandler) DownloadImage(c *gin.Context) {

	imagePath := c.Query("fp")
	if imagePath == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No fp query argument given"})
		return
	}
	downloadRequest := resources.UserDownloadImageRequest{ImagePath: imagePath}

	bucketName, imageName, err := downloadRequest.ToBucketNameImageName()
	log.Printf("%v %v %v", bucketName, imageName, err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageData, err := h.storage.DownloadImage(bucketName, imageName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return
	}

	c.DataFromReader(
		http.StatusOK,
		imageData.PayloadSize,
		"image/png",
		imageData.Payload,
		map[string]string{},
	)
}
