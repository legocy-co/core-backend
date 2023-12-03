package admin

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	"github.com/legocy-co/legocy/internal/domain/calculator/models"
	repository "github.com/legocy-co/legocy/internal/domain/calculator/repository/admin"
)

type LegoSetValuationAdminService struct {
	r repository.LegoSetValuationAdminRepository
}

func NewService(r repository.LegoSetValuationAdminRepository) LegoSetValuationAdminService {
	return LegoSetValuationAdminService{r: r}
}

func (s LegoSetValuationAdminService) GetLegoSetValuations(
	c context.Context) ([]*models.LegoSetValuation, *errors.AppError) {

	return s.r.GetLegoSetValuations(c)
}

func (s LegoSetValuationAdminService) GetLegoSetValuationByID(
	c context.Context, id int) (*models.LegoSetValuation, *errors.AppError) {

	return s.r.GetLegoSetValuationByID(c, id)
}

func (s LegoSetValuationAdminService) AddLegoSetValuation(
	c context.Context, vo models.LegoSetValuationValueObject) *errors.AppError {

	return s.r.AddLegoSetValuation(c, vo)
}

func (s LegoSetValuationAdminService) DeleteLegoSetValuationByID(
	c context.Context, id int) *errors.AppError {

	return s.r.DeleteLegoSetValuationByID(c, id)
}

func (s LegoSetValuationAdminService) UpdateLegoSetValuationByID(
	c context.Context, id int, vo models.LegoSetValuationValueObject) *errors.AppError {

	return s.r.UpdateLegoSetValuationByID(c, id, vo)
}
