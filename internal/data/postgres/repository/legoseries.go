package postgres

import (
	"context"
	d "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	models "legocy-go/internal/domain/lego/models"
)

type LegoSeriesPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewLegoSeriesPostgresRepository(conn d.DataBaseConnection) LegoSeriesPostgresRepository {
	return LegoSeriesPostgresRepository{conn: conn}
}

func (r LegoSeriesPostgresRepository) CreateLegoSeries(c context.Context, s *models.LegoSeriesValueObject) error {
	db := r.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	entity := entities.FromLegoSeriesValueObject(s)
	result := db.Create(&entity)
	return result.Error
}

func (r LegoSeriesPostgresRepository) GetLegoSeriesList(c context.Context) ([]*models.LegoSeries, error) {

	db := r.conn.GetDB()

	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var entitiesList []*entities.LegoSeriesPostgres
	db.Find(&entitiesList)

	series := make([]*models.LegoSeries, 0, len(entitiesList))
	for _, entity := range entitiesList {
		series = append(series, entity.ToLegoSeries())
	}
	return series, nil
}

func (r LegoSeriesPostgresRepository) GetLegoSeries(
	c context.Context, id int) (*models.LegoSeries, error) {

	var entity *entities.LegoSeriesPostgres
	var series *models.LegoSeries

	db := r.conn.GetDB()
	if db == nil {
		return series, d.ErrConnectionLost
	}

	db.First(&entity, id)
	series = entity.ToLegoSeries()
	return series, nil
}

func (r LegoSeriesPostgresRepository) GetLegoSeriesByName(
	c context.Context, name string) (*models.LegoSeries, error) {

	db := r.conn.GetDB()

	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var entity *entities.LegoSeriesPostgres
	err := db.Where(entities.LegoSeriesPostgres{Name: name}).First(&entity).Error
	if err != nil {
		return nil, err
	}

	return entity.ToLegoSeries(), nil
}

func (r LegoSeriesPostgresRepository) DeleteLegoSeries(c context.Context, id int) error {
	db := r.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	result := db.Delete(&entities.LegoSeriesPostgres{}, id)
	return result.Error
}
