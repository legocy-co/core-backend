package filter

import (
	"github.com/legocy-co/legocy/pkg/filter"
	"gorm.io/gorm"
)

type GormAndSpecification[T any] struct {
	db   *gorm.DB
	spec filter.AndSpecification[T]
	dto  gorm.Model
}

func (gs *GormAndSpecification[T]) IsSatisfiedBy(item T) bool {
	return gs.spec.IsSatisfiedBy(item)
}

func (gs *GormAndSpecification[T]) Apply(db *gorm.DB) *gorm.DB {
	return db
}
