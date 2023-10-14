package admin

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/resources/users/admin"
	models "legocy-go/internal/domain/users/models"
	"net/http"
)

// GetUsersAdmin
//
//	@Summary	Get Users (Admin)
//	@Tags		users_admin
//	@ID			list_users_admin
//	@Produce	json
//	@Success	200	{object}	[]admin.UserAdminDetailResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/users/ [get]
//
//	@Security	JWT
func (uah *UserAdminHandler) GetUsersAdmin(c *gin.Context) {
	var users []*models.UserAdmin

	users, err := uah.service.GetUsers(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var response = make([]admin.UserAdminDetailResponse, 0, len(users))
	for _, user := range users {
		response = append(response, admin.GetUserAdminDetailResponse(user))
	}

	c.JSON(http.StatusOK, response)
	return
}