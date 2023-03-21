package auth

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/resources"
	ser "legocy-go/internal/delievery/http/service/auth"
	models "legocy-go/internal/domain/auth/models"
	"legocy-go/internal/storage"
	s "legocy-go/internal/storage/models"
	"legocy-go/internal/storage/provider"
	"net/http"
	"strconv"
)

type UserImageHandler struct {
	service ser.UserImageUseCase
	storage storage.ImageStorage
}

func NewUserImageHandler(service ser.UserImageUseCase, storage storage.ImageStorage) UserImageHandler {
	return UserImageHandler{service: service, storage: storage}
}

func (h *UserImageHandler) UploadUserImage(c *gin.Context) {

	err := h.storage.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	src, err := file.Open()
	defer src.Close()

	img := s.ImageUnit{
		ID:          userID,
		Payload:     src,
		PayloadName: file.Filename,
		PayloadSize: file.Size,
	}

	imgUrl, err := h.storage.UploadFile(c.Request.Context(), img, provider.UserObjectsBucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	userImg := &models.UserImage{
		UserID:      userID,
		FilepathURL: imgUrl,
	}

	err = h.service.StoreUserImage(c.Request.Context(), userImg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	response := v1.DataMetaResponse{Data: userImg, Meta: v1.SuccessMetaResponse}
	v1.Respond(c.Writer, response)
}
