package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/legocy-co/legocy/docs"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/pkg/auth/jwt"
	"net/http"
)

// UserRegister
//
//	@Summary	register new user
//	@Tags		authentication
//	@ID			user-register
//	@Produce	json
//	@Param		data	body		schemas.UserRegistrationRequest	true	"user data"
//	@Success	200		{object}	schemas.JWTResponse
//	@Failure	400		{object}	map[string]interface{}
//	@Router		/users/auth/register [post]
func (th *TokenHandler) UserRegister(c *gin.Context) {

	var registerReq schemas.UserRegistrationRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := registerReq.ToUser()
	if err := th.service.CreateUser(c, user, registerReq.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
