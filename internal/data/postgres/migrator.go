package postgres

import entities "github.com/legocy-co/legocy/internal/data/postgres/entity"

func (c *Connection) applyMigrations() error {
	return c.db.Debug().AutoMigrate(
		entities.UserPostgres{},
		entities.UserImagePostgres{},

		entities.LegoSeriesPostgres{},
		entities.LegoSetPostgres{},
		entities.LegoSetValuationPostgres{},

		entities.MarketItemPostgres{},
		entities.MarketItemImagePostgres{},
		entities.MarketItemLikePostgres{},

		entities.UserImagePostgres{},

		entities.UserReviewPostgres{},
		entities.UserLegoSetPostgres{},

		entities.LegoSetImagePostgres{},
	)
}
