package users

import (
	"context"
	"github.com/gin-gonic/gin"
	resources "legocy-go/internal/delievery/http/resources/users"
	service "legocy-go/internal/delievery/http/service/users"
	m "legocy-go/internal/domain/users/models"
	"legocy-go/pkg/storage"
	"legocy-go/pkg/storage/client"
	"legocy-go/pkg/storage/models"
	"log"
	"net/http"
	"strconv"
)

type UserImageHandler struct {
	service service.UserImageUseCase
	storage client.ImageStorage
}

func NewUserImageHandler(
	service service.UserImageUseCase, storage client.ImageStorage) UserImageHandler {
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

	err = h.service.StoreUserImage(context.Background(),
		&m.UserImage{UserID: userID, FilepathURL: imgUrl})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
	}

	response := resources.GetUserImageUploadResponse(imgUrl)
	c.JSON(http.StatusOK, response)
}

func (h UserImageHandler) DownloadImage(c *gin.Context) {

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

func (h UserImageHandler) ListImages(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	imagesList, err := h.service.GetUserImages(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	imageResponseList := make([]resources.UserImageInfoResponse, len(imagesList))
	for _, image := range imagesList {
		imageResponseList = append(
			imageResponseList,
			resources.GetUserInfoResponse(image),
		)
	}

	response := resources.GetUserImagesListResponse(imageResponseList)
	c.JSON(http.StatusOK, response)
}
