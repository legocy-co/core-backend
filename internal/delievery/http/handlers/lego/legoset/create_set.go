package legoset

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/errors"
	"legocy-go/internal/delievery/http/resources/lego"
	"net/http"
)

// SetCreate
//
//	@Summary	Create Lego Set object
//	@Tags		lego_sets_admin
//	@ID			set_create
//	@Param		data	body	lego.LegoSetRequest	true	"create data"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/sets/ [post]
//
//	@Security	JWT
func (lsh *LegoSetHandler) SetCreate(c *gin.Context) {
	var setRequest lego.LegoSetRequest
	if _err := c.ShouldBindJSON(&setRequest); _err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": _err.Error()})
		return
	}

	legoSetValueObject := setRequest.ToLegoSeriesValueObject()
	err := lsh.service.LegoSetCreate(c.Request.Context(), legoSetValueObject)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
