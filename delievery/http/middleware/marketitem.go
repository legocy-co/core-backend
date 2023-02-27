package v1

import (
	"github.com/gin-gonic/gin"
	auth "legocy-go/pkg/auth/middleware"
	models "legocy-go/pkg/auth/models"
	r "legocy-go/pkg/marketplace/repository"
	"net/http"
	"strconv"
)

// ItemOwnerOrAdmin checks if Current User is owner of MarketItem or Admin user
func ItemOwnerOrAdmin(
	lookUpParam string, repo r.MarketItemRepository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// Get Token Header
		tokenPayload, err := GetUserPayload(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if tokenPayload.Role == models.ADMIN {
			ctx.Next()
		}

		itemID, err := strconv.Atoi(ctx.Param(lookUpParam))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sellerID, err := repo.GetMarketItemSellerID(ctx, itemID)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if tokenPayload.ID != sellerID {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": "User does not have permission"})
			return
		}

		ctx.Next()
	}
}

// HasFreeMarketItemsSlot Checks if Given User
func HasFreeMarketItemsSlot(
	maxAmount int, repo r.MarketItemRepository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := GetAuthTokenHeader(ctx)
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "Token Header not found"})
			ctx.Abort()
			return
		}

		tokenPayload, ok := auth.ParseTokenClaims(tokenString)
		if !ok {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{" error": "Error parsing Token Claims"})
			return
		}

		userItemsCount, err := repo.GetSellerMarketItemsAmount(ctx, tokenPayload.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if user has free slots
		if userItemsCount >= int64(maxAmount) {
			ctx.AbortWithStatusJSON(
				http.StatusForbidden, gin.H{
					"error": "User has exceeded limit of publishing items"})
			return
		}

		ctx.Next()
	}
}
