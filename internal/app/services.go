package app

import (
	collection "legocy-go/internal/domain/collections/service/collection"
	lego "legocy-go/internal/domain/lego/service"
	marketplace "legocy-go/internal/domain/marketplace/service"
	marketplaceAdmin "legocy-go/internal/domain/marketplace/service/admin"
	users "legocy-go/internal/domain/users/service"
	useradmin "legocy-go/internal/domain/users/service/admin"
)

// Start Admin

func (a *App) GetUserAdminService() useradmin.UserAdminService {
	return useradmin.NewUserAdminService(a.GetUserAdminRepository())
}

func (a *App) GetMarketItemAdminService() marketplaceAdmin.MarketItemAdminService {
	return marketplaceAdmin.NewMarketItemAdminService(a.GetMarketItemAdminRepository())
}

// End Admin

func (a *App) GetUserService() users.UserService {
	return users.NewUserService(a.GetUserRepo())
}

func (a *App) GetUserImagesService() users.UserImageUseCase {
	return users.NewUserImageUseCase(a.GetUserImagesRepo())
}

func (a *App) GetLegoSeriesService() lego.LegoSeriesService {
	return lego.NewLegoSeriesService(a.GetLegoSeriesRepo())
}

func (a *App) GetLegoSetService() lego.LegoSetService {
	return lego.NewLegoSetService(a.GetLegoSetRepo())
}

func (a *App) GetLocationService() marketplace.LocationUseCase {
	return marketplace.NewLocationUseCase(a.GetLocationRepo())
}

func (a *App) GetCurrencyService() marketplace.CurrencyService {
	return marketplace.NewCurrencyService(a.GetCurrencyRepo())
}

func (a *App) GetMarketItemService() marketplace.MarketItemService {
	return marketplace.NewMarketItemService(a.GetMarketItemRepo())
}

func (a *App) GetUserReviewService() marketplace.UserReviewService {
	return marketplace.NewUserReviewService(a.GetUserReviewRepo())
}

func (a *App) GetUserCollectionService() collection.UserCollectionService {
	return collection.NewUserCollectionService(
		a.GetUserLegoSetsRepository(),
		a.GetLegoSetsValuationRepository(),
		a.GetUserRepo(),
	)
}
