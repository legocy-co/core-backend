package postgres

import (
	"legocy-go/internal/domain/collections/models"
)

type LegoSetValuation struct {
	Model
	LegoSetID  int              `gorm:"uniqueIndex:lego_set_currency_unique_idx; not null"`
	LegoSet    LegoSetPostgres  `gorm:"foreignKey:LegoSetID;constraint:OnDelete:CASCADE"`
	CurrencyID int              `gorm:"uniqueIndex:lego_set_currency_unique_idx; not null"`
	Currency   CurrencyPostgres `gorm:"foreignKey:CurrencyID;constraint:OnDelete: SET NULL"`
	State      string           `gorm:"not null;"`
	Valuation  float32          `gorm:"not null"`
}

func (LegoSetValuation) TableName() string {
	return "lego_sets_valuation"
}

func (e LegoSetValuation) ToLegoSetValuation() *models.LegoSetValuation {
	return &models.LegoSetValuation{
		ID:               int(e.ID),
		LegoSet:          *e.LegoSet.ToLegoSet(),
		State:            e.State,
		CompanyValuation: e.Valuation,
		Currency:         *e.Currency.ToCurrency(),
	}
}
