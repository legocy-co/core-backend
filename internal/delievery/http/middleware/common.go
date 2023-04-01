package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/errors"
	"legocy-go/internal/domain/users/middleware"
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
