package admin

import (
	"context"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	"github.com/legocy-co/legocy/internal/data/postgres/utils"
	e "github.com/legocy-co/legocy/internal/domain/marketplace/errors"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
	"github.com/legocy-co/legocy/pkg/kafka"
	"github.com/legocy-co/legocy/pkg/pagination"
)

type MarketItemAdminPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewMarketItemAdminPostgresRepository(conn d.DataBaseConnection) MarketItemAdminPostgresRepository {
	return MarketItemAdminPostgresRepository{conn: conn}
}

func (m MarketItemAdminPostgresRepository) GetMarketItems(
	ctx pagination.PaginationContext,
) (pagination.Page[*models.MarketItemAdmin], *errors.AppError) {

	db := m.conn.GetDB()

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

func (m MarketItemAdminPostgresRepository) GetMarketItemByID(c context.Context, id int) (*models.MarketItemAdmin, *errors.AppError) {
	db := m.conn.GetDB()
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

func (m MarketItemAdminPostgresRepository) CreateMarketItem(c context.Context, vo *models.MarketItemAdminValueObject) *errors.AppError {
	db := m.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	entity := entities.FromMarketItemAdminValueObject(*vo)
	if entity == nil {
		return &d.ErrItemNotFound
	}

	result := tx.Create(&entity)

	tx.Commit()

	err := kafka.ProduceJSONEvent(
		kafka.MarketItemUpdatesTopic,
		map[string]interface{}{
			"itemID": int(entity.ID),
		},
	)
	if err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.InternalError, err.Error())
		return &appErr
	}

	if result.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return &appErr
	}

	return nil
}

func (m MarketItemAdminPostgresRepository) UpdateMarketItemByID(
	c context.Context, itemId int, vo *models.MarketItemAdminValueObject) (*models.MarketItemAdmin, *errors.AppError) {
	db := m.conn.GetDB()

	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	_ = db.First(&entity, itemId)
	if entity == nil {
		return nil, &e.ErrMarketItemsNotFound
	}

	entityUpdated := entity.GetUpdatedMarketItemAdmin(*vo)
	db.Save(entityUpdated)

	return m.GetMarketItemByID(c, itemId)
}

func (m MarketItemAdminPostgresRepository) DeleteMarketItemByID(c context.Context, itemId int) *errors.AppError {
	db := m.conn.GetDB()

	if db == nil {
		return &d.ErrConnectionLost
	}

	result := db.Delete(entities.MarketItemPostgres{}, itemId)
	if result.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return &appErr
	}

	return nil
}
