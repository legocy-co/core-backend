package postgres

type UserLegoSetPostgres struct {
	Model
	UserID     int              `gorm:"index;not null"`
	User       UserPostgres     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	LegoSetID  int              `gorm:"index; not null"`
	LegoSet    LegoSetPostgres  `gorm:"foreignKey:LegoSetID;constraint:OnDelete:CASCADE"`
	State      string           `gorm:"not null;"`
	BuyPrice   float32          `gorm:"not null;type:numeric"`
	CurrencyID int              `gorm:"index;not null"`
	Currency   CurrencyPostgres `gorm:"foreignKey:CurrencyID;constraint:OnDelete: SET NULL"`
}

func (UserLegoSetPostgres) TableName() string {
	return "users_lego_sets"
}
