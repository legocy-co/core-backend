package marketplace

import (
	"github.com/legocy-co/legocy/internal/app/errors"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
)

type LikeRepository interface {
	AddLike(vo models.LikeValueObject) *errors.AppError
	RemoveLike(vo models.LikeValueObject) *errors.AppError
	GetLikesByUserID(userID int) ([]models.Like, *errors.AppError)
}
