package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeleteUser
//
// DeleteUserAdmin
//
//	@Summary	Delete User
//	@Tags		users_admin
//	@ID			delete_user_admin
//	@Param		userId	path	int	true	"user ID"
//	@Produce	json
//	@Success	200	{object}	map[string]bool
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/users/{userId} [delete]
//
//	@Security	ApiKeyAuth
//	@param		Authorization	header	string	true	"Authorization"
func (uah *UserAdminHandler) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	err = uah.service.DeleteUser(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, map[string]bool{"status": true})
}
