package app

import (
	"gorm.io/gorm"
	router "legocy-go/api/v1/router"
	"legocy-go/api/v1/usecase/auth"
	lego2 "legocy-go/api/v1/usecase/lego"
	"legocy-go/api/v1/usecase/marketplace"
	"legocy-go/internal/config"
	p "legocy-go/internal/db/postgres"
	repo "legocy-go/internal/db/postgres/repository"
)

type Application interface {
	setup() router.V1router
	Run(port string)
}

func New(configFilepath string) Application {
	return &App{cfg: configFilepath}
}

type App struct {
	cfg     string
	isSetUp bool
}

func (a *App) setup() router.V1router {
	// Config

	err := config.SetupFromJSON(a.cfg)
	if err != nil {
		panic(err)
	}

	cfg := config.GetAppConfig()
	if cfg == nil {
		panic("empty config")
	}

	// Database
	dbCfg := cfg.DbConf
	var db *gorm.DB
	conn, err := p.CreateConnection(&dbCfg, db)
	if err != nil {
		panic(err)
	}
	conn.Init()

	// Repositories
	userRepo := repo.NewUserPostgresRepository(conn)
	seriesRepo := repo.NewLegoSeriesPostgresRepository(conn)
	setsRepo := repo.NewLegoSetPostgresRepository(conn)
	locationRepo := repo.NewLocationPostgresRepository(conn)

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
