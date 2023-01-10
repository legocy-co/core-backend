package repository

import (
	"context"
	"legocy-go/pkg/lego/models"
)

type LegoSetRepository interface {
	CreateLegoSet(c *context.Context, s *models.LegoSet) error
	GetLegoSets(c *context.Context) ([]*models.LegoSet, error)
	DeleteLegoSet(c *context.Context, id int) error
}
