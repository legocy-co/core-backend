package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/api/v1/handlers/marketplace"
	m "legocy-go/api/v1/middleware"
	s "legocy-go/api/v1/usecase/marketplace"
)

func (r V1router) addCurrencies(rg *gin.RouterGroup, service s.CurrencyUseCase) {
	handler := marketplace.NewCurrencyHandler(service)

	currencies := rg.Group("/currencies").Use(m.Auth())
	{
		currencies.GET("/", handler.ListCurrencies)
		currencies.GET("/:currencySymbol", handler.CurrencyDetail)
	}

	currenciesAdmin := rg.Group("/currencies/admin").Use(m.AdminUserOnly())
	{
		currenciesAdmin.POST("/", handler.CreateCurrency)
	}

}
