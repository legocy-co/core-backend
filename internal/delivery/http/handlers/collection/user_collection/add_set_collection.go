package user_collection

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delivery/http/errors"
	"legocy-go/internal/delivery/http/resources/collections"
	v1 "legocy-go/pkg/auth/jwt/middleware"
	"net/http"
)

// AddLegoSetToUserCollection
//
//	@Summary	Add Set To Collection
//	@Tags		user_collections
//	@ID			add_set_user_collections
//	@Param		data	body	collections.CollectionLegoSetAddRequest	true	"data"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/collections/ [post]
//
//	@Security	JWT
func (h UserLegoCollectionHandler) AddLegoSetToUserCollection(c *gin.Context) {
	tokenPayload, err := v1.GetUserPayload(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := tokenPayload.ID

	var createRequest *collections.CollectionLegoSetAddRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	vo, err := createRequest.ToCollectionLegoSetValueObject()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	appErr := h.s.AddSetToUserCollection(c, userID, *vo)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
