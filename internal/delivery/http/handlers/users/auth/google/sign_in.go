package google

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/pkg/auth/jwt"
	"google.golang.org/api/idtoken"
)

// SignIn godoc
// @Summary Sign in with Google
// @Description Sign in with Google
// @ID sign-in-google
// @Tags authentication
// @Accept json
// @Produce json
// @Param data body schemas.GoogleSignInUpRequest true "Google sign in request"
// @Success 200 {object} schemas.JWTResponse
// @Failure 400 {object} map[string]interface{}
// @Router /users/auth/google/sign-in [post]
func (h Handler) SignIn(ctx *gin.Context) {

	var req schemas.GoogleSignInUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": err.Error()})
		return
	}

	cfg := config.GetAppConfig()
	payload, err := idtoken.Validate(ctx, req.Token, cfg.GoogleClientID)
	if err != nil {
		ctx.AbortWithStatusJSON(403, gin.H{"error": err.Error()})
		return
	}

	user, appErr := h.r.GetByExternalID(ctx, payload.Subject)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
	}

	accessToken, err := jwt.GenerateAccessToken(
		user.Email,
		user.ID,
		user.Role,
		cfg.JwtConf.SecretKey,
		cfg.JwtConf.AccessTokenLifeTime,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(
		user.Email,
		user.ID,
		user.Role,
		cfg.JwtConf.SecretKey,
		cfg.JwtConf.RefreshTokenLifeTime,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, schemas.JWTResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}
