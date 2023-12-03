package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/calculator"
	"net/http"
)

// CreateValuation
//
//	@Summary	Create LegoSetValuation (Admin)
//	@Tags		calculator_admin
//	@ID			create_lego_set_valuation_admin
//	@Param		data	body	calculator.LegoSetValuationCreateRequest	true	"data"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	409	{object}	map[string]interface{}
//	@Failure	422	{object}	map[string]interface{}
//	@Router		/admin/sets-valuations/ [post]
//
//	@Security	JWT
func (h Handler) CreateValuation(ctx *gin.Context) {

	var createRequest *calculator.LegoSetValuationCreateRequest
	if e := ctx.ShouldBindJSON(&createRequest); e != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": e.Error()})
		return
	}

	vo, e := createRequest.ToLegoSetValuationVO()
	if e != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": e.Error()})
		return
	}

	err := h.service.AddLegoSetValuation(ctx, *vo)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
