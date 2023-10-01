package admin

import (
	"github.com/gin-gonic/gin"
	service "legocy-go/internal/domain/users/service/admin"
	"net/http"
	"strconv"
)

type UserAdminHandler struct {
	service service.UserAdminService
}

func NewUserAdminHandler(service service.UserAdminService) UserAdminHandler {
	return UserAdminHandler{service: service}
}

// DeleteUser
//
// DeleteUserAdmin
//
//	@Summary	Delete User
//	@Tags		user_admin
//	@ID			delete_user
//	@Param		userID	path	int	true	"user ID"
//	@Produce	json
//	@Success	200	{object}	map[string]bool
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/users/{userID} [delete]
//
//	@Security	JWT
func (uah *UserAdminHandler) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	err = uah.service.DeleteUser(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			"Error deleting User object")
	}

	c.JSON(http.StatusOK, map[string]bool{"status": true})
}
