package app

import (
	postgres "legocy-go/internal/db/postgres/repository"
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
