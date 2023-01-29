package auth

import (
	"fmt"
	r "legocy-go/api/v1/resources"
	"legocy-go/api/v1/resources/auth"
	ser "legocy-go/api/v1/usecase/auth"
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

	var jwtRequest auth.JWTRequest
	if err := c.ShouldBindJSON(&jwtRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	errIsValidUser := th.service.ValidateUser(c.Request.Context(), jwtRequest)
	if errIsValidUser != nil {
		r.ErrorRespond(c.Writer, errIsValidUser.Error())
		return
	}

	user, err := th.service.GetUserByEmail(c.Request.Context(), jwtRequest.Email)
	if err != nil {
		r.ErrorRespond(c.Writer, err.Error())
		return
	}

	token, err := jwt.GenerateJWT(user.Email, user.Role)
	if err != nil {
		fmt.Println(err)
		r.ErrorRespond(c.Writer, "Error generating token")
		return
	}

	r.Respond(c.Writer, r.DataMetaResponse{
		Data: auth.JWTResponse{
			AccessToken: token,
		},
		Meta: r.SuccessMetaResponse,
	})

}

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

	response := r.DataMetaResponse{
		Data: auth.GetUserResponse(user),
		Meta: r.SuccessMetaResponse,
	}
	r.Respond(c.Writer, response)

}

// Admin handlers

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

	response := r.DataMetaResponse{
		Data: auth.GetUserResponse(user),
		Meta: r.SuccessMetaResponse,
	}
	r.Respond(c.Writer, response)
}
