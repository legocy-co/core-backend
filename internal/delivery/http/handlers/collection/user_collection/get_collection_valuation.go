package user_collection

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/collections"
	v1 "github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
	"net/http"
)

// GetUserCollectionValuation
//
//	@Summary	Get User Collection Valuation
//	@Tags		user_collections
//	@ID			get_user_collection_valuation
//	@Produce	json
//	@Success	200	{object} 	collections.UserCollectionValuationResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/collections/calculator/ [get]
//
//	@Security	JWT
func (h UserLegoCollectionHandler) GetUserCollectionValuation(c *gin.Context) {
	tokenPayload, err := v1.GetUserPayload(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := tokenPayload.ID

	setsValuations, user, appErr := h.s.GetUserCollectionValuation(c, userID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	response := collections.FromUserCollectionValuation(setsValuations, *user)
	c.JSON(http.StatusOK, response)
}
