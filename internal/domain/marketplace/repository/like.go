package marketplace

import (
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

type LikeRepository interface {
	AddLike(vo models.LikeValueObject) *errors.AppError
	RemoveLike(vo models.LikeValueObject) *errors.AppError
	GetLikesByUserID(userID int) ([]models.Like, *errors.AppError)
}
