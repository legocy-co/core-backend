package postgres

import (
	"context"
	d "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	"legocy-go/internal/domain/errors"
	models "legocy-go/internal/domain/lego/models"
)

type LegoSeriesPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewLegoSeriesPostgresRepository(conn d.DataBaseConnection) LegoSeriesPostgresRepository {
	return LegoSeriesPostgresRepository{conn: conn}
}

func (r LegoSeriesPostgresRepository) CreateLegoSeries(c context.Context, s *models.LegoSeriesValueObject) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &d.ErrConnectionLost
	}

	var err *errors.AppError

	entity := entities.FromLegoSeriesValueObject(s)
	_err := db.Create(&entity).Error

	if _err != nil {
		*err = errors.NewAppError(errors.ConflictError, _err.Error())
	}

	return err
}

func (r LegoSeriesPostgresRepository) GetLegoSeriesList(c context.Context) ([]*models.LegoSeries, *errors.AppError) {

	db := r.conn.GetDB()

	if db == nil {
		return nil, &d.ErrConnectionLost
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
	var series *models.LegoSeries
	var err *errors.AppError

	db := r.conn.GetDB()
	if db == nil {
		return series, &d.ErrConnectionLost
	}

	_err := db.First(&entity, id).Error
	if _err != nil {
		*err = errors.NewAppError(errors.NotFoundError, _err.Error())
		return series, err
	}

	series = entity.ToLegoSeries()
	return series, nil
}

func (r LegoSeriesPostgresRepository) GetLegoSeriesByName(
	c context.Context, name string) (*models.LegoSeries, *errors.AppError) {

	db := r.conn.GetDB()

	var err *errors.AppError

	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.LegoSeriesPostgres
	_err := db.Where(entities.LegoSeriesPostgres{Name: name}).First(&entity).Error
	if _err != nil {
		*err = errors.NewAppError(errors.NotFoundError, _err.Error())
		return nil, err
	}

	return entity.ToLegoSeries(), nil
}

func (r LegoSeriesPostgresRepository) DeleteLegoSeries(c context.Context, id int) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &d.ErrConnectionLost
	}

	var err *errors.AppError

	_err := db.Delete(&entities.LegoSeriesPostgres{}, id).Error
	if _err != nil {
		*err = errors.NewAppError(errors.ConflictError, _err.Error())
	}

	return err
}
