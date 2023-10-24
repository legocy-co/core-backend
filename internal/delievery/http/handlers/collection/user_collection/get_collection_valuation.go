package user_collection

import (
	"github.com/gin-gonic/gin"
	v1 "legocy-go/internal/delievery/http/middleware"
	"legocy-go/internal/delievery/http/resources/collections"
	"net/http"
	"strconv"
)

// GetUserCollectionValuation
//
//	@Summary	Get User Collection Valuation
//	@Tags		user_collections
//	@ID			get_user_collection_valuation
//	@Produce	json
//	@Success	200	{object} 	collections.UserCollectionValuationResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Param		currencyID	path	int	true	"currency ID"
//	@Router		/collections/valuation/{currencyID} [get]
//
//	@Security	JWT
func (h UserLegoCollectionHandler) GetUserCollectionValuation(c *gin.Context) {
	tokenPayload, err := v1.GetUserPayload(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currencyId, err := strconv.Atoi(c.Param("currencyID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Currency ID"})
		return
	}

	userID := tokenPayload.ID

	setsValuations, user, err := h.s.GetUserCollectionValuation(c, userID, currencyId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := collections.FromUserCollectionValuation(setsValuations, *user)
	c.JSON(http.StatusOK, response)
}
