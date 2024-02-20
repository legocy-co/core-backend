package repository

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/pkg/pagination"
)

type LegoSetRepository interface {
	CreateLegoSet(c context.Context, s *models.LegoSetValueObject) *errors.AppError
	GetLegoSets(c context.Context) ([]*models.LegoSet, *errors.AppError)
	GetSetsPage(ctx pagination.PaginationContext) (pagination.Page[models.LegoSet], *errors.AppError)
	GetLegoSetByID(c context.Context, id int) (*models.LegoSet, *errors.AppError)
	DeleteLegoSet(c context.Context, id int) *errors.AppError
	UpdateLegoSetByID(legoSetID int, vo *models.LegoSetValueObject) *errors.AppError
}
