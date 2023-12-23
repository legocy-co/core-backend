package app

import (
	postgres "github.com/legocy-co/legocy/internal/data/postgres/repository"
	postgresAdmin "github.com/legocy-co/legocy/internal/data/postgres/repository/admin"
	calculator "github.com/legocy-co/legocy/internal/domain/calculator/repository"
	calculatorAdmin "github.com/legocy-co/legocy/internal/domain/calculator/repository/admin"
	collections "github.com/legocy-co/legocy/internal/domain/collections/repository"
	lego "github.com/legocy-co/legocy/internal/domain/lego/repository"
	marketplace "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
	users "github.com/legocy-co/legocy/internal/domain/users/repository"
)

// Start Admin

func (a *App) GetMarketItemAdminRepository() marketplace.MarketItemAdminRepository {
	return postgresAdmin.NewMarketItemAdminPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserAdminRepository() users.UserAdminRepository {
	return postgresAdmin.NewUserAdminPostgresRepository(a.GetDatabase())
}

func (a *App) GetLegoSetsValuationAdminRepository() calculatorAdmin.LegoSetValuationAdminRepository {
	return postgresAdmin.NewLegoSetValuationPostgresAdminRepository(a.GetDatabase())
}

// End Admin

func (a *App) GetUserRepo() users.UserRepository {
	return postgres.NewUserPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserImagesRepo() users.UserImageRepository {
	return postgres.NewUserImagePostgresRepository(a.GetDatabase())
}

func (a *App) GetLegoSeriesRepo() lego.LegoSeriesRepository {
	return postgres.NewLegoSeriesPostgresRepository(a.GetDatabase())
}

func (a *App) GetLegoSetRepo() lego.LegoSetRepository {
	return postgres.NewLegoSetPostgresRepository(a.GetDatabase())
}

func (a *App) GetMarketItemRepo() marketplace.MarketItemRepository {
	return postgres.NewMarketItemPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserReviewRepo() marketplace.UserReviewRepository {
	return postgres.NewUserReviewPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserLegoSetsRepository() collections.UserCollectionRepository {
	return postgres.NewCollectionPostgresRepository(a.GetDatabase())
}

func (a *App) GetLegoSetsValuationRepository() calculator.LegoSetValuationRepository {
	return postgres.NewLegoSetValuationPostgresRepository(a.GetDatabase())
}

func (a *App) GetMarketItemImageRepository() marketplace.MarketItemImageRepository {
	return postgres.NewMarketItemImagePostgresRepository(a.GetDatabase())
}
