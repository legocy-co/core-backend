package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/delievery/http/handlers/users/admin"
	m "legocy-go/internal/delievery/http/middleware"
	s "legocy-go/internal/domain/users/service/admin"
)

func (r V1router) addUserAdmin(rg *gin.RouterGroup, service s.UserAdminService) {
	handler := h.NewUserAdminHandler(service)

	userAdmin := rg.Group("/admin/users").Use(m.AdminUserOnly())
	{
		userAdmin.DELETE("/:userID", handler.DeleteUser)
	}
}
