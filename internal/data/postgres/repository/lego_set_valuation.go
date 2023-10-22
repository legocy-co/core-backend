package postgres

import (
	"context"
	d "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	"legocy-go/internal/domain/collections/models"
)

type LegoSetValuationPostgresRepository struct {
	conn d.DataBaseConnection
}

func (r LegoSetValuationPostgresRepository) GetLegoSetValuationsList(c context.Context, legoSetID int) ([]models.LegoSetValuation, error) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var setValuations []entities.LegoSetValuation

	res := db.Model(
		&entities.LegoSetValuation{}).Preload("LegoSet").Preload("Currency").Find(
		&setValuations, "lego_set_id = ?", legoSetID)
	if res.Error != nil {
		return nil, res.Error
	}

	setValuationsDomain := make([]models.LegoSetValuation, 0, len(setValuations))
	for _, entity := range setValuations {
		setValuationsDomain = append(setValuationsDomain, *entity.ToLegoSetValuation())
	}

	return setValuationsDomain, nil
}

func (r LegoSetValuationPostgresRepository) GetLegoSetValuationByID(c context.Context, id int) (*models.LegoSetValuation, error) {
	//TODO implement me
	panic("implement me")
}

func (r LegoSetValuationPostgresRepository) GetLegoSetValuationBySetStateCurrency(c context.Context, setID int, setState string, currencyID int) (*models.LegoSetValuation, error) {
	//TODO implement me
	panic("implement me")
}

func (r LegoSetValuationPostgresRepository) AddLegoSetValuation(c context.Context, vo models.LegoSetValuationValueObject) error {
	//TODO implement me
	panic("implement me")
}

func (r LegoSetValuationPostgresRepository) DeleteLegoSetValuationByID(c context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (r LegoSetValuationPostgresRepository) UpdateLegoSetValuationByID(c context.Context, id int, vo models.LegoSetValuationValueObject) error {
	//TODO implement me
	panic("implement me")
}
