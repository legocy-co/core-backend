package userImage

import (
	"github.com/gin-gonic/gin"
	resources "legocy-go/internal/delivery/http/resources/users"
	"net/http"
	"strconv"
)

// ListImages
//
//	@Summary	Get User Images List
//	@Tags		users_images
//	@ID			get_user_images
//	@Produce	json
//	@Param		userID	path int	true	"user ID"
//	@Success	200		{object}	resources.UserImagesListResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/images/:userID [get]
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