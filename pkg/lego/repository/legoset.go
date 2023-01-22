package repository

import (
	"context"
	models "legocy-go/pkg/lego/models"
)

type LegoSetRepository interface {
	CreateLegoSet(c context.Context, s *models.LegoSet) error
	GetLegoSets(c context.Context) ([]*models.LegoSet, error)
	GetLegoSetByID(c context.Context, id int) (*models.LegoSet, error)
	DeleteLegoSet(c context.Context, id int) error
}
