package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users/admin"
	"net/http"
	"strconv"
)

// GetUserByID
//
//	@Summary	Get User by ID (Admin)
//	@Tags		users_admin
//	@ID			detail_user_admin
//	@Param		userId	path	int	true	"user ID"
//	@Produce	json
//	@Success	200	{object}	admin.UserAdminDetailResponse
//	@Failure	404	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/users/{userId} [get]
//
//	@Security	JWT
func (h *UserAdminHandler) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	userDomain, appErr := h.service.GetUserByID(c, userID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
	}

	userResponse := admin.GetUserAdminDetailResponse(userDomain)
	c.JSON(http.StatusOK, userResponse)
	return
}
