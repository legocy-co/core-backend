package google

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/pkg/auth/jwt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/idtoken"
)

// SignUp godoc
// @Summary Sign up with Google
// @Description Sign up with Google
//
//	@ID			sign-up-google
//
//	@Tags		authentication
//
// @Accept  json
// @Produce  json
// @Param data body schemas.GoogleSignInUpRequest true "Google sign up request"
// @Success 200 {object} schemas.JWTResponse
// @Failure 400 {object} map[string]interface{}
// @Router /users/auth/google/sign-up [post]
func (h Handler) SignUp(ctx *gin.Context) {

	var req schemas.GoogleSignInUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": err.Error()})
		return
	}

	cfg := config.GetAppConfig()
	payload, err := idtoken.Validate(ctx, req.Token, cfg.GoogleClientID)
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	vo := schemas.FromGoogleToken(payload)
	if appErr := h.r.CreateUser(ctx, vo); appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	go h.saveGoogleImage(payload)

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

func (h Handler) saveGoogleImage(payload *idtoken.Payload) {

	user, appErr := h.r.GetByExternalID(context.Background(), payload.Subject)
	if appErr != nil {
		return
	}

	err := h.uploadImage(user, payload.Claims["picture"].(string))
	if err != nil {
		log.Errorf("Failed to save user image: error - %v", err.Error())
	}

}
