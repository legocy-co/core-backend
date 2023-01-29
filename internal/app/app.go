package app

import (
	"gorm.io/gorm"
	router "legocy-go/api/v1/router"
	"legocy-go/api/v1/usecase/auth"
	lego2 "legocy-go/api/v1/usecase/lego"
	"legocy-go/api/v1/usecase/marketplace"
	"legocy-go/internal/config"
	d "legocy-go/internal/db"
	p "legocy-go/internal/db/postgres"
	repo "legocy-go/internal/db/postgres/repository"
	"legocy-go/internal/storage"
	"legocy-go/internal/storage/provider"
)

type Application interface {
	setup() router.V1router
	Run(port string)
}

type App struct {
	database     d.DataBaseConnection
	imageStorage storage.ImageStorage
}

func New(configFilepath string) Application {
	// Load config
	err := config.SetupFromJSON(configFilepath)
	if err != nil {
		panic(err)
	}

	cfg := config.GetAppConfig()
	if cfg == nil {
		panic("empty config")
	}

	//Database

	dbCfg := config.GetDBConfig()
	var db *gorm.DB
	conn, err := p.CreateConnection(dbCfg, db)
	if err != nil {
		panic(err)
	}
	conn.Init()

	// Item Storage
	minioCfg := config.GetMinioConfig()
	imgStorage, err := provider.NewMinioProvider(
		minioCfg.Url,
		minioCfg.User, minioCfg.Password,
		minioCfg.Ssl)

	if err != nil {
		panic(err)
	}

	return &App{
		database:     conn,
		imageStorage: imgStorage,
	}
}

func (a *App) setup() router.V1router {

	// Repositories
	userRepo := repo.NewUserPostgresRepository(a.database)
	seriesRepo := repo.NewLegoSeriesPostgresRepository(a.database)
	setsRepo := repo.NewLegoSetPostgresRepository(a.database)
	locationRepo := repo.NewLocationPostgresRepository(a.database)

	// Services
	userService := auth.NewUserUsecase(&userRepo)
	seriesService := lego2.NewLegoSeriesService(&seriesRepo)
	setsService := lego2.NewLegoSetUseCase(&setsRepo)
	locationService := marketplace.NewLocationUseCase(&locationRepo)

	// Router
	v1router := router.InitRouter(
		userService,
		seriesService,
		setsService,
		locationService)

	return v1router
}

func (a *App) Run(port string) {
	router := a.setup()
	err := router.Run(port)
	if err != nil {
		panic(err)
	}
}
