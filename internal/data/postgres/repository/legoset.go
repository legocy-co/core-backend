package postgres

import (
	"context"
	d "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	models "legocy-go/internal/domain/lego/models"
	"legocy-go/pkg/filter"
)

type LegoSetPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewLegoSetPostgresRepository(conn d.DataBaseConnection) LegoSetPostgresRepository {
	return LegoSetPostgresRepository{conn: conn}
}

func (psql LegoSetPostgresRepository) CreateLegoSet(c context.Context, s *models.LegoSetValueObject) error {
	db := psql.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	entity := entities.FromLegoSetValueObject(s)
	db.Create(entity)
	return nil
}

func (psql LegoSetPostgresRepository) GetLegoSets(c context.Context) ([]*models.LegoSet, error) {
	db := psql.conn.GetDB()

	if db == nil {
		return nil, d.ErrConnectionLost
	}
	pagination := c.Value("pagination").(*filter.QueryParams)

	var entitiesList []*entities.LegoSetPostgres
	db.Model(entities.LegoSetPostgres{}).
		Scopes(filter.FilterDbByQueryParams(pagination, filter.PAGINATE)).
		Preload("LegoSeries").Find(&entitiesList)

	legoSets := make([]*models.LegoSet, 0, len(entitiesList))
	for _, entity := range entitiesList {
		legoSets = append(legoSets, entity.ToLegoSet())
	}

	return legoSets, nil

}

func (psql LegoSetPostgresRepository) GetLegoSetByID(c context.Context, id int) (*models.LegoSet, error) {
	var legoSet *models.LegoSet
	db := psql.conn.GetDB()

	if db == nil {
		return legoSet, d.ErrConnectionLost
	}

	var entity *entities.LegoSetPostgres
	db.Preload("LegoSeries").First(&entity, id)

	legoSet = entity.ToLegoSet()
	return legoSet, nil
}

func (psql LegoSetPostgresRepository) DeleteLegoSet(c context.Context, id int) error {
	db := psql.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	db.Delete(&entities.LegoSetPostgres{}, id)
	return nil
}
