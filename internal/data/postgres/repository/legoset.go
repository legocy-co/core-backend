package postgres

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	"github.com/legocy-co/legocy/internal/domain/lego"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSetPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewLegoSetPostgresRepository(conn d.DataBaseConnection) LegoSetPostgresRepository {
	return LegoSetPostgresRepository{conn: conn}
}

func (r LegoSetPostgresRepository) CreateLegoSet(c context.Context, s *models.LegoSetValueObject) *errors.AppError {
	db := r.conn.GetDB()

	var err *errors.AppError

	if db == nil {
		return &d.ErrConnectionLost
	}

	entity := entities.FromLegoSetValueObject(s)
	result := db.Create(entity)

	if result.Error != nil {
		*err = errors.NewAppError(errors.ConflictError, result.Error.Error())
	}

	return err
}

func (r LegoSetPostgresRepository) GetLegoSets(c context.Context) ([]*models.LegoSet, *errors.AppError) {
	db := r.conn.GetDB()

	var err *errors.AppError

	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entitiesList []*entities.LegoSetPostgres
	_err := db.Model(
		entities.LegoSetPostgres{},
	).Preload("LegoSeries").Preload("Images").
		Find(&entitiesList).Error

	if _err != nil {
		appErr := errors.NewAppError(errors.InternalError, _err.Error())
		err = &appErr
	}

	legoSets := make([]*models.LegoSet, 0, len(entitiesList))
	for _, entity := range entitiesList {
		legoSets = append(legoSets, entity.ToLegoSet())
	}

	return legoSets, err

}

func (r LegoSetPostgresRepository) GetLegoSetByID(c context.Context, id int) (*models.LegoSet, *errors.AppError) {
	var legoSet *models.LegoSet
	var err *errors.AppError
	db := r.conn.GetDB()

	if db == nil {
		return legoSet, &d.ErrConnectionLost
	}

	var entity *entities.LegoSetPostgres
	query := db.Preload("LegoSeries").Preload("Images").First(&entity, id)

	if query.Error != nil {
		*err = errors.NewAppError(errors.NotFoundError, query.Error.Error())
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, &lego.ErrLegoSetsNotFound
	}

	legoSet = entity.ToLegoSet()
	return legoSet, err
}

func (r LegoSetPostgresRepository) DeleteLegoSet(c context.Context, id int) *errors.AppError {
	db := r.conn.GetDB()

	var err *errors.AppError

	if db == nil {
		return &d.ErrConnectionLost
	}

	_err := db.Delete(&entities.LegoSetPostgres{}, id).Error
	if _err != nil {
		*err = errors.NewAppError(errors.ConflictError, _err.Error())
	}

	return nil
}
