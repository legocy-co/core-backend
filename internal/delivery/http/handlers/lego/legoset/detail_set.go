package legoset

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/lego"
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
func (h *LegoSetHandler) SetDetail(c *gin.Context) {
	setID, _err := strconv.Atoi(c.Param("setID"))
	if _err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't extract ID from URL path"})
		return
	}

	legoSet, err := h.service.LegoSetDetail(c.Request.Context(), setID)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	legoSetResponse := lego.GetLegoSetResponse(legoSet)

	c.JSON(http.StatusOK, legoSetResponse)
}
