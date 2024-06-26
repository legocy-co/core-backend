package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware/auth"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
)

// CurrentUserProfileHeader
//
//	@Summary	Get Current User Profile Info
//	@Tags		user_profile_pages
//	@ID			basic_user_profile_page
//	@Produce	json
//	@Success	200	{object}	users.UserDetailResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/users/profile/header/ [get]
//
//	@Security	JWT
func (h *UserProfilePageHandler) CurrentUserProfileHeader(ctx *gin.Context) {
	tokenPayload, err := auth.GetUserPayload(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, appErr := h.userService.GetUserByID(ctx, tokenPayload.ID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.JSON(httpErr.Status, httpErr)
		return
	}

	userResponse := users.GetUserDetailResponse(user)
	ctx.JSON(200, userResponse)
}
