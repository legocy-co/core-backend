package userImage

import (
	"context"
	"github.com/gin-gonic/gin"
	resources "legocy-go/internal/delivery/http/resources/users"
	userModels "legocy-go/internal/domain/users/models"
	"legocy-go/pkg/storage"
	"legocy-go/pkg/storage/models"
	"net/http"
	"strconv"
)

// UploadUserImage
//
//	@Summary	Download User Image
//	@Tags		users_images
//	@ID			upload_user_image
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		file	formData  file	true	"filepath"
//	@Success	200		{object}	resources.UserImageUploadResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/images/:userID [post]
func (h UserImageHandler) UploadUserImage(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	src, err := file.Open()
	defer src.Close()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	img := models.ImageUnitFromFile(src, userID, file.Filename, file.Size)

	imgUrl, err := h.storage.UploadImage(img, storage.UserBucketName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.StoreUserImage(context.Background(),
		&userModels.UserImage{UserID: userID, FilepathURL: imgUrl})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
	}

	response := resources.GetUserImageUploadResponse(imgUrl)
	c.JSON(http.StatusOK, response)
}
