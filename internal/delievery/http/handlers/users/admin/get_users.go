package admin

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/errors"
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
func (h *UserAdminHandler) GetUsersAdmin(c *gin.Context) {
	var users []*models.UserAdmin

	users, err := h.service.GetUsers(c)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	var response = make([]admin.UserAdminDetailResponse, 0, len(users))
	for _, user := range users {
		response = append(response, admin.GetUserAdminDetailResponse(user))
	}

	c.JSON(http.StatusOK, response)
	return
}
