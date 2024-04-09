package repository

import (
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

type LegoSetImageRepository interface {
	Get(legoSetID int) ([]*models.LegoSetImage, *errors.AppError)
	Store(vo models.LegoSetImageValueObject) (*models.LegoSetImage, *errors.AppError)
	Delete(id int) *errors.AppError
}
