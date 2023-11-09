package postgres

import (
	"golang.org/x/net/context"
	d "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	"legocy-go/internal/domain/errors"
	models "legocy-go/internal/domain/marketplace/models"
)

type CurrencyPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewCurrencyPostgresRepository(conn d.DataBaseConnection) *CurrencyPostgresRepository {
	return &CurrencyPostgresRepository{conn: conn}
}

func (cpr *CurrencyPostgresRepository) GetCurrencies(c context.Context) ([]*models.Currency, *errors.AppError) {
	var currenciesDB []*entities.CurrencyPostgres
	db := cpr.conn.GetDB()

	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	db.Find(&currenciesDB)

	currencies := make([]*models.Currency, 0, len(currenciesDB))
	for _, entity := range currenciesDB {
		currencies = append(currencies, entity.ToCurrency())
	}

	return currencies, nil
}

func (cpr *CurrencyPostgresRepository) GetCurrency(c context.Context, symbol string) (*models.Currency, *errors.AppError) {
	var currency *models.Currency
	var currencyDB *entities.CurrencyPostgres
	db := cpr.conn.GetDB()

	if db == nil {
		return currency, &d.ErrConnectionLost
	}

	db.Model(entities.CurrencyPostgres{}).First(&currencyDB, "symbol = ?", symbol)

	if currencyDB == nil {
		return currency, &d.ErrItemNotFound
	}

	currency = currencyDB.ToCurrency()
	return currency, nil
}

func (cpr *CurrencyPostgresRepository) CreateCurrency(c context.Context, currency *models.CurrencyValueObject) *errors.AppError {
	db := cpr.conn.GetDB()

	if db == nil {
		return &d.ErrConnectionLost
	}

	var entity *entities.CurrencyPostgres = entities.FromCurrencyValueObject(currency)
	result := db.Create(&entity)

	var err *errors.AppError
	if result.Error != nil {
		*err = errors.NewAppError(errors.ConflictError, result.Error.Error())
	}

	return err
}
