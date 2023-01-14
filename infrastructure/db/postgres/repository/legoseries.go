package postgres

import (
	"context"
	d "legocy-go/infrastructure/db"
	p "legocy-go/infrastructure/db/postgres"
	entities "legocy-go/infrastructure/db/postgres/entities"
	models "legocy-go/pkg/lego/models"
)

type LegoSeriesPostgresRepository struct {
	conn *p.PostrgresConnection
}

func NewLegoSeriesPostgresRepository(conn *p.PostrgresConnection) LegoSeriesPostgresRepository {
	return LegoSeriesPostgresRepository{conn: conn}
}

func (psql *LegoSeriesPostgresRepository) CreateLegoSeries(c context.Context, s *models.LegoSeries) error {
	db := psql.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	entity := entities.FromLegoSeries(s)
	db.Create(entity)
	return nil
}

func (psql *LegoSeriesPostgresRepository) GetLegoSeriesList(c context.Context) ([]*models.LegoSeries, error) {
	var series []*models.LegoSeries
	db := psql.conn.GetDB()

	if db == nil {
		return series, d.ErrConnectionLost
	}

	var entitiesList []*entities.LegoSeriesPostgres
	db.Find(entitiesList)

	for _, entity := range entitiesList {
		series = append(series, entity.ToLegoSeries())
	}

	return series, nil
}

func (psql *LegoSeriesPostgresRepository) GetLegoSeries(c context.Context, id int) (*models.LegoSeries, error) {
	var entity *entities.LegoSeriesPostgres
	var series *models.LegoSeries

	db := psql.conn.GetDB()
	if db == nil {
		return series, d.ErrConnectionLost
	}

	db.First(&entity, id)
	series = entity.ToLegoSeries()
	return series, nil
}

func (psql *LegoSeriesPostgresRepository) DeleteLegoSeries(c context.Context, id int) error {
	db := psql.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	db.Delete(&entities.LegoSeriesPostgres{}, id)
	return nil
}
