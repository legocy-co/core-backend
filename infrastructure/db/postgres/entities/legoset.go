package postgres

type LegoSetPostgres struct {
	Model
	Number   int    `gorm:"unique"`
	Name     string `gorm:"unique"`
	NPieces  int
	SeriesID int `gorm:"not null"`
}
