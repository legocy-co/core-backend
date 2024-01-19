package repository

import (
	"github.com/legocy-co/legocy/internal/app/errors"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSetImageRepository interface {
	Get(legoSetID int) ([]*models.LegoSetImage, *errors.AppError)
	Store(vo models.LegoSetImageValueObject) (*models.LegoSetImage, *errors.AppError)
	Delete(id int) *errors.AppError
}
