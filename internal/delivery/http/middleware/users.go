package middleware

import (
	"legocy-go/internal/delivery/http/errors"
	models "legocy-go/internal/domain/users/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := GetAuthTokenHeader(ctx)
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "Token Header not found"})
			ctx.Abort()
			return
		}

		err := ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func AdminUserOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := GetAuthTokenHeader(ctx)
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "Token Header not found"})
			ctx.Abort()
			return
		}

		err := ValidateAdminToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func UserIdOrAdmin(lookUpParam string) gin.HandlerFunc {
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
