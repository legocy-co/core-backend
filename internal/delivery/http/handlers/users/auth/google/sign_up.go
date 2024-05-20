package google

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/pkg/config"
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
func (h *Handler) SignUp(ctx *gin.Context) {

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

	h.SignIn(ctx)
}
