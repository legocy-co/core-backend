package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/calculator"
	"net/http"
	"strconv"
)

// GetValuationByID
//
//	@Summary	Get LegoSetValuations (Admin)
//	@Tags		calculator_admin
//	@ID			detail_lego_set_valuation_admin
//	@Param valuationID path int true "id"
//	@Produce	json
//	@Success	200	{object}	calculator.LegoSetValuationResponse
//	@Failure	409	{object}	map[string]interface{}
//	@Failure	422	{object}	map[string]interface{}
//	@Router		/admin/sets-valuations/{valuationID} [get]
//
//	@Security	JWT
func (h Handler) GetValuationByID(ctx *gin.Context) {

	valuationId, e := strconv.Atoi(ctx.Param("valuationID"))
	if e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	valuation, err := h.service.GetLegoSetValuationByID(ctx, valuationId)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	response := calculator.FromLegoSetValuation(*valuation)
	ctx.JSON(http.StatusOK, response)
}
