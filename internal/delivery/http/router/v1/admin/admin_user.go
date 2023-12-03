package admin

import (
	"github.com/gin-gonic/gin"
	h "github.com/legocy-co/legocy/internal/delivery/http/handlers/users/admin"
	s "github.com/legocy-co/legocy/internal/domain/users/service/admin"
	m "github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
)

func AddUserAdmin(rg *gin.RouterGroup, service s.UserAdminService) {
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
