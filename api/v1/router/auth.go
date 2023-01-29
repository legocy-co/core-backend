package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/api/v1/handlers/auth"
	m "legocy-go/api/v1/middleware"
	s "legocy-go/api/v1/usecase/auth"
)

func (r V1router) addAuth(rg *gin.RouterGroup, service s.UserUseCase) {
	handler := auth.NewTokenHandler(service)

	auth := rg.Group("/auth")
	{
		auth.POST("/token", handler.GenerateToken)
		auth.POST("/register", handler.UserRegister)
	}

	authPrivate := rg.Group("/admin/auth").Use(m.AdminUserOnly())
	{
		authPrivate.POST("/", handler.AdminRegister)
	}
}
