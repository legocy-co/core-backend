package postgres

import (
	"context"
	d "legocy-go/infrastructure/db"
	p "legocy-go/infrastructure/db/postgres"
	entities "legocy-go/infrastructure/db/postgres/entities"
	models "legocy-go/pkg/lego/models"
)

type LegoSetPostgresRepository struct {
	conn *p.PostrgresConnection
}

func (psql *LegoSetPostgresRepository) CreateLegoSet(c context.Context, s *models.LegoSet) error {
	db := psql.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	entity := entities.FromLegoSet(s)
	db.Create(entity)
	return nil
}

func (psql *LegoSetPostgresRepository) GetLegoSets(c *context.Context) ([]*models.LegoSet, error) {
	var legoSets []*models.LegoSet
	db := psql.conn.GetDB()

	if db == nil {
		return legoSets, d.ErrConnectionLost
	}

	var entitiesList []*entities.LegoSetPostgres
	db.Find(&entitiesList)

	for _, entity := range entitiesList {
		legoSets = append(legoSets, entity.ToLegoSet())
	}

	return legoSets, nil

}

func (psql *LegoSetPostgresRepository) GetLegoSetByID(c context.Context, id int) (*models.LegoSet, error) {
	var legoSet *models.LegoSet
	db := psql.conn.GetDB()

	if db == nil {
		return legoSet, d.ErrConnectionLost
	}

	var entity *entities.LegoSetPostgres
	db.First(&entity, id)

	legoSet = entity.ToLegoSet()
	return legoSet, nil
}

func (psql *LegoSetPostgresRepository) DeleteLegoSet(c *context.Context, id int) error {
	db := psql.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	db.Delete(&entities.LegoSetPostgres{}, id)
	return nil
}
