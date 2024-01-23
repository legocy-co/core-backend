package userImage

import (
	"github.com/gin-gonic/gin"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
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
//	@Success	200		{object}	schemas.UserImagesListResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/images/{userID} [get]
func (h UserImageHandler) ListImages(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}

	imagesList, err := h.service.GetUserImages(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
	}

	imageResponseList := make([]schemas.UserImageInfoResponse, len(imagesList))
	for _, image := range imagesList {
		imageResponseList = append(
			imageResponseList,
			schemas.GetUserImageResponse(image),
		)
	}

	response := schemas.GetUserImagesListResponse(imageResponseList)
	c.JSON(http.StatusOK, response)
}
