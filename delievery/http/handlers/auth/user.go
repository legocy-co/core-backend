package auth

import (
	"github.com/gin-gonic/gin"
	"legocy-go/delievery/http/resources/auth"
	ser "legocy-go/delievery/http/usecase/auth"
	_ "legocy-go/docs"
	jwt "legocy-go/pkg/auth/middleware"
	"net/http"
)

type TokenHandler struct {
	service ser.UserUseCase
}

func NewTokenHandler(service ser.UserUseCase) TokenHandler {
	return TokenHandler{service: service}
}

// GenerateToken
//
//	@Summary	generate jwt token
//	@Tags		authentication
//	@ID			create-jwt
//	@Produce	json
//	@Param		data	body		auth.JWTRequest	true	"jwt request"
//	@Success	200		{object}	auth.JWTResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/auth/token [post]
func (th *TokenHandler) GenerateToken(c *gin.Context) {

	var jwtRequest auth.JWTRequest
	if err := c.ShouldBindJSON(&jwtRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	err := th.service.ValidateUser(c.Request.Context(), jwtRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := th.service.GetUserByEmail(c.Request.Context(), jwtRequest.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateJWT(user.Email, user.ID, user.Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auth.JWTResponse{
		AccessToken: token,
	})

}

// UserRegister
//
//	@Summary	register new user
//	@Tags		authentication
//	@ID			user-register
//	@Produce	json
//	@Param		data	body		auth.UserRegistrationRequest	true	"user data"
//	@Success	200		{object}	auth.UserRegistrationResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/auth/register [post]
func (th *TokenHandler) UserRegister(c *gin.Context) {

	var registerReq auth.UserRegistrationRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := registerReq.ToUser()
	if err := th.service.CreateUser(c, user, registerReq.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, auth.GetUserResponse(user))
}

// AdminRegister
//
//	@Summary	Create Admin User
//	@Tags		authentication
//	@ID			create-admin
//	@Produce	json
//	@Param		data	body		auth.UserRegistrationRequest	true	"reg request"
//	@Success	200		{object}	auth.UserRegistrationResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/admin/auth [post]
func (th *TokenHandler) AdminRegister(c *gin.Context) {

	var registerReq auth.UserRegistrationRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := registerReq.ToAdmin()
	if err := th.service.CreateUser(c, user, registerReq.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, auth.GetUserResponse(user))
}
