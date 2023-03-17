package app

import (
	postgres "legocy-go/internal/db/postgres/repository"
	"legocy-go/internal/domain/auth/repository"
	repository2 "legocy-go/internal/domain/lego/repository"
	marketplace2 "legocy-go/internal/domain/marketplace/repository"
)

func (a *App) GetUserRepo() repository.UserRepository {
	return postgres.NewUserPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserImagesRepo() repository.UserImageRepository {
	return postgres.NewUserImagePostgresRepository(a.GetDatabase())
}

func (a *App) GetLegoSeriesRepo() repository2.LegoSeriesRepository {
	return postgres.NewLegoSeriesPostgresRepository(a.GetDatabase())
}

func (a *App) GetLegoSetRepo() repository2.LegoSetRepository {
	return postgres.NewLegoSetPostgresRepository(a.GetDatabase())
}

func (a *App) GetLocationRepo() marketplace2.LocationRepository {
	return postgres.NewLocationPostgresRepository(a.GetDatabase())
}

func (a *App) GetCurrencyRepo() marketplace2.CurrencyRepository {
	return postgres.NewCurrencyPostgresRepository(a.GetDatabase())
}

func (a *App) GetMarketItemRepo() marketplace2.MarketItemRepository {
	return postgres.NewMarketItemPostgresRepository(a.GetDatabase())
}
