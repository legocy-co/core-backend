package middleware

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delivery/http/errors"
)

func GetAuthTokenHeader(ctx *gin.Context) string {
	return ctx.GetHeader("Authorization")
}

func GetUserPayload(ctx *gin.Context) (*JWTClaim, error) {
	tokenString := GetAuthTokenHeader(ctx)
	if tokenString == "" {
		return nil, errors.ErrTokenHeaderNotFound
	}

	tokenPayload, ok := ParseTokenClaims(tokenString)
	if !ok {
		return nil, errors.ErrParsingClaims
	}

	return tokenPayload, nil
}
