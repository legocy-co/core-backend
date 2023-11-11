package postgres

import (
	"context"
	d "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	"legocy-go/internal/domain/calculator"
	"legocy-go/internal/domain/calculator/models"
	"legocy-go/internal/domain/errors"
)

type LegoSetValuationPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewLegoSetValuationPostgresRepository(conn d.DataBaseConnection) LegoSetValuationPostgresRepository {
	return LegoSetValuationPostgresRepository{conn: conn}
}

func (r LegoSetValuationPostgresRepository) GetLegoSetValuationsList(c context.Context, legoSetID int) ([]models.LegoSetValuation, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var setValuations []entities.LegoSetValuation

	query := db.Model(
		&entities.LegoSetValuation{}).Preload("LegoSet").Preload("Currency").Find(
		&setValuations, "lego_set_id = ?", legoSetID)
	if query.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, query.Error.Error())
		return nil, &appErr
	}

	setValuationsDomain := make([]models.LegoSetValuation, 0, len(setValuations))
	for _, entity := range setValuations {
		setValuationsDomain = append(setValuationsDomain, *entity.ToLegoSetValuation())
	}

	return setValuationsDomain, nil
}

func (r LegoSetValuationPostgresRepository) GetLegoSetValuationByID(c context.Context, id int) (*models.LegoSetValuation, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.LegoSetValuation
	query := db.Model(
		&entities.LegoSetValuation{}).Preload("LegoSet").Preload("Currency").First(&entity, id)
	if query.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, query.Error.Error())
		return nil, &appErr
	}

	if entity == nil {
		return nil, &calculator.ErrLegoSetValuationNotFound
	}
	return entity.ToLegoSetValuation(), nil
}

func (r LegoSetValuationPostgresRepository) GetLegoSetValuationBySetStateCurrency(c context.Context, setID int, setState string, currencyID int) (*models.LegoSetValuation, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.LegoSetValuation
	query := db.Model(
		&entities.LegoSetValuation{}).Preload("LegoSet").Preload("Currency").First(
		&entity, "lego_set_id = ?", setID, "state = ?", setState, "currency_id = ?", currencyID)
	if query.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, query.Error.Error())
		return nil, &appErr
	}

	if entity == nil {
		return nil, &calculator.ErrLegoSetValuationNotFound
	}

	return entity.ToLegoSetValuation(), nil
}
