package app

import (
	calculator "github.com/legocy-co/legocy/internal/domain/calculator/service"
	calculatorAdmin "github.com/legocy-co/legocy/internal/domain/calculator/service/admin"
	collection "github.com/legocy-co/legocy/internal/domain/collections/service/collection"
	lego "github.com/legocy-co/legocy/internal/domain/lego/service"
	marketplace "github.com/legocy-co/legocy/internal/domain/marketplace/service"
	marketplaceAdmin "github.com/legocy-co/legocy/internal/domain/marketplace/service/admin"
	users "github.com/legocy-co/legocy/internal/domain/users/service"
	usersAdmin "github.com/legocy-co/legocy/internal/domain/users/service/admin"
)

// Start Admin

func (a *App) GetUserAdminService() usersAdmin.UserAdminService {
	return usersAdmin.NewUserAdminService(a.GetUserAdminRepository())
}

func (a *App) GetMarketItemAdminService() marketplaceAdmin.MarketItemAdminService {
	return marketplaceAdmin.NewMarketItemAdminService(a.GetMarketItemAdminRepository())
}

func (a *App) GetLegoSetValuationAdminService() calculatorAdmin.LegoSetValuationAdminService {
	return calculatorAdmin.NewService(a.GetLegoSetsValuationAdminRepository())
}

// End Admin

func (a *App) GetUserService() users.UserService {
	return users.NewUserService(a.GetUserRepo())
}

func (a *App) GetUserImagesService() users.UserImageService {
	return users.NewUserImageUseCase(a.GetUserImagesRepo())
}

func (a *App) GetLegoSeriesService() lego.LegoSeriesService {
	return lego.NewLegoSeriesService(a.GetLegoSeriesRepo())
}

func (a *App) GetLegoSetService() lego.LegoSetService {
	return lego.NewLegoSetService(a.GetLegoSetRepo())
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

func (a *App) GetMarketItemImageService() marketplace.MarketItemImageService {
	return marketplace.NewMarketItemImageService(a.GetMarketItemImageRepository())
}

func (a *App) GetLegoSetValuationService() calculator.LegoSetValuationService {
	return calculator.NewLegoSetValuationService(a.GetLegoSetsValuationRepository())
}
