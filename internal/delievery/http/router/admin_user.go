package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/delievery/http/handlers/users/admin"
	m "legocy-go/internal/delievery/http/middleware"
	s "legocy-go/internal/domain/users/service/admin"
)

func (r V1router) addUserAdmin(rg *gin.RouterGroup, service s.UserAdminService) {
	handler := h.NewUserAdminHandler(service)

	usersAdmin := rg.Group("/users").Use(m.AdminUserOnly())
	{
		usersAdmin.GET("/", handler.GetUsersAdmin)
		usersAdmin.GET("/:userId", handler.GetUserByID)
		usersAdmin.POST("/register", handler.AdminRegister)
		usersAdmin.PUT("/:userId", handler.UpdateUserByID)
		usersAdmin.DELETE("/:userId", handler.DeleteUser)
	}
}
