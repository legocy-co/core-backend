package marketplace

import models "legocy-go/pkg/marketplace/models"

type CurrencyRequest struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (cr *CurrencyRequest) ToCurrencyBasic() *models.CurrencyBasic {
	return &models.CurrencyBasic{
		Name:   cr.Name,
		Symbol: cr.Symbol,
	}
}

type CurrencyResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func GetCurrencyResponse(c *models.Currency) CurrencyResponse {
	return CurrencyResponse{
		ID:     c.ID,
		Name:   c.Name,
		Symbol: c.Symbol,
	}
}
