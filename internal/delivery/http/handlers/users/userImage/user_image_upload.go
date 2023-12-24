package userImage

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/utils/image"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	userModels "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/pkg/storage"
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
//	@Success	200		{object}	schemas.UserImageUploadResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/images/{userID} [post]
func (h UserImageHandler) UploadUserImage(c *gin.Context) {

	uploadHandler := image.NewUploadHandler(
		h.storage,
		storage.UserBucketName,
		"userID",
	)

	imgUrl, err := uploadHandler(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	userID, _ := strconv.Atoi(c.Param("userID"))

	err = h.service.StoreUserImage(context.Background(),
		&userModels.UserImage{UserID: userID, FilepathURL: imgUrl})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
	}

	response := schemas.GetUserImageUploadResponse(imgUrl)
	c.JSON(http.StatusOK, response)
}
