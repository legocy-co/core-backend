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

func (h CurrencyHandler) ListCurrencies(c *gin.Context) {
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

func (h CurrencyHandler) CurrencyDetail(c *gin.Context) {
	currencySymbol := c.Param("currencySymbol")
	if currencySymbol == "" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Could not extact currency symbol"})
		return
	}

	curr, err := h.service.CurrencyDetail(c, currencySymbol)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	currResponse := res.GetCurrencyResponse(curr)
	response := r.DataMetaResponse{
		Data: currResponse,
		Meta: r.SuccessMetaResponse,
	}

	r.Respond(c.Writer, response)
}

func (h CurrencyHandler) CreateCurrency(c *gin.Context) {
	var currencyReq *res.CurrencyRequest
	if err := c.ShouldBindJSON(&currencyReq); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currency := currencyReq.ToCurrencyBasic()
	err := h.service.CreateCurrency(c, currency)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	r.Respond(c.Writer, r.DataMetaResponse{
		Data: currencyReq,
		Meta: r.SuccessMetaResponse,
	})
}
