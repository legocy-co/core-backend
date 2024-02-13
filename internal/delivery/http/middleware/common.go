package middleware

import (
	"github.com/gin-gonic/gin"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func IsOwnerOrAdmin(lookupParam string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Get Token Header
		tokenPayload, err := middleware.GetUserPayload(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if tokenPayload.Role == models.ADMIN {
			log.Println("Admin user. Access granted")
			ctx.Next()
			return
		}

		paramToCheck, err := strconv.Atoi(ctx.Param(lookupParam))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if tokenPayload.ID != paramToCheck {
			ctx.AbortWithStatusJSON(
				http.StatusForbidden, gin.H{"error": "User does not have permission"})
			return
		}

		log.Println("User is owner. Access granted")
		ctx.Next()
	}
}
