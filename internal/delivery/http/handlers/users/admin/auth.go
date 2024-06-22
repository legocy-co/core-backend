package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/pkg/jwt"
	"net/http"
)

// LoginAdmin
//
//	@Summary	generate jwt token admin
//	@Tags		authentication_admin
//	@ID			create-jwt-admin
//	@Produce	json
//	@Param		data	body		schemas.SignInRequest	true	"jwt request"
//	@Success	200		{object}	schemas.JWTResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/admin/users/auth/sign-in [post]
func (h UserAdminHandler) LoginAdmin(c *gin.Context) {

	var jwtRequest schemas.SignInRequest
	if err := c.ShouldBindJSON(&jwtRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, appErr := h.service.LoginAdmin(c.Request.Context(), jwtRequest.Email, jwtRequest.Password)
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

	c.JSON(http.StatusOK, schemas.JWTResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

}
