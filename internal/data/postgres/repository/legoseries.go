package postgres

import (
	"context"
	d "github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/data/postgres"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	"github.com/legocy-co/legocy/internal/domain/lego"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/pkg/errors"
)

type LegoSeriesPostgresRepository struct {
	conn d.Storage
}

func (r LegoSeriesPostgresRepository) UpdateLegoSeries(legoSeriesID int, vo *models.LegoSeriesValueObject) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &postgres.ErrConnectionLost
	}

	var currentEntity *entities.LegoSeriesPostgres
	if db.Find(&currentEntity, legoSeriesID).RowsAffected <= 0 {
		e := errors.NewAppError(errors.NotFoundError, "LegoSeries not found")
		return &e
	}

	currentEntity = entities.GetUpdatedLegoSeriesPostgres(currentEntity, vo)

	tx := db.Begin()

	if err := tx.Save(currentEntity).Error; err != nil {
		tx.Rollback()
		e := errors.NewAppError(errors.ConflictError, err.Error())
		return &e
	}

	tx.Commit()
	return nil
}

func NewLegoSeriesPostgresRepository(conn d.Storage) LegoSeriesPostgresRepository {
	return LegoSeriesPostgresRepository{conn: conn}
}

func (r LegoSeriesPostgresRepository) CreateLegoSeries(c context.Context, s *models.LegoSeriesValueObject) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &postgres.ErrConnectionLost
	}

	entity := entities.FromLegoSeriesValueObject(s)
	_err := db.Create(&entity).Error

	if _err != nil {
		appErr := errors.NewAppError(errors.ConflictError, _err.Error())
		return &appErr
	}

	return nil
}

func (r LegoSeriesPostgresRepository) GetLegoSeriesList(c context.Context) ([]*models.LegoSeries, *errors.AppError) {

	db := r.conn.GetDB()

	if db == nil {
		return nil, &postgres.ErrConnectionLost
	}

	var entitiesList []*entities.LegoSeriesPostgres

	_err := db.Find(&entitiesList).Error
	if _err != nil {
		err := errors.NewAppError(errors.NotFoundError, _err.Error())
		return nil, &err
	}

	series := make([]*models.LegoSeries, 0, len(entitiesList))
	for _, entity := range entitiesList {
		series = append(series, entity.ToLegoSeries())
	}
	return series, nil
}

func (r LegoSeriesPostgresRepository) GetLegoSeries(
	c context.Context, id int) (*models.LegoSeries, *errors.AppError) {

	var entity *entities.LegoSeriesPostgres

	db := r.conn.GetDB()
	if db == nil {
		return nil, &postgres.ErrConnectionLost
	}

	if ok := db.First(&entity, id).RowsAffected > 0; !ok {
		return nil, &lego.ErrLegoSeriesNotFound
	}

	return entity.ToLegoSeries(), nil
}

func (r LegoSeriesPostgresRepository) GetLegoSeriesByName(
	c context.Context, name string) (*models.LegoSeries, *errors.AppError) {

	db := r.conn.GetDB()

	if db == nil {
		return nil, &postgres.ErrConnectionLost
	}

	var entity *entities.LegoSeriesPostgres
	_err := db.Where(entities.LegoSeriesPostgres{Name: name}).First(&entity).Error
	if _err != nil {
		appErr := errors.NewAppError(errors.NotFoundError, _err.Error())
		return nil, &appErr
	}

	return entity.ToLegoSeries(), nil
}

func (r LegoSeriesPostgresRepository) DeleteLegoSeries(c context.Context, id int) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &postgres.ErrConnectionLost
	}

	var err *errors.AppError

	_err := db.Delete(&entities.LegoSeriesPostgres{}, id).Error
	if _err != nil {
		*err = errors.NewAppError(errors.ConflictError, _err.Error())
	}

	return err
}
