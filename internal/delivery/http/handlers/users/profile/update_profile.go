package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"net/http"
	"strconv"
)

// UpdateUserProfile
//
//		@Summary	Update User Profile Page
//		@Tags		user_profile_pages
//		@ID			update_user_profile
//		@Param		userID	path	int	true	"user ID"
//	 	@Param		data	body	schemas.UserUpdateRequest	true	"User Update Request"
//		@Produce	json
//		@Success	200	{object}	map[string]interface{}
//		@Failure	400	{object}	map[string]interface{}
//		@Router		/users/profile/{userID} [put]
//
//		@Security	JWT
func (h UserProfilePageHandler) UpdateUserProfile(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	var req schemas.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	vo := req.ToUserValueObject()
	if appErr := h.userService.UpdateUser(userID, *vo); appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
}
