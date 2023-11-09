package auth

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	"legocy-go/internal/delievery/http/errors"
	jwt "legocy-go/internal/delievery/http/middleware"
	resources "legocy-go/internal/delievery/http/resources/users"
	"net/http"
)

// GenerateToken
//
//	@Summary	generate jwt token
//	@Tags		authentication
//	@ID			create-jwt
//	@Produce	json
//	@Param		data	body		resources.JWTRequest	true	"jwt request"
//	@Success	200		{object}	resources.JWTResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/auth/token [post]
func (th *TokenHandler) GenerateToken(c *gin.Context) {

	var jwtRequest resources.JWTRequest
	if err := c.ShouldBindJSON(&jwtRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	appErr := th.service.ValidateUser(c.Request.Context(), jwtRequest)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	user, appErr := th.service.GetUserByEmail(c.Request.Context(), jwtRequest.Email)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	token, err := jwt.GenerateJWT(user.Email, user.ID, user.Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resources.JWTResponse{
		AccessToken: token,
	})

}
