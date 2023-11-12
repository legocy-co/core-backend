package postgres

import (
	"golang.org/x/net/context"
	"legocy-go/internal/app/errors"
	database "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	models "legocy-go/internal/domain/marketplace/models"
)

type LocationPostgresRepository struct {
	conn database.DataBaseConnection
}

func NewLocationPostgresRepository(conn database.DataBaseConnection) LocationPostgresRepository {
	return LocationPostgresRepository{conn: conn}
}

func (lpr LocationPostgresRepository) GetLocations(c context.Context) ([]*models.Location, *errors.AppError) {
	var locationsDB []*entities.LocationPostgres

	db := lpr.conn.GetDB()
	if db == nil {
		return nil, &database.ErrConnectionLost
	}

	db.Find(&locationsDB)

	locations := make([]*models.Location, 0, len(locationsDB))
	for _, entity := range locationsDB {
		locations = append(locations, entity.ToLocation())
	}

	var err *errors.AppError
	if len(locations) == 0 {
		err = &database.ErrItemNotFound
	}

	return locations, err
}

func (lpr LocationPostgresRepository) GetCountryLocations(c context.Context, country string) ([]*models.Location, *errors.AppError) {
	var locationsDB []*entities.LocationPostgres

	db := lpr.conn.GetDB()
	if db == nil {
		return nil, &database.ErrConnectionLost
	}

	db.Model(entities.LocationPostgres{}).Find(&locationsDB,
		entities.LocationPostgres{Country: country})

	locations := make([]*models.Location, 0, len(locationsDB))
	for _, entity := range locationsDB {
		locations = append(locations, entity.ToLocation())
	}

	var err *errors.AppError
	if len(locations) == 0 {
		err = &database.ErrItemNotFound
	}

	return locations, err
}

func (lpr LocationPostgresRepository) CreateLocation(c context.Context, location *models.LocationValueObject) *errors.AppError {
	db := lpr.conn.GetDB()
	if db == nil {
		return &database.ErrConnectionLost
	}

	var entity *entities.LocationPostgres = entities.FromLocationValueObject(location)
	if entity == nil {
		return &database.ErrItemNotFound
	}

	var err *errors.AppError

	_err := db.Create(&entity).Error
	if _err != nil {
		*err = errors.NewAppError(errors.ConflictError, _err.Error())
	}

	return err
}

func (lpr LocationPostgresRepository) DeleteLocation(c context.Context, id int) *errors.AppError {
	db := lpr.conn.GetDB()
	if db == nil {
		return &database.ErrConnectionLost
	}

	var err *errors.AppError

	_err := db.Delete(entities.LocationPostgres{}, id).Error
	if _err != nil {
		*err = errors.NewAppError(errors.ConflictError, _err.Error())
	}

	return err
}
