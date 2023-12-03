package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/calculator"
	"net/http"
	"strconv"
)

// UpdateValuation
//
//		@Summary	Update LegoSetValuation (Admin)
//		@Tags		calculator_admin
//		@ID			update_lego_set_valuation_admin
//		@Param		data	body	calculator.LegoSetValuationUpdateRequest	true	"data"
//	 	@Param  	valuationID path int true "id"
//		@Produce	json
//		@Success	200	{object}	map[string]interface{}
//		@Failure	409	{object}	map[string]interface{}
//		@Failure	422	{object}	map[string]interface{}
//		@Router		/admin/sets-valuations/{valuationID} [put]
//
//		@Security	JWT
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
