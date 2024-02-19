package legoset

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/lego"
	"net/http"
	"strconv"
)

// SetUpdate
//
//	@Summary	Update Lego Set object
//	@Tags		lego_sets_admin
//	@ID			set_update
//	@Param		data	body	lego.LegoSetRequest	true	"create data"
//	@Param		setID	path	int	true	"Lego Set ID"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/sets/{setID} [put]

// @Security	JWT
func (h *LegoSetHandler) SetUpdate(c *gin.Context) {
	var setRequest lego.LegoSetRequest
	if _err := c.ShouldBindJSON(&setRequest); _err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": _err.Error()})
		return
	}

	setID, _err := strconv.Atoi(c.Param("setID"))
	if _err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	legoSetValueObject := setRequest.ToLegoSeriesValueObject()
	err := h.service.LegoSetUpdate(setID, legoSetValueObject)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
