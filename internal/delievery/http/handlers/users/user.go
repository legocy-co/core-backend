package users

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	auth2 "legocy-go/internal/delievery/http/resources/users"
	ser "legocy-go/internal/delievery/http/service/users"
	jwt "legocy-go/internal/domain/users/middleware"
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
//	@Param		data	body		users.JWTRequest	true	"jwt request"
//	@Success	200		{object}	users.JWTResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/token [post]
func (th *TokenHandler) GenerateToken(c *gin.Context) {

	var jwtRequest auth2.JWTRequest
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

	c.JSON(http.StatusOK, auth2.JWTResponse{
		AccessToken: token,
	})

}

// UserRegister
//
//	@Summary	register new user
//	@Tags		authentication
//	@ID			user-register
//	@Produce	json
//	@Param		data	body		users.UserRegistrationRequest	true	"user data"
//	@Success	200		{object}	users.UserRegistrationResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/register [post]
func (th *TokenHandler) UserRegister(c *gin.Context) {

	var registerReq auth2.UserRegistrationRequest

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

	c.JSON(http.StatusOK, auth2.GetUserResponse(user))
}

// AdminRegister
//
//	@Summary	Create Admin User
//	@Tags		authentication_admin
//	@ID			create-admin
//	@Produce	json
//	@Param		data	body		users.UserRegistrationRequest	true	"reg request"
//	@Success	200		{object}	users.UserRegistrationResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/admin/users [post]
//
//	@Security	ApiKeyAuth
//	@param		Authorization	header	string	true	"Authorization"
func (th *TokenHandler) AdminRegister(c *gin.Context) {

	var registerReq auth2.UserRegistrationRequest

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

	c.JSON(http.StatusOK, auth2.GetUserResponse(user))
}
