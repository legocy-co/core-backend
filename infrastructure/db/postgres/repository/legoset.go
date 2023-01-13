package postgres

import (
	"context"
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
		return p.ErrConnectionLost
	}

	entity := entities.FromLegoSet(s)
	db.Create(entity)
	return nil
}
