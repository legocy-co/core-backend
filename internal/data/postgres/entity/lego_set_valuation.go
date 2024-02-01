package postgres

import (
	"github.com/legocy-co/legocy/internal/domain/calculator/models"
)

type LegoSetValuationPostgres struct {
	Model
	LegoSetID int             `gorm:"uniqueIndex:lego_set_state_unique_idx;not null"`
	LegoSet   LegoSetPostgres `gorm:"foreignKey:LegoSetID;constraint:OnDelete:CASCADE"`
	State     string          `gorm:"uniqueIndex:lego_set_state_unique_idx;not null"`
	Valuation float32         `gorm:"not null"`
}

func (LegoSetValuationPostgres) TableName() string {
	return "lego_sets_valuation"
}

func (e LegoSetValuationPostgres) ToLegoSetValuation() *models.LegoSetValuation {
	return &models.LegoSetValuation{
		ID:               int(e.ID),
		LegoSet:          *e.LegoSet.ToLegoSet(),
		State:            e.State,
		CompanyValuation: e.Valuation,
	}
}

func FromLegoSetValuationVO(vo models.LegoSetValuationValueObject) *LegoSetValuationPostgres {
	return &LegoSetValuationPostgres{
		LegoSetID: vo.LegoSetID,
		State:     vo.State,
		Valuation: vo.CompanyValuation,
	}
}

func GetUpdatedLegoSetValuationPostgres(
	vo models.LegoSetValuationValueObject, entity *LegoSetValuationPostgres) *LegoSetValuationPostgres {

	entity.LegoSetID = vo.LegoSetID
	entity.Valuation = vo.CompanyValuation
	entity.State = vo.State

	return entity
}
