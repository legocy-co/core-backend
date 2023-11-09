package legoset

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/errors"
	"net/http"
	"strconv"
)

// SetDelete
//
//	@Summary	Delete Lego Set object
//	@Tags		lego_sets_admin
//	@ID			set_delete
//	@Param		setID	path	int	true	"LegoSet ID"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/sets/{setID} [delete]
//
//	@Security	JWT
func (lsh *LegoSetHandler) SetDelete(c *gin.Context) {
	setID, _err := strconv.Atoi(c.Param("setID"))
	if _err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	err := lsh.service.LegoSetDelete(c, setID)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
