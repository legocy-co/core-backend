package users

import (
	"github.com/gin-gonic/gin"
	resources "legocy-go/internal/delievery/http/resources/users"
	"net/http"
	"strconv"
)

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
