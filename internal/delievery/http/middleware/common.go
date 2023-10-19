package middleware

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/errors"
	"legocy-go/internal/domain/users/middleware"
	models "legocy-go/internal/domain/users/models"
	"net/http"
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

func IsAdminUser() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := GetAuthTokenHeader(ctx)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "provide a Authorization header"})
			return
		}

		tokenPayload, ok := auth.ParseTokenClaims(tokenString)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token data"})
			return
		}

		if tokenPayload.Role != models.ADMIN {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied(admin users only)"})
			return
		}

		ctx.Next()
	}

}
