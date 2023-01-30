package v1

import (
	auth "legocy-go/pkg/auth/middleware"
	models "legocy-go/pkg/auth/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "Token Header not found"})
			ctx.Abort()
			return
		}

		err := auth.ValidateToken(tokenString)
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
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "Token Header not found"})
			ctx.Abort()
			return
		}

		err := auth.ValidateAdminToken(tokenString)
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

		// Get Token Header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "Token Header not found"})
			ctx.Abort()
			return
		}

		// Get UserID param
		userID, err := strconv.Atoi(ctx.Param(lookUpParam))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		tokenPayload, ok := auth.ParseTokenClaims(tokenString)
		if !ok {

		}

		// check if User itself or admin
		if tokenPayload.ID != userID || tokenPayload.Role != models.ADMIN {
			ctx.JSON(http.StatusForbidden, "User does not have permission")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
