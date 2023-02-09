package auth

import (
	"github.com/gin-gonic/gin"
	r "legocy-go/api/v1/resources"
	ser "legocy-go/api/v1/usecase/auth"
	"legocy-go/internal/storage"
	s "legocy-go/internal/storage/models"
	models "legocy-go/pkg/auth/models"
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
		UserID:      userID,
		Payload:     src,
		PayloadName: file.Filename,
		PayloadSize: file.Size,
	}

	imgUrl, err := h.storage.UploadFile(c.Request.Context(), img)
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

	response := r.DataMetaResponse{Data: userImg, Meta: r.SuccessMetaResponse}
	r.Respond(c.Writer, response)
}
