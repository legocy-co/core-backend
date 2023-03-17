package app

import (
	auth2 "legocy-go/internal/delievery/http/usecase/auth"
	lego2 "legocy-go/internal/delievery/http/usecase/lego"
	marketplace2 "legocy-go/internal/delievery/http/usecase/marketplace"
)

func (a *App) GetUserService() auth2.UserUseCase {
	return auth2.NewUserUsecase(a.GetUserRepo())
}

func (a *App) GetUserImagesService() auth2.UserImageUseCase {
	return auth2.NewUserImageUseCase(a.GetUserImagesRepo())
}

func (a *App) GetLegoSeriesService() lego2.LegoSeriesService {
	return lego2.NewLegoSeriesService(a.GetLegoSeriesRepo())
}

func (a *App) GetLegoSetService() lego2.LegoSetUseCase {
	return lego2.NewLegoSetUseCase(a.GetLegoSetRepo())
}

func (a *App) GetLocationService() marketplace2.LocationUseCase {
	return marketplace2.NewLocationUseCase(a.GetLocationRepo())
}

func (a *App) GetCurrencyService() marketplace2.CurrencyUseCase {
	return marketplace2.NewCurrencyUseCase(a.GetCurrencyRepo())
}

func (a *App) GetMarketItemService() marketplace2.MarketItemService {
	return marketplace2.NewMarketItemSerivce(a.GetMarketItemRepo())
}
