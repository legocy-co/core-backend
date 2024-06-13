package app

import (
	postgres "github.com/legocy-co/legocy/internal/data/postgres/repository"
	postgresAdmin "github.com/legocy-co/legocy/internal/data/postgres/repository/admin"
	"github.com/legocy-co/legocy/internal/data/postgres/repository/facebook"
	"github.com/legocy-co/legocy/internal/data/postgres/repository/google"
	calculator "github.com/legocy-co/legocy/internal/domain/calculator/repository"
	calculatorAdmin "github.com/legocy-co/legocy/internal/domain/calculator/repository/admin"
	collections "github.com/legocy-co/legocy/internal/domain/collections/repository"
	lego "github.com/legocy-co/legocy/internal/domain/lego/repository"
	marketplace "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
	users "github.com/legocy-co/legocy/internal/domain/users/repository"
	"github.com/legocy-co/legocy/internal/pkg/di"
)

// Start Admin

func (a *App) GetMarketItemAdminRepository() marketplace.MarketItemAdminRepository {
	return postgresAdmin.NewMarketItemAdminPostgresRepository(a.GetDatabase(), di.ProvideDispatcher())
}

func (a *App) GetUserAdminRepository() users.UserAdminRepository {
	return postgresAdmin.NewUserAdminPostgresRepository(a.GetDatabase(), di.ProvideDispatcher())
}

func (a *App) GetLegoSetsValuationAdminRepository() calculatorAdmin.LegoSetValuationAdminRepository {
	return postgresAdmin.NewLegoSetValuationPostgresAdminRepository(a.GetDatabase())
}

// End Admin

func (a *App) GetUserRepo() users.UserRepository {
	return postgres.NewUserPostgresRepository(a.GetDatabase(), di.ProvideDispatcher())
}

func (a *App) GetUserImagesRepo() users.UserImageRepository {
	return postgres.NewUserImagePostgresRepository(a.GetDatabase(), di.ProvideDispatcher())
}

func (a *App) GetLegoSeriesRepo() lego.LegoSeriesRepository {
	return postgres.NewLegoSeriesPostgresRepository(a.GetDatabase())
}

func (a *App) GetLegoSetRepo() lego.LegoSetRepository {
	return postgres.NewLegoSetPostgresRepository(a.GetDatabase())
}

func (a *App) GetMarketItemRepo() marketplace.MarketItemRepository {
	return postgres.NewMarketItemPostgresRepository(a.GetDatabase(), di.ProvideDispatcher())
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

func (a *App) GetMarketItemImageRepository() marketplace.MarketItemImageRepository {
	return postgres.NewMarketItemImagePostgresRepository(a.GetDatabase(), di.ProvideDispatcher())
}

func (a *App) GetLegoSetImageRepository() lego.LegoSetImageRepository {
	return postgres.NewLegoSetImagePostgresRepository(a.GetDatabase(), di.ProvideDispatcher())
}

func (a *App) GetMarketItemLikeRepository() marketplace.LikeRepository {
	return postgres.NewLikePostgresRepository(a.GetDatabase())
}

func (a *App) GetGoogleAuthRepository() users.UserExternalAuthRepository {
	return google.NewUserAuthRepository(a.GetDatabase(), di.ProvideDispatcher())
}

func (a *App) GetFacebookAuthRepository() users.UserExternalAuthRepository {
	return facebook.NewUserAuthRepository(a.GetDatabase(), di.ProvideDispatcher())
}
