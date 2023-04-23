package legoset

import (
	"github.com/gin-gonic/gin"
	v1 "legocy-go/internal/delievery/http/resources"
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
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	err = lsh.service.LegoSetDelete(c, setID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	response := v1.DataMetaResponse{
		Data: true,
		Meta: v1.SuccessMetaResponse,
	}

	v1.Respond(c.Writer, response)
}
