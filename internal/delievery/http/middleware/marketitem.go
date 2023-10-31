package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	r "legocy-go/internal/domain/marketplace/repository"
	models "legocy-go/internal/domain/users/models"
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
			return
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

// IsMarketItemOwner checks if Current User is owner of MarketItem
func IsMarketItemOwner(
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
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": "user method"})
			return
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

		tokenPayload, ok := ParseTokenClaims(tokenString)
		if !ok {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{" error": "Error parsing Token Claims"})
			return
		}

		logrus.Info("Getting seller market item amount")
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
		logrus.Info("User has free slots to create market item")
		ctx.Next()
	}
}
