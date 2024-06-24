package user_collection

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	v1 "github.com/legocy-co/legocy/internal/delivery/http/middleware/auth"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/collections"
	"net/http"
)

// GetUserCollection
//
//	@Summary	Get User Collection
//	@Tags		user_collections
//	@ID			get_user_collection
//	@Produce	json
//	@Success	200	{object} 	collections.UserLegoSetCollectionResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/collections/ [get]
//
//	@Security	JWT
func (h UserLegoCollectionHandler) GetUserCollection(c *gin.Context) {
	tokenPayload, err := v1.GetUserPayload(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userID := tokenPayload.ID

	userCollection, appErr := h.s.GetUserCollection(c, userID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	collectionValuations, _, appErr := h.s.GetUserCollectionValuation(c, userID)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	userCollectionResponse := collections.GetUserLegoCollectionResponse(*userCollection, collectionValuations)

	c.JSON(http.StatusOK, userCollectionResponse)
}
