package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/delivery/http/handlers/users/admin"
	s "legocy-go/internal/domain/users/service/admin"
	m "legocy-go/pkg/auth/jwt/middleware"
)

func (r V1router) addUserAdmin(rg *gin.RouterGroup, service s.UserAdminService) {
	handler := h.NewUserAdminHandler(service)

	usersAdmin := rg.Group("/users").Use(m.IsAdmin())
	{
		usersAdmin.GET("/", handler.GetUsersAdmin)
		usersAdmin.GET("/:userId", handler.GetUserByID)
		usersAdmin.POST("/register", handler.AdminRegister)
		usersAdmin.PUT("/:userId", handler.UpdateUserByID)
		usersAdmin.DELETE("/:userId", handler.DeleteUser)
	}
}
