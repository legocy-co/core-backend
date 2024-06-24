package admin

import (
	"github.com/gin-gonic/gin"
	h "github.com/legocy-co/legocy/internal/delivery/http/handlers/users/admin"
	m "github.com/legocy-co/legocy/internal/delivery/http/middleware/auth"
	s "github.com/legocy-co/legocy/internal/domain/users/service/admin"
)

func AddUserAdmin(rg *gin.RouterGroup, service s.UserAdminService) {
	handler := h.NewUserAdminHandler(service)

	authAdmin := rg.Group("/users/auth")
	{
		authAdmin.POST("/sign-in", handler.LoginAdmin)
	}

	usersAdmin := rg.Group("/users").Use(m.IsAdmin())
	{
		usersAdmin.GET("/", handler.GetUsersAdmin)
		usersAdmin.GET("/:userId", handler.GetUserByID)
		usersAdmin.POST("/register", handler.AdminRegister)
		usersAdmin.PUT("/:userId", handler.UpdateUserByID)
		usersAdmin.DELETE("/:userId", handler.DeleteUser)
	}
}
