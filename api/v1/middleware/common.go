package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/api/v1/errors"
	auth "legocy-go/pkg/auth/middleware"
)

func GetAuthTokenHeader(ctx *gin.Context) string {
	return ctx.GetHeader("Authorization")
}

func GetUserPayload(ctx *gin.Context) (*auth.JWTClaim, error) {
	tokenString := GetAuthTokenHeader(ctx)
	if tokenString == "" {
		return nil, errors.ErrTokenHeaderNotFound
	}

	tokenPayload, ok := auth.ParseTokenClaims(tokenString)
	if !ok {
		return nil, errors.ErrParsingClaims
	}

	return tokenPayload, nil
}
