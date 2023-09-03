package app

import (
	lego "legocy-go/internal/domain/lego/service"
	maketplace "legocy-go/internal/domain/marketplace/service"
	users "legocy-go/internal/domain/users/service"
)

func (a *App) GetUserService() users.UserUseCase {
	return users.NewUserUsecase(a.GetUserRepo())
}

func (a *App) GetUserImagesService() users.UserImageUseCase {
	return users.NewUserImageUseCase(a.GetUserImagesRepo())
}

func (a *App) GetLegoSeriesService() lego.LegoSeriesService {
	return lego.NewLegoSeriesService(a.GetLegoSeriesRepo())
}

func (a *App) GetLegoSetService() lego.LegoSetUseCase {
	return lego.NewLegoSetUseCase(a.GetLegoSetRepo())
}

func (a *App) GetLocationService() maketplace.LocationUseCase {
	return maketplace.NewLocationUseCase(a.GetLocationRepo())
}

func (a *App) GetCurrencyService() maketplace.CurrencyUseCase {
	return maketplace.NewCurrencyUseCase(a.GetCurrencyRepo())
}

func (a *App) GetMarketItemService() maketplace.MarketItemService {
	return maketplace.NewMarketItemService(a.GetMarketItemRepo())
}

func (a *App) GetUserReviewService() maketplace.UserReviewService {
	return maketplace.NewUserReviewService(a.GetUserReviewRepo())
}
