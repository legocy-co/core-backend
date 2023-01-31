package marketplace

import (
	"github.com/gin-gonic/gin"
	r "legocy-go/api/v1/resources"
	res "legocy-go/api/v1/resources/marketplace"
	s "legocy-go/api/v1/usecase/marketplace"
	"net/http"
)

type CurrencyHandler struct {
	service s.CurrencyUseCase
}

func NewCurrencyHandler(service s.CurrencyUseCase) CurrencyHandler {
	return CurrencyHandler{service: service}
}

func (h *CurrencyHandler) ListCurrencies(c *gin.Context) {
	currenciesList, err := h.service.CurrenciesList(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	var currenciesResponse []res.CurrencyResponse
	for _, currency := range currenciesList {
		currenciesResponse = append(currenciesResponse, res.GetCurrencyResponse(currency))
	}

	response := r.DataMetaResponse{
		Data: currenciesResponse,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
}
