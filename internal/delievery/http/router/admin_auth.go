package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/delievery/http/handlers/users/admin/auth"
	m "legocy-go/internal/delievery/http/middleware"
	s "legocy-go/internal/domain/users/service/admin"
)

func (r V1router) addAuthAdmin(rg *gin.RouterGroup, service s.UserAdminService) {
	handler := h.NewAdminHandler(service)

	authAdmin := rg.Group("/admin/users").Use(m.AdminUserOnly())
	{
		authAdmin.POST("/register", handler.AdminRegister)
	}
}
