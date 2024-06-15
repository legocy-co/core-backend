package fb

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/pkg/auth/jwt"
	"github.com/legocy-co/legocy/pkg/facebook"
)

// SignUp godoc
// @Summary Sign up with Facebook
// @Description Sign up with Facebook
// @ID sign-up-facebook
// @Tags authentication
// @Router /users/auth/fb/sign-up [get]
func (h Handler) SignUp(ctx *gin.Context) {
	ctx.Redirect(307, facebook.GetOAuthConfig(false).AuthCodeURL("state"))
}

// SignUpCallback godoc
// @Summary Sign up with Facebook callback
// @Description Sign up with Facebook callback
// @ID sign-up-facebook-callback
// @Tags authentication
// @Router /users/auth/fb/sign-up/callback [get]
func (h Handler) SignUpCallback(ctx *gin.Context) {

	cfg := facebook.GetOAuthConfig(false)
	code := ctx.Query("code")

	payload, err := facebook.GetUserInfo(ctx, cfg, code)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	vo := payload.ToUserVO()
	if appErr := h.r.CreateUser(ctx, vo); appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	user, appErr := h.r.GetByExternalID(ctx, payload.ID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
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
