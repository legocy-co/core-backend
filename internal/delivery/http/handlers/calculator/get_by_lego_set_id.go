package calculator

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	schemas "github.com/legocy-co/legocy/internal/delivery/http/schemas/calculator"
	"net/http"
	"strconv"
)

// GetValuationsByLegoSetID godoc
// @Summary Get valuations by lego set id
// @Description Get valuations by lego set id
// @Tags calculator
// @Accept json
// @Produce json
// @Param legoSetID path int true "Lego Set ID"
// @Success 200 {object} []schemas.LegoSetValuationResponse
// @Failure 400 {object} map[string]interface{}
// @Router /sets-valuations/{legoSetID} [get]
// @Security JWT
func (h LegoSetValuationHandler) GetValuationsByLegoSetID(ctx *gin.Context) {
	legoSetId, e := strconv.Atoi(ctx.Param("legoSetID"))
	if e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid lego set id"})
	}

	valuationsDomain, err := h.service.GetLegoSetValuations(ctx, legoSetId)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	valuationsResponse := make([]schemas.LegoSetValuationResponse, 0, len(valuationsDomain))
	for _, v := range valuationsDomain {
		valuationsResponse = append(valuationsResponse, schemas.FromLegoSetValuation(v))
	}

	ctx.JSON(http.StatusOK, valuationsResponse)
}
