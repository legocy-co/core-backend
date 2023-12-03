package postgres

import (
	"github.com/legocy-co/legocy/internal/domain/collections/models"
)

type UserLegoSetPostgres struct {
	Model
	UserID    int             `gorm:"index;not null"`
	User      UserPostgres    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	LegoSetID int             `gorm:"index; not null"`
	LegoSet   LegoSetPostgres `gorm:"foreignKey:LegoSetID;constraint:OnDelete:CASCADE"`
	State     string          `gorm:"not null;"`
	BuyPrice  float32         `gorm:"not null;type:decimal(25,2);"`
}

func (UserLegoSetPostgres) TableName() string {
	return "users_lego_sets"
}

func (lsp UserLegoSetPostgres) ToCollectionLegoSet() models.CollectionLegoSet {
	return models.CollectionLegoSet{
		ID:           int(lsp.ID),
		LegoSet:      *lsp.LegoSet.ToLegoSet(),
		CurrentState: lsp.State,
		BuyPrice:     lsp.BuyPrice,
	}
}

func GetUpdatedUserLegoSet(vo *models.CollectionLegoSetValueObject, entity *UserLegoSetPostgres, userID int) *UserLegoSetPostgres {
	entity.LegoSetID = vo.LegoSetID
	entity.UserID = userID
	entity.State = vo.CurrentState
	entity.BuyPrice = vo.BuyPrice

	return entity
}

func GetCreatedUserLegoSet(vo *models.CollectionLegoSetValueObject, userID int) *UserLegoSetPostgres {
	return &UserLegoSetPostgres{
		UserID:    userID,
		LegoSetID: vo.LegoSetID,
		State:     vo.CurrentState,
		BuyPrice:  vo.BuyPrice,
	}
}
