package auth

import (
	"github.com/gin-gonic/gin"
	"legocy-go/config"
	_ "legocy-go/docs"
	"legocy-go/internal/delivery/http/errors"
	resources "legocy-go/internal/delivery/http/resources/users"
	"legocy-go/pkg/auth/jwt"
	"net/http"
)

// GenerateToken
//
//	@Summary	generate jwt token
//	@Tags		authentication
//	@ID			create-jwt
//	@Produce	json
//	@Param		data	body		resources.SignInRequest	true	"jwt request"
//	@Success	200		{object}	resources.JWTResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/auth/sign-in [post]
func (th *TokenHandler) GenerateToken(c *gin.Context) {

	var jwtRequest resources.SignInRequest
	if err := c.ShouldBindJSON(&jwtRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	appErr := th.service.ValidateUserCredentials(c.Request.Context(), jwtRequest)
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

	accessToken, err := jwt.GenerateAccessToken(
		user.Email,
		user.ID,
		user.Role,
		config.GetAppConfig().JwtConf.SecretKey,
		config.GetAppConfig().JwtConf.AccessTokenLifeTime,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(
		user.Email,
		user.ID,
		user.Role,
		config.GetAppConfig().JwtConf.SecretKey,
		config.GetAppConfig().JwtConf.RefreshTokenLifeTime,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resources.JWTResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

}
