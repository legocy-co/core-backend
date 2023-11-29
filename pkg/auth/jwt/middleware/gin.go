package middleware

import (
	"legocy-go/config"
	"legocy-go/internal/delivery/http/errors"
	models "legocy-go/internal/domain/users/models"
	"legocy-go/pkg/auth/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAuthTokenHeader(ctx *gin.Context) string {
	return ctx.GetHeader("Authorization")
}

func GetUserPayload(ctx *gin.Context) (*jwt.JWTClaim, error) {
	tokenString := GetAuthTokenHeader(ctx)
	if tokenString == "" {
		return nil, errors.ErrTokenHeaderNotFound
	}

	tokenPayload, ok := jwt.ParseTokenClaims(tokenString, config.GetAppConfig().JwtConf.SecretKey)
	if !ok {
		return nil, errors.ErrParsingClaims
	}

	return tokenPayload, nil
}

func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := GetAuthTokenHeader(ctx)
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "Token Header not found"})
			ctx.Abort()
			return
		}

		err := jwt.ValidateAccessToken(tokenString, config.GetAppConfig().JwtConf.SecretKey)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := GetAuthTokenHeader(ctx)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Token Header not found"})
			return
		}

		err := jwt.ValidateAdminAccessToken(tokenString, models.ADMIN, config.GetAppConfig().JwtConf.SecretKey)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		ctx.Next()
	}
}

func IsOwnerOrAdmin(lookUpParam string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenPayload, err := GetUserPayload(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get UserID param
		userID, err := strconv.Atoi(ctx.Param(lookUpParam))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": errors.ErrParamNotFound.Error()})
			return
		}

		// check if User itself or admin
		if tokenPayload.ID != userID || tokenPayload.Role != models.ADMIN {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "User does not have permission"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
