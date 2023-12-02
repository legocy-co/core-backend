package admin

import (
	"context"
	"legocy-go/internal/app/errors"
	d "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	"legocy-go/internal/domain/calculator"
	"legocy-go/internal/domain/calculator/models"
	"legocy-go/internal/domain/collections"
)

type LegoSetValuationAdminPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewLegoSetValuationPostgresAdminRepository(conn d.DataBaseConnection) LegoSetValuationAdminPostgresRepository {
	return LegoSetValuationAdminPostgresRepository{conn: conn}
}

func (r LegoSetValuationAdminPostgresRepository) GetLegoSetValuations(
	c context.Context) ([]*models.LegoSetValuation, *errors.AppError) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var setValuations []entities.LegoSetValuationPostgres

	query := db.Model(
		&entities.LegoSetValuationPostgres{}).Preload("LegoSet").Find(&setValuations)
	if query.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, query.Error.Error())
		return nil, &appErr
	}

	if query.RowsAffected == 0 {
		return nil, &calculator.ErrLegoSetValuationNotFound
	}

	setValuationsDomain := make([]*models.LegoSetValuation, 0, len(setValuations))
	for _, entity := range setValuations {
		setValuationsDomain = append(setValuationsDomain, entity.ToLegoSetValuation())
	}

	return setValuationsDomain, nil
}

func (r LegoSetValuationAdminPostgresRepository) GetLegoSetValuationByID(
	c context.Context, id int) (*models.LegoSetValuation, *errors.AppError) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.LegoSetValuationPostgres
	res := db.Model(
		&entities.LegoSetValuationPostgres{}).Preload("LegoSet").First(&entity, id)
	if res.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return nil, &appErr
	}

	if entity == nil {
		return nil, &collections.ErrValuationNotFound
	}

	return entity.ToLegoSetValuation(), nil
}

func (r LegoSetValuationAdminPostgresRepository) GetLegoSetValuationBySetStateCurrency(
	c context.Context, setID int, setState string, currencyID int) (*models.LegoSetValuation, *errors.AppError) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.LegoSetValuationPostgres
	res := db.Model(
		&entities.LegoSetValuationPostgres{}).Preload("LegoSet").First(
		&entity, "lego_set_id = ?", setID, "state = ?", setState, "currency_id = ?", currencyID)
	if res.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return nil, &appErr
	}

	if entity == nil {
		return nil, &collections.ErrValuationNotFound
	}

	return entity.ToLegoSetValuation(), nil
}

func (r LegoSetValuationAdminPostgresRepository) AddLegoSetValuation(
	c context.Context, vo models.LegoSetValuationValueObject) *errors.AppError {

	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	entityToCreate := entities.FromLegoSetValuationVO(vo)

	tx := db.Begin()

	err := tx.Create(entityToCreate).Error
	if err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	tx.Commit()
	return nil
}

func (r LegoSetValuationAdminPostgresRepository) DeleteLegoSetValuationByID(c context.Context, id int) *errors.AppError {
	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	err := tx.Delete(&entities.LegoSetValuationPostgres{}, id).Error
	if err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	tx.Commit()
	return nil
}

func (r LegoSetValuationAdminPostgresRepository) UpdateLegoSetValuationByID(
	c context.Context, id int, vo models.LegoSetValuationValueObject) *errors.AppError {

	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	var entity *entities.LegoSetValuationPostgres
	err := db.First(&entity, id).Error
	if err != nil || entity == nil {
		return &calculator.ErrLegoSetValuationNotFound
	}

	entity = entities.GetUpdatedLegoSetValuationPostgres(vo, entity)

	err = tx.Save(entity).Error
	if err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	tx.Commit()
	return nil
}
