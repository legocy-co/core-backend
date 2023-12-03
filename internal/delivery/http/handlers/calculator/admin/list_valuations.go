package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/errors"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/calculator"
	"net/http"
)

func (h Handler) GetLegoSetValuations(ctx *gin.Context) {

	valuations, err := h.service.GetLegoSetValuations(ctx)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		ctx.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	response := make([]calculator.LegoSetValuationResponse, 0, len(valuations))
	for _, valuation := range valuations {
		response = append(response, calculator.FromLegoSetValuation(*valuation))
	}

	ctx.JSON(http.StatusOK, response)
}
