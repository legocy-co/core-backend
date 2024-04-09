package middleware

import (
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/pkg/auth/jwt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAuthTokenHeader(ctx *gin.Context) string {

	// Bearer <token> -> <token>
	// If Invalid -> ""

	value := ctx.GetHeader("Authorization")
	value = strings.TrimLeft(value, " ")

	if value == "" {
		return value
	}

	if len(value) < 7 {
		return ""
	}

	if value[:7] != "Bearer " {
		return ""
	}

	return strings.TrimLeft(value[7:], " ")

}

func GetUserPayload(ctx *gin.Context) (*jwt.JWTClaim, error) {
	tokenString := GetAuthTokenHeader(ctx)
	if tokenString == "" {
		return nil, errors.ErrInvaldTokenHeader
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

		log.Println("Checking user payload")
		tokenPayload, err := GetUserPayload(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get UserID param
		log.Println("Checking user id from url")
		userID, err := strconv.Atoi(ctx.Param(lookUpParam))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": errors.ErrParamNotFound.Error()})
			return
		}

		// check if User itself or admin
		if tokenPayload.ID != userID && tokenPayload.Role != models.ADMIN {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User does not have permission"})
			return
		}

		log.Println("User is owner or admin. Access granted.")
		ctx.Next()
	}
}
