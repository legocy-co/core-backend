package postgres

type LegoSetValuation struct {
	Model
	LegoSetID  int              `gorm:"uniqueIndex:lego_set_currency_unique_idx; not null"`
	LegoSet    LegoSetPostgres  `gorm:"foreignKey:LegoSetID;constraint:OnDelete:CASCADE"`
	CurrencyID int              `gorm:"uniqueIndex:lego_set_currency_unique_idx; not null"`
	Currency   CurrencyPostgres `gorm:"foreignKey:CurrencyID;constraint:OnDelete: SET NULL"`
	Valuation  float32          `gorm:"not null"`
}

func (LegoSetValuation) TableName() string {
	return "lego_sets_valuation"
}
