package app

import (
	postgres "legocy-go/internal/db/postgres/repository"
	auth "legocy-go/pkg/auth/repository"
	lego "legocy-go/pkg/lego/repository"
	marketplace "legocy-go/pkg/marketplace/repository"
)

func (a *App) GetUserRepo() auth.UserRepository {
	return postgres.NewUserPostgresRepository(a.GetDatabase())
}

func (a *App) GetUserImagesRepo() auth.UserImageRepository {
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
