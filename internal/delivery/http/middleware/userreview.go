package middleware

import (
	"github.com/gin-gonic/gin"
	r "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
	"net/http"
	"strconv"
)

// ReviewOwnerOrAdmin checks if Current User is owner of User Review or Admin user
func ReviewOwnerOrAdmin(
	lookUpParam string, repo r.UserReviewRepository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// Get Token Header
		tokenPayload, err := middleware.GetUserPayload(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if tokenPayload.Role == models.ADMIN {
			ctx.Next()
		}

		reviewID, err := strconv.Atoi(ctx.Param(lookUpParam))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		reviewerID, err := repo.GetReviewerID(ctx, reviewID)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if tokenPayload.ID != reviewerID {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": "User does not have permission"})
			return
		}

		ctx.Next()
	}
}
