package auth

import (
	r "legocy-go/delievery/http/resources"
	"legocy-go/delievery/http/resources/auth"
	ser "legocy-go/delievery/http/usecase/auth"
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

// Generate Token
// @Summary      JWT Token Generation
// @Description  generate Token by
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /token [post]
func (th *TokenHandler) GenerateToken(c *gin.Context) {

	var jwtRequest auth.JWTRequest
	if err := c.ShouldBindJSON(&jwtRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	err := th.service.ValidateUser(c.Request.Context(), jwtRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	user, err := th.service.GetUserByEmail(c.Request.Context(), jwtRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	token, err := jwt.GenerateJWT(user.Email, user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
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
