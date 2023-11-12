package app

import (
	postgres "legocy-go/internal/data/postgres/repository"
	postgresAdmin "legocy-go/internal/data/postgres/repository/admin"
	calculator "legocy-go/internal/domain/calculator/repository"
	collections "legocy-go/internal/domain/collections/repository"
	lego "legocy-go/internal/domain/lego/repository"
	marketplace "legocy-go/internal/domain/marketplace/repository"
	users "legocy-go/internal/domain/users/repository"
)

// Start Admin

func (a *App) GetMarketItemAdminRepository() marketplace.MarketItemAdminRepository {
	return postgresAdmin.NewMarketItemAdminPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserAdminRepository() users.UserAdminRepository {
	return postgresAdmin.NewUserAdminPostgresRepository(a.GetDatabase())
}

func (a *App) GetCurrencyRepo() marketplace.CurrencyRepository {
	return postgresAdmin.NewCurrencyPostgresRepository(a.GetDatabase())
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

func (a *App) GetLocationRepo() marketplace.LocationRepository {
	return postgres.NewLocationPostgresRepository(a.GetDatabase())
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
