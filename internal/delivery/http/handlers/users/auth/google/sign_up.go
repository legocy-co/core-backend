package google

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"google.golang.org/api/idtoken"
)

func (h *Handler) SignUp(ctx *gin.Context) {

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

	vo := schemas.FromGoogleToken(payload)
	if appErr := h.r.CreateUser(ctx, vo); appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	h.SignIn(ctx)
}
