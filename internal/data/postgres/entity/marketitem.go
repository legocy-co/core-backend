package postgres

import (
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

type MarketItemPostgres struct {
	Model
	Price             float32
	LegoSetPostgresID uint                      `filter:"param:setId;searchable,filterable"`
	LegoSet           LegoSetPostgres           `gorm:"ForeignKey:LegoSetPostgresID;"`
	UserPostgresID    uint                      `filter:"param:sellerId;searchable,filterable"`
	Seller            UserPostgres              `gorm:"ForeignKey:UserPostgresID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Location          string                    `gorm:"not null"`
	Status            string                    `gorm:"not null"`
	SetState          string                    `gorm:"not null"`
	Description       string                    `gorm:"not null"`
	Images            []MarketItemImagePostgres `gorm:"foreignKey:MarketItemID"`
	Likes             []MarketItemLikePostgres  `gorm:"foreignKey:MarketItemID"`
}

func (mp MarketItemPostgres) TableName() string {
	return "market_items"
}

func (mp *MarketItemPostgres) ToMarketItem() (*models.MarketItem, *errors.AppError) {

	images := make([]*models.MarketItemImage, 0, len(mp.Images))
	for _, img := range mp.Images {
		images = append(images, img.ToMarketItemImage())
	}

	return models.NewMarketItem(
		int(mp.ID),
		*mp.LegoSet.ToLegoSet(),
		*mp.Seller.ToUser(),
		mp.Price,
		mp.Location,
		mp.SetState,
		mp.Status,
		mp.Description,
		images,
		len(mp.Likes) != 0,
	)
}

func FromMarketItemValueObject(mi *models.MarketItemValueObject) *MarketItemPostgres {
	return &MarketItemPostgres{
		Price:             mi.Price,
		LegoSetPostgresID: uint(mi.LegoSetID),
		UserPostgresID:    uint(mi.SellerID),
		Location:          mi.Location,
		Status:            mi.Status,
		SetState:          mi.SetState,
		Description:       mi.Description,
	}
}

func (mp *MarketItemPostgres) GetUpdatedMarketItem(
	vo models.MarketItemValueObject) *MarketItemPostgres {
	mp.LegoSetPostgresID = uint(vo.LegoSetID)
	mp.Price = vo.Price
	mp.Location = vo.Location
	mp.UserPostgresID = uint(vo.SellerID)
	mp.Status = vo.Status
	mp.SetState = vo.SetState
	mp.Description = vo.Description

	return mp
}

func (mp *MarketItemPostgres) GetUpdatedMarketItemAdmin(
	vo models.MarketItemAdminValueObject) *MarketItemPostgres {
	mp.LegoSetPostgresID = uint(vo.LegoSetID)
	mp.Location = vo.Location
	mp.Price = vo.Price
	mp.UserPostgresID = uint(vo.SellerID)
	mp.Status = models.ListingStatusActive
	mp.SetState = vo.SetState
	mp.Description = vo.Description

	return mp
}

func FromMarketItemAdminValueObject(vo models.MarketItemAdminValueObject) *MarketItemPostgres {
	return &MarketItemPostgres{
		Price:             vo.Price,
		LegoSetPostgresID: uint(vo.LegoSetID),
		UserPostgresID:    uint(vo.SellerID),
		Location:          vo.Location,
		Status:            vo.Status,
		SetState:          vo.SetState,
		Description:       vo.Description,
	}
}

func (mp *MarketItemPostgres) ToMarketItemAdmin() *models.MarketItemAdmin {

	images := make([]*models.MarketItemImage, 0, len(mp.Images))
	for _, img := range mp.Images {
		images = append(images, img.ToMarketItemImage())
	}

	return &models.MarketItemAdmin{
		ID:          int(mp.ID),
		LegoSet:     *mp.LegoSet.ToLegoSet(),
		Seller:      *mp.Seller.ToUser(),
		Price:       mp.Price,
		Location:    mp.Location,
		Status:      mp.Status,
		SetState:    mp.SetState,
		Description: mp.Description,
		Images:      images,
	}
}
