package postgres

import models "github.com/legocy-co/legocy/internal/domain/lego/models"

type LegoSetImagePostgres struct {
	Model
	LegoSetID int             `gorm:"not null"`
	LegoSet   LegoSetPostgres `gorm:"foreignKey:LegoSetID;constraint:OnDelete:CASCADE;"`
	ImageURL  string          `gorm:"not null"`
	IsMain    bool            `gorm:"default=false"`
}

func (i LegoSetImagePostgres) TableName() string {
	return "lego_set_images"
}

func (i LegoSetImagePostgres) ToLegoSetImage() *models.LegoSetImage {
	return &models.LegoSetImage{
		ID:        int(i.ID),
		LegoSetID: int(i.LegoSetID),
		ImageURL:  i.ImageURL,
		IsMain:    i.IsMain,
	}
}

func FromLegoSetImageValueObject(vo models.LegoSetImageValueObject) *LegoSetImagePostgres {
	return &LegoSetImagePostgres{
		LegoSetID: vo.LegoSetID,
		ImageURL:  vo.ImageURL,
		IsMain:    vo.IsMain,
	}
}
