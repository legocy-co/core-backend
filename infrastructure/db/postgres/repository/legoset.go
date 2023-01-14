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
