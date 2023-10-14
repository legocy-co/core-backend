package admin

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/resources/users/admin"
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
func (uah *UserAdminHandler) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	userDomain, err := uah.service.GetUserByID(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	userResponse := admin.GetUserAdminDetailResponse(userDomain)
	c.JSON(http.StatusOK, userResponse)
	return
}