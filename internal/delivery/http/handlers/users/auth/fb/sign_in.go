package fb

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/pkg/jwt"
)

// SignIn godoc
// @Summary Sign in with Facebook callback
// @Description Sign in with Facebook callback
// @ID sign-in-facebook-callback
// @Tags authentication
// @Accept json
// @Produce json
// @Param data body schemas.FacebookSignInRequest true "Facebook sign in request"
// @Param X-Secret-Key header string true "Sign in key"
// @Success 200 {object} schemas.JWTResponse
// @Failure 400 {object} string
// @Failure 403 {object} string
// @Failure 500 {object} string
// @Router /users/auth/fb/sign-in [post]
func (h Handler) SignIn(ctx *gin.Context) {

	var data schemas.FacebookSignInRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	user, appErr := h.r.GetByExternalID(ctx, data.FacebookID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	if err := checkSignInKey(ctx, data, *user); err != nil {
		ctx.AbortWithStatusJSON(403, gin.H{"error": err.Error()})
		return
	}

	appCfg := config.GetAppConfig()

	accessToken, err := jwt.GenerateAccessToken(
		user.Email,
		user.ID,
		user.Role,
		appCfg.JwtConf.SecretKey,
		appCfg.JwtConf.AccessTokenLifeTime,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(
		user.Email,
		user.ID,
		user.Role,
		appCfg.JwtConf.SecretKey,
		appCfg.JwtConf.RefreshTokenLifeTime,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, schemas.JWTResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}
