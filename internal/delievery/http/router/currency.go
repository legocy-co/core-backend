package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/handlers/marketplace"
	m "legocy-go/internal/delievery/http/middleware"
	s "legocy-go/internal/delievery/http/usecase/marketplace"
)

func (r V1router) addCurrencies(rg *gin.RouterGroup, service s.CurrencyUseCase) {
	handler := marketplace.NewCurrencyHandler(service)

	currencies := rg.Group("/currencies").Use(m.Auth())
	{
		currencies.GET("/", handler.ListCurrencies)
		currencies.GET("/:currencySymbol", handler.CurrencyDetail)
	}

	currenciesAdmin := rg.Group("/admin/currencies").Use(m.AdminUserOnly())
	{
		currenciesAdmin.POST("/", handler.CreateCurrency)
	}

}
