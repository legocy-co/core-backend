package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"net/http"
	"strconv"
)

// DeleteValuation
//
//	@Summary	Delete LegoSetValuations (Admin)
//	@Tags		calculator_admin
//	@ID			delete_lego_set_valuation_admin
//	@Param valuationID path int true "id"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	409	{object}	map[string]interface{}
//	@Failure	422	{object}	map[string]interface{}
//	@Router		/admin/sets-valuations/{valuationID} [delete]
//
//	@Security	JWT
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
