package postgres

import (
	"context"
	modelErr "legocy-go/pkg/lego/errors"
	models "legocy-go/pkg/lego/models"

	"github.com/jinzhu/gorm"
)

type LegoSeriesPostgresRepository struct {
	db *gorm.DB
}

// TODO:
func (psql *LegoSeriesPostgresRepository) CreateLegoSeries(c *context.Context, s *models.LegoSeries) error {
	if psql.db == nil {
		return modelErr.ErrSeriesNotFound
	}

	return modelErr.ErrSeriesAlreadyExists
}

func (psql *LegoSeriesPostgresRepository) GetLegoSeries(c *context.Context) ([]*models.LegoSeries, error) {
	var output []*models.LegoSeries
	return output, nil
}

func (psql *LegoSeriesPostgresRepository) DeleteLegoSeries(c *context.Context, id int) error {
	return modelErr.ErrSeriesNotFound
}
