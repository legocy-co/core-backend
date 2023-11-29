package user_collection

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delivery/http/errors"
	"legocy-go/internal/delivery/http/resources/collections"
	v1 "legocy-go/pkg/auth/jwt/middleware"
	"net/http"
	"strconv"
)

// UpdateUserCollectionSet
//
//	@Summary	Update Collection Set Info
//	@Tags		user_collections
//	@ID			update_set_user_collection
//	@Param		setID	path	int	true	"set ID"
//	@Param		data	body	collections.CollectionLegoSetUpdateRequest	true	"data"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/collections/{setID} [put]
//
//	@Security	JWT
func (h UserLegoCollectionHandler) UpdateUserCollectionSet(c *gin.Context) {
	tokenPayload, err := v1.GetUserPayload(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := tokenPayload.ID

	collectionSetId, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error extracting set ID from URL"})
		return
	}

	var updateRequest *collections.CollectionLegoSetUpdateRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	vo, err := updateRequest.ToCollectionLegoSetValueObject()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	appErr := h.s.UpdateUserCollectionSet(c, userID, collectionSetId, *vo)
	if appErr != nil {
		httpErr := errors.FromAppError(*appErr)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})

}
