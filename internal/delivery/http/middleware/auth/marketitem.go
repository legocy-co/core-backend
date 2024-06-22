package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	r "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/pkg/jwt"
	log "github.com/sirupsen/logrus"
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

		sellerID, e := repo.GetMarketItemSellerID(ctx, itemID)
		if e != nil {
			httpErr := errors.FromAppError(*e)
			ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
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
				http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("Token Payload: %v", err.Error())},
			)
			return
		}

		//if tokenPayload.Role == models.ADMIN {
		//	ctx.AbortWithStatusJSON(
		//		http.StatusBadRequest, gin.H{"error": "user method"})
		//	return
		//}

		itemID, err := strconv.Atoi(ctx.Param(lookUpParam))
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": "Invalid ID path param value"})
			return
		}

		sellerID, e := repo.GetMarketItemSellerID(ctx, itemID)
		if e != nil {
			httpErr := errors.FromAppError(*e)
			ctx.AbortWithStatusJSON(
				httpErr.Status, fmt.Sprintf("sellerID: %v", httpErr.Message),
			)
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

		tokenPayload, ok := jwt.ParseTokenClaims(tokenString, config.GetAppConfig().JwtConf.SecretKey)
		if !ok {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"error": "Error parsing Token Claims"})
			return
		}

		log.Info("Checking if user has free slots to create market item")
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
		log.Info("User has free slots to create market item")
		ctx.Next()
	}
}
