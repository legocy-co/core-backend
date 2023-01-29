package postgres

import (
	"context"
	d "legocy-go/internal/db"
	p "legocy-go/internal/db/postgres"
	entities "legocy-go/internal/db/postgres/entities"
	models "legocy-go/pkg/lego/models"
	"log"
)

type LegoSeriesPostgresRepository struct {
	conn *p.PostgresConnection
}

func NewLegoSeriesPostgresRepository(conn *p.PostgresConnection) LegoSeriesPostgresRepository {
	return LegoSeriesPostgresRepository{conn: conn}
}

func (psql *LegoSeriesPostgresRepository) CreateLegoSeries(c context.Context, s *models.LegoSeriesBasic) error {
	db := psql.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	entity := entities.FromLegoSeriesBasic(s)
	result := db.Create(&entity)
	return result.Error
}

func (psql *LegoSeriesPostgresRepository) GetLegoSeriesList(c context.Context) ([]*models.LegoSeries, error) {
	var series []*models.LegoSeries
	db := psql.conn.GetDB()

	if db == nil {
		return series, d.ErrConnectionLost
	}

	var entitiesList []*entities.LegoSeriesPostgres
	db.Find(&entitiesList)

	for _, entity := range entitiesList {
		series = append(series, entity.ToLegoSeries())
	}
	log.Print(series)
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

	result := db.Delete(&entities.LegoSeriesPostgres{}, id)
	return result.Error
}
