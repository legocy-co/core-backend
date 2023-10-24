package app

import (
	postgres "legocy-go/internal/data/postgres/repository"
	"legocy-go/internal/data/postgres/repository/admin"
	collections "legocy-go/internal/domain/collections/repository"
	lego "legocy-go/internal/domain/lego/repository"
	marketplace "legocy-go/internal/domain/marketplace/repository"
	"legocy-go/internal/domain/users/repository"
)

func (a *App) GetUserRepo() repository.UserRepository {
	return postgres.NewUserPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserImagesRepo() repository.UserImageRepository {
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

func (a *App) GetCurrencyRepo() marketplace.CurrencyRepository {
	return postgres.NewCurrencyPostgresRepository(a.GetDatabase())
}

func (a *App) GetMarketItemRepo() marketplace.MarketItemRepository {
	return postgres.NewMarketItemPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserReviewRepo() marketplace.UserReviewRepository {
	return postgres.NewUserReviewPostgresRepository(a.GetDatabase())
}

func (a *App) GetMarketItemAdminRepository() marketplace.MarketItemAdminRepository {
	return admin.NewMarketItemAdminPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserAdminRepository() repository.UserAdminRepository {
	return admin.NewUserAdminPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserLegoSetsRepository() collections.UserCollectionRepository {
	return postgres.NewCollectionPostgresRepository(a.GetDatabase())
}

func (a *App) GetLegoSetsValuationRepository() collections.LegoSetValuationRepository {
	return postgres.NewLegoSetValuationPostgresRepository(a.GetDatabase())
}
