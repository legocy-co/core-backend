package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/calculator"
	"net/http"
	"strconv"
)

func (h Handler) UpdateValuation(ctx *gin.Context) {

	valuationID, e := strconv.Atoi(ctx.Param("valuationID"))
	if e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	var updateRequest *calculator.LegoSetValuationUpdateRequest
	if e := ctx.ShouldBindJSON(&updateRequest); e != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": e.Error()})
		return
	}

	vo, e := updateRequest.ToLegoSetValuationVO()
	if e != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": e.Error()})
		return
	}

	err := h.service.UpdateLegoSetValuationByID(ctx, valuationID, *vo)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
