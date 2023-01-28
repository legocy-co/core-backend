package postgres

import (
	"golang.org/x/net/context"
	database "legocy-go/infrastructure/db"
	postgres "legocy-go/infrastructure/db/postgres"
	entities "legocy-go/infrastructure/db/postgres/entities"
	models "legocy-go/pkg/marketplace/models"
)

type LocationPostgresRepository struct {
	conn *postgres.PostgresConnection
}

func NewLocationPostgresRepository(conn *postgres.PostgresConnection) LocationPostgresRepository {
	return LocationPostgresRepository{conn: conn}
}

func (lpr *LocationPostgresRepository) GetLocations(c context.Context) ([]*models.Location, error) {
	var locations []*models.Location
	var locationsDB []*entities.LocationPostgres

	db := lpr.conn.GetDB()
	if db == nil {
		return locations, database.ErrConnectionLost
	}

	db.Find(&locationsDB)

	for _, entity := range locationsDB {
		locations = append(locations, entity.ToLocation())
	}

	var err error
	if len(locations) == 0 {
		err = database.ErrItemNotFound
	}

	return locations, err
}

func (lpr *LocationPostgresRepository) GetCountryLocations(c context.Context, country string) ([]*models.Location, error) {
	var locations []*models.Location
	var locationsDB []*entities.LocationPostgres

	db := lpr.conn.GetDB()
	if db == nil {
		return locations, database.ErrConnectionLost
	}

	db.Model(entities.LocationPostgres{}).Find(&locationsDB,
		entities.LocationPostgres{Country: country})

	for _, entity := range locationsDB {
		locations = append(locations, entity.ToLocation())
	}

	var err error
	if len(locations) == 0 {
		err = database.ErrItemNotFound
	}

	return locations, err
}

func (lpr *LocationPostgresRepository) CreateLocation(c context.Context, location *models.LocationBasic) error {
	db := lpr.conn.GetDB()
	if db == nil {
		return database.ErrConnectionLost
	}

	var entity *entities.LocationPostgres = entities.FromLocationBasic(location)
	if entity == nil {
		return database.ErrItemNotFound
	}

	result := db.Create(&entity)
	return result.Error
}
