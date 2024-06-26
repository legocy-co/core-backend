package auth

import (
	"github.com/gin-gonic/gin"
	r "github.com/legocy-co/legocy/internal/domain/collections/repository"
	"github.com/legocy-co/legocy/internal/domain/users/models"
	"net/http"
	"strconv"
)

func CollectionSetOwnerOrAdmin(
	lookUpParam string, repo r.UserCollectionRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenPayload, err := GetUserPayload(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if tokenPayload.Role == models.ADMIN {
			ctx.Next()
			return
		}

		setID, err := strconv.Atoi(ctx.Param(lookUpParam))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "cannot extract set ID from URL"})
			return
		}

		setOwnerID, e := repo.GetCollectionSetOwner(ctx, setID)
		if e != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if tokenPayload.ID != setOwnerID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
			return
		}

		ctx.Next()
	}
}
