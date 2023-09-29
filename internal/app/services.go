package app

import (
	lego "legocy-go/internal/domain/lego/service"
	marketplace "legocy-go/internal/domain/marketplace/service"
	"legocy-go/internal/domain/marketplace/service/admin"
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

func (a *App) GetLocationService() marketplace.LocationUseCase {
	return marketplace.NewLocationUseCase(a.GetLocationRepo())
}

func (a *App) GetCurrencyService() marketplace.CurrencyUseCase {
	return marketplace.NewCurrencyUseCase(a.GetCurrencyRepo())
}

func (a *App) GetMarketItemService() marketplace.MarketItemService {
	return marketplace.NewMarketItemService(a.GetMarketItemRepo())
}

func (a *App) GetUserReviewService() marketplace.UserReviewService {
	return marketplace.NewUserReviewService(a.GetUserReviewRepo())
}

func (a *App) GetMarketItemAdminService() admin.MarketItemAdminService {
	return admin.NewMarketItemAdminService(a.GetMarketItemAdminRepository())
}
