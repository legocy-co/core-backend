package postgres

import (
	"context"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entities"
	"legocy-go/pkg/filter"
	models "legocy-go/pkg/lego/models"
)

type LegoSetPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewLegoSetPostgresRepository(conn d.DataBaseConnection) LegoSetPostgresRepository {
	return LegoSetPostgresRepository{conn: conn}
}

func (psql LegoSetPostgresRepository) CreateLegoSet(c context.Context, s *models.LegoSetBasic) error {
	db := psql.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	entity := entities.FromLegoSetBasic(s)
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
