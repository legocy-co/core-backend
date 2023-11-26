package userImage

import (
	"github.com/gin-gonic/gin"
	resources "legocy-go/internal/delivery/http/resources/users"
	"log"
	"net/http"
)

// DownloadImage
//
//	@Summary	Download User Image
//	@Tags		users_images
//	@ID			download_user_image
//	@Param		fp	query string	true	"filepath"
//	@Success	200		{file}	file
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/images/download [get]
func (h UserImageHandler) DownloadImage(c *gin.Context) {

	imagePath := c.Query("fp")
	if imagePath == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No fp query argument given"})
		return
	}

	downloadRequest, err := resources.NewUserDownloadImageRequest(imagePath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

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
