package utils

import (
	"github.com/legocy-co/legocy/lib/pagination"
	"gorm.io/gorm"
)

func AddPaginationQuery(db *gorm.DB, ctx pagination.PaginationContext) *gorm.DB {
	return db.Limit(ctx.GetLimit()).Offset(ctx.GetOffset())
}
