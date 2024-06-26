package postgres

import (
	"context"
	d "github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/data/postgres"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	"github.com/legocy-co/legocy/internal/data/postgres/utils"
	pgFilter "github.com/legocy-co/legocy/internal/data/postgres/utils/filters"
	"github.com/legocy-co/legocy/internal/domain/lego"
	"github.com/legocy-co/legocy/internal/domain/lego/filters"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/pkg/errors"
	"github.com/legocy-co/legocy/lib/pagination"
)

type LegoSetPostgresRepository struct {
	conn d.Storage
}

func NewLegoSetPostgresRepository(conn d.Storage) LegoSetPostgresRepository {
	return LegoSetPostgresRepository{conn: conn}
}

func (r LegoSetPostgresRepository) CreateLegoSet(c context.Context, s *models.LegoSetValueObject) (*models.LegoSet, *errors.AppError) {
	db := r.conn.GetDB()

	if db == nil {
		return nil, &postgres.ErrConnectionLost
	}

	entity := entities.FromLegoSetValueObject(s)
	result := db.Create(entity)

	if result.Error != nil {
		_e := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return nil, &_e
	}

	return r.GetLegoSetByID(c, int(entity.ID))
}

func (r LegoSetPostgresRepository) GetLegoSets(c context.Context) ([]*models.LegoSet, *errors.AppError) {
	db := r.conn.GetDB()

	var err *errors.AppError

	if db == nil {
		return nil, &postgres.ErrConnectionLost
	}

	var entitiesList []*entities.LegoSetPostgres
	_err := db.Model(
		entities.LegoSetPostgres{},
	).Preload("LegoSeries").Find(&entitiesList).Error

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
		return legoSet, &postgres.ErrConnectionLost
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

	if db == nil {
		return &postgres.ErrConnectionLost
	}

	_err := db.Delete(&entities.LegoSetPostgres{}, id).Error
	if _err != nil {
		_e := errors.NewAppError(errors.ConflictError, _err.Error())
		return &_e
	}

	return nil
}

func (r LegoSetPostgresRepository) GetSetsPage(ctx pagination.PaginationContext, filter *filters.LegoSetFilterCriteria) (pagination.Page[models.LegoSet], *errors.AppError) {
	var err *errors.AppError
	db := r.conn.GetDB()

	if db == nil {
		return pagination.Page[models.LegoSet]{}, &postgres.ErrConnectionLost
	}

	var total int64
	totalQuery := pgFilter.AddLegoSetFilters(
		db.Model(&entities.LegoSetPostgres{}), filter, false, "",
	)
	totalQuery.Count(&total)

	var entitiesList []*entities.LegoSetPostgres

	query := db.Model(
		entities.LegoSetPostgres{},
	).Preload("LegoSeries").Preload("Images").Order("id desc")

	query = pgFilter.AddLegoSetFilters(query, filter, false, "")
	query = utils.AddPaginationQuery(query, ctx)

	_err := query.Find(&entitiesList).Error
	if _err != nil {
		appErr := errors.NewAppError(errors.InternalError, _err.Error())
		return pagination.NewEmptyPage[models.LegoSet](), &appErr
	}

	legoSets := make([]models.LegoSet, 0, len(entitiesList))
	for _, entity := range entitiesList {
		legoSets = append(legoSets, *entity.ToLegoSet())
	}

	page := pagination.NewPage[models.LegoSet](
		legoSets, int(total), ctx.GetLimit(), ctx.GetOffset())

	return page, err
}

func (r LegoSetPostgresRepository) UpdateLegoSetByID(legoSetID int, vo *models.LegoSetValueObject) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &postgres.ErrConnectionLost
	}

	var currentEntity *entities.LegoSetPostgres
	if db.Find(&currentEntity, legoSetID).RowsAffected <= 0 {
		e := errors.NewAppError(errors.NotFoundError, "LegoSet not found")
		return &e
	}

	currentEntity = entities.GetUpdatedLegoSetPostgres(currentEntity, vo)

	tx := db.Begin()

	if err := tx.Save(currentEntity).Error; err != nil {
		tx.Rollback()
		e := errors.NewAppError(errors.ConflictError, err.Error())
		return &e
	}

	tx.Commit()
	return nil

}
