package legoset

import (
	"github.com/gin-gonic/gin"
	v1 "legocy-go/internal/delievery/http/resources"
	"legocy-go/internal/delievery/http/resources/lego"
	"net/http"
)

// SetCreate
// @Summary	Create Lego Set object
// @Tags		lego_sets_admin
// @ID			set_create
// @Param		data	body	lego.LegoSetRequest	true	"create data"
// @Produce	json
// @Success	200	{object}	map[string]interface{}
// @Failure	400	{object}	map[string]interface{}
// @Router		/admin/sets/ [post]
//
// @Security	JWT
func (lsh *LegoSetHandler) SetCreate(c *gin.Context) {
	var setRequest lego.LegoSetRequest
	if err := c.ShouldBindJSON(&setRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	legoSetValueObject := setRequest.ToLegoSeriesValueObject()
	err := lsh.service.LegoSetCreate(c.Request.Context(), legoSetValueObject)
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
