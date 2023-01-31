package postgres

import (
	"golang.org/x/net/context"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entities"
	models "legocy-go/pkg/marketplace/models"
)

type CurrencyPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewCurrencyPostgresRepository(conn d.DataBaseConnection) *CurrencyPostgresRepository {
	return &CurrencyPostgresRepository{conn: conn}
}

func (cpr *CurrencyPostgresRepository) GetCurrencies(c context.Context) ([]*models.Currency, error) {
	var currencies []*models.Currency
	var currenciesDB []*entities.CurrencyPostgres
	db := cpr.conn.GetDB()

	if db == nil {
		return currencies, d.ErrConnectionLost
	}

	db.Find(&currenciesDB)

	for _, entity := range currenciesDB {
		currencies = append(currencies, entity.ToCurrency())
	}

	return currencies, nil
}

func (cpr *CurrencyPostgresRepository) GetCurrency(c context.Context, symbol string) (*models.Currency, error) {
	var currency *models.Currency
	var currencyDB *entities.CurrencyPostgres
	db := cpr.conn.GetDB()

	if db == nil {
		return currency, d.ErrConnectionLost
	}

	db.Model(entities.CurrencyPostgres{}).First(&currencyDB,
		entities.CurrencyPostgres{Symbol: symbol})

	if currencyDB == nil {
		return currency, d.ErrItemNotFound
	}

	currency = currencyDB.ToCurrency()
	return currency, nil
}

func (cpr *CurrencyPostgresRepository) CreateCurrency(c context.Context, currency *models.CurrencyBasic) error {
	db := cpr.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	var entity *entities.CurrencyPostgres = entities.FromCurrencyBasic(currency)
	result := db.Create(&entity)
	return result.Error
}
