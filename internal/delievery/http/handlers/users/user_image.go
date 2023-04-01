package users

import (
	"github.com/gin-gonic/gin"
	resources "legocy-go/internal/delievery/http/resources/users"
	service "legocy-go/internal/delievery/http/service/users"
	"legocy-go/pkg/storage"
	"legocy-go/pkg/storage/client"
	"legocy-go/pkg/storage/models"
	"net/http"
	"strconv"
)

type UserImageHandler struct {
	service service.UserImageUseCase
	storage client.ImageStorage
}

func NewUserImageHandler(service service.UserImageUseCase, storage client.ImageStorage) UserImageHandler {
	return UserImageHandler{
		service: service,
		storage: storage,
	}
}

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

	response := resources.GetUserImageUploadResponse(imgUrl)
	c.JSON(http.StatusOK, response)
}
