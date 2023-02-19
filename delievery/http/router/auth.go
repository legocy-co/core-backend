package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/delievery/http/handlers/auth"
	m "legocy-go/delievery/http/middleware"
	s "legocy-go/delievery/http/usecase/auth"
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
