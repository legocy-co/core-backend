package admin

import (
	"context"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	"github.com/legocy-co/legocy/internal/data/postgres/utils"
	e "github.com/legocy-co/legocy/internal/domain/marketplace/errors"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
	"github.com/legocy-co/legocy/internal/pkg/events"
	"github.com/legocy-co/legocy/pkg/pagination"
)

type MarketItemAdminPostgresRepository struct {
	conn       d.DBConn
	dispatcher events.Dispatcher
}

func NewMarketItemAdminPostgresRepository(conn d.DBConn, dispatcher events.Dispatcher) MarketItemAdminPostgresRepository {
	return MarketItemAdminPostgresRepository{
		conn:       conn,
		dispatcher: dispatcher,
	}
}

func (r MarketItemAdminPostgresRepository) GetMarketItems(ctx pagination.PaginationContext) (pagination.Page[*models.MarketItemAdmin], *errors.AppError) {

	db := r.conn.GetDB()

	if db == nil {
		return pagination.NewEmptyPage[*models.MarketItemAdmin](), &d.ErrConnectionLost
	}

	var itemsDB []*entities.MarketItemPostgres

	query := db.Model(
		&entities.MarketItemPostgres{},
	).
		Preload("Seller").
		Preload("Images").
		Preload("LegoSet").
		Preload("LegoSet.LegoSeries")

	query = utils.AddPaginationQuery(query, ctx)

	err := query.Find(&itemsDB).Error
	if err != nil {
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return pagination.NewEmptyPage[*models.MarketItemAdmin](), &appErr
	}

	marketItemsAdmin := make([]*models.MarketItemAdmin, 0, len(itemsDB))
	for _, entity := range itemsDB {
		marketItemsAdmin = append(marketItemsAdmin, entity.ToMarketItemAdmin())
	}

	var total int64
	db.Model(&entities.MarketItemPostgres{}).Count(&total)

	return pagination.NewPage[*models.MarketItemAdmin](
		marketItemsAdmin,
		int(total),
		ctx.GetLimit(),
		ctx.GetOffset(),
	), nil

}

func (r MarketItemAdminPostgresRepository) GetMarketItemByID(c context.Context, id int) (*models.MarketItemAdmin, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	query := db.Preload(
		"Seller").Preload("Images").Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Find(&entity, "id = ?", id)

	if query.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, query.Error.Error())
		return nil, &appErr
	}

	// Not Found
	if query.RowsAffected == 0 {
		return nil, &e.ErrMarketItemsNotFound
	}

	return entity.ToMarketItemAdmin(), nil
}

func (r MarketItemAdminPostgresRepository) CreateMarketItem(c context.Context, vo *models.MarketItemAdminValueObject) *errors.AppError {
	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	entity := entities.FromMarketItemAdminValueObject(*vo)
	if entity == nil {
		return &d.ErrItemNotFound
	}

	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(
			errors.ConflictError,
			err.Error(),
		)
		return &appErr
	}

	return nil
}

func (r MarketItemAdminPostgresRepository) UpdateMarketItemByID(
	c context.Context, itemId int, vo *models.MarketItemAdminValueObject) (*models.MarketItemAdmin, *errors.AppError) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	if ok := db.First(&entity, itemId).RowsAffected > 0; !ok {
		return nil, &e.ErrMarketItemsNotFound
	}

	tx := db.Begin()

	entityUpdated := entity.GetUpdatedMarketItemAdmin(*vo)
	if err := tx.Save(entityUpdated).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(
			errors.ConflictError,
			err.Error(),
		)
		return nil, &appErr
	}

	tx.Commit()

	return r.GetMarketItemByID(c, itemId)
}

func (r MarketItemAdminPostgresRepository) DeleteMarketItemByID(c context.Context, itemId int) *errors.AppError {

	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	if err := tx.Delete(entities.MarketItemPostgres{}, itemId).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(
			errors.ConflictError,
			err.Error(),
		)
		return &appErr
	}

	tx.Commit()

	return nil
}
