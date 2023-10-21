package postgres

import (
	models "legocy-go/internal/domain/marketplace/models"
)

type LocationPostgres struct {
	Model
	Country string `gorm:"uniqueIndex:idx_country_city"`
	City    string `gorm:"uniqueIndex:idx_country_city"`
}

func (lp *LocationPostgres) ToLocation() *models.Location {
	return &models.Location{
		ID:      int(lp.ID),
		Country: lp.Country,
		City:    lp.City,
	}
}

func FromLocationValueObject(loc *models.LocationValueObject) *LocationPostgres {
	return &LocationPostgres{
		City:    loc.City,
		Country: loc.Country,
	}
}

func FromLocation(loc *models.Location) *LocationPostgres {
	return &LocationPostgres{
		Country: loc.Country,
		City:    loc.City,
	}
}
