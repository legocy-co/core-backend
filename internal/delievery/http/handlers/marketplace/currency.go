package marketplace

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/errors"
	"legocy-go/internal/delievery/http/resources"
	"legocy-go/internal/delievery/http/resources/marketplace"
	s "legocy-go/internal/domain/marketplace/service"
	"net/http"
)

type CurrencyHandler struct {
	service s.CurrencyUseCase
}

func NewCurrencyHandler(service s.CurrencyUseCase) CurrencyHandler {
	return CurrencyHandler{service: service}
}

// ListCurrencies
//
//	@Summary	Get all currencies
//	@Tags		currencies
//	@ID			list_currencies
//	@Produce	json
//	@Success	200	{object}	[]marketplace.CurrencyResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/currencies/ [get]
//
//	@Security	JWT
func (h CurrencyHandler) ListCurrencies(c *gin.Context) {
	currenciesList, err := h.service.CurrenciesList(c)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	var currenciesResponse []marketplace.CurrencyResponse
	for _, currency := range currenciesList {
		currenciesResponse = append(currenciesResponse, marketplace.GetCurrencyResponse(currency))
	}

	c.JSON(http.StatusOK, currenciesResponse)
}

// CurrencyDetail
//
//	@Summary	Get currency by symbol
//	@Tags		currencies
//	@ID			currency_detail
//	@Param		currencySymbol	path	string	true	"currency symbol"
//	@Produce	json
//	@Success	200	{object}	marketplace.CurrencyResponse
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/currencies/{currencySymbol} [get]
//
//	@Security	JWT
func (h CurrencyHandler) CurrencyDetail(c *gin.Context) {
	currencySymbol := c.Param("currencySymbol")
	if currencySymbol == "" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": "Could not extact currency symbol"})
		return
	}

	curr, err := h.service.CurrencyDetail(c, currencySymbol)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	currResponse := marketplace.GetCurrencyResponse(curr)
	c.JSON(http.StatusOK, currResponse)
}

// CreateCurrency
//
//	@Summary	Get currency by symbol
//	@Tags		currencies_admin
//	@ID			currency_create
//	@Param		data	body	marketplace.CurrencyRequest	true	"currency symbol"
//	@Produce	json
//	@Success	200	{object}	map[string]interface{}
//	@Failure	400	{object}	map[string]interface{}
//	@Router		/admin/currencies/ [post]
//
//	@Security	JWT
func (h CurrencyHandler) CreateCurrency(c *gin.Context) {
	var currencyReq *marketplace.CurrencyRequest
	if err := c.ShouldBindJSON(&currencyReq); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currency := currencyReq.ToCurrencyValueObject()
	err := h.service.CreateCurrency(c, currency)
	if err != nil {
		httpErr := errors.FromAppError(*err)
		c.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
		return
	}

	v1.Respond(c.Writer, v1.DataMetaResponse{
		Data: currencyReq,
		Meta: v1.SuccessMetaResponse,
	})
}
