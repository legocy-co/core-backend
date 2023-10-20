package user_collection

import (
	"github.com/gin-gonic/gin"
	v1 "legocy-go/internal/delievery/http/middleware"
	"net/http"
	"strconv"
)

// DeleteUserCollectionLegoSet
//
//	@Summary	Delete User Collection Set
//	@Tags		user_collections
//	@ID			delete_set_user_collections
//	@Param		setID	path	int	true	"set ID"
//	@Produce	json
//	@Success	200	{object}	map[string]bool
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/collections/{setID} [delete]
//
//	@Security	JWT
func (h UserLegoCollectionHandler) DeleteUserCollectionLegoSet(c *gin.Context) {
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

	err = h.s.RemoveSetFromUserCollection(c, userID, collectionSetId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
