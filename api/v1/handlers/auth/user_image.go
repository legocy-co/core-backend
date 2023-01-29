package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ser "legocy-go/api/v1/usecase/auth"
	"legocy-go/internal/storage"
	"legocy-go/internal/storage/models"
	"net/http"
	"strconv"
)

type UserImageHandler struct {
	service ser.UserImageUseCase
	storage storage.ImageStorage
}

func (h *UserImageHandler) UploadUserImage(c *gin.Context) {
	//TODO: add permission(user itself or admin) & store imgUrl in database

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

	img := models.ImageUnit{
		UserID:      userID,
		Payload:     src,
		PayloadName: file.Filename,
		PayloadSize: file.Size,
	}

	imgUrl, err := h.storage.UploadFile(c.Request.Context(), img)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
	}

	logrus.Info(imgUrl)
}
