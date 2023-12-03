package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"net/http"
	"strconv"
)

func (h Handler) DeleteValuation(ctx *gin.Context) {
	valuationID, e := strconv.Atoi(ctx.Param("valuationID"))
	if e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	err := h.service.DeleteLegoSetValuationByID(ctx, valuationID)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
