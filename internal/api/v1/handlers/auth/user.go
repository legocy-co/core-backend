package auth

import (
	"fmt"
	"legocy-go/internal/api/v1/resources"
	auth2 "legocy-go/internal/api/v1/resources/auth"
	ser "legocy-go/internal/api/v1/usecase/auth"
	jwt "legocy-go/pkg/auth/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	service ser.UserUseCase
}

func NewTokenHandler(service ser.UserUseCase) TokenHandler {
	return TokenHandler{service: service}
}

func (th *TokenHandler) GenerateToken(c *gin.Context) {

	var jwtRequest auth2.JWTRequest
	if err := c.ShouldBindJSON(&jwtRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	errIsValidUser := th.service.ValidateUser(c.Request.Context(), jwtRequest)
	if errIsValidUser != nil {
		v1.ErrorRespond(c.Writer, errIsValidUser.Error())
		return
	}

	user, err := th.service.GetUserByEmail(c.Request.Context(), jwtRequest.Email)
	if err != nil {
		v1.ErrorRespond(c.Writer, err.Error())
		return
	}

	token, err := jwt.GenerateJWT(user.Email, user.Role)
	if err != nil {
		fmt.Println(err)
		v1.ErrorRespond(c.Writer, "Error generating token")
		return
	}

	v1.Respond(c.Writer, v1.DataMetaResponse{
		Data: auth2.JWTResponse{
			AccessToken: token,
		},
		Meta: v1.SuccessMetaResponse,
	})

}

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

	response := v1.DataMetaResponse{
		Data: auth2.GetUserResponse(user),
		Meta: v1.SuccessMetaResponse,
	}
	v1.Respond(c.Writer, response)

}

// Admin handlers

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

	response := v1.DataMetaResponse{
		Data: auth2.GetUserResponse(user),
		Meta: v1.SuccessMetaResponse,
	}
	v1.Respond(c.Writer, response)
}
