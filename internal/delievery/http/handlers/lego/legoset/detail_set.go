package legoset

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/resources/lego"
	"net/http"
	"strconv"
)

// SetDetail
//
//	@Summary	Get LegoSet by ID
//	@Tags		lego_sets
//	@ID			set_detail
//	@Param		setID	path	int	true	"Lego Set ID"
//	@Produce	json
//	@Success	200	{object}	lego.LegoSetResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/sets/{setID} [get]
//
//	@Security	JWT
func (lsh *LegoSetHandler) SetDetail(c *gin.Context) {
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		c.Abort()
		return
	}

	legoSet, err := lsh.service.LegoSetDetail(c.Request.Context(), setID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	legoSetResponse := lego.GetLegoSetResponse(legoSet)

	c.JSON(http.StatusOK, legoSetResponse)
}
