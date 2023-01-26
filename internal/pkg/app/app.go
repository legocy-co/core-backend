package app

import (
	"gorm.io/gorm"
	r "legocy-go/api/v1/router"
	s "legocy-go/api/v1/usecase/auth"
	"legocy-go/api/v1/usecase/lego"
	"legocy-go/config"
	p "legocy-go/infrastructure/db/postgres"
	repo "legocy-go/infrastructure/db/postgres/repository"
)

type Application interface {
	setup() r.V1router
	Run(port string)
}

func New(configFilepath string) Application {
	return &App{configFilepath: configFilepath}
}

type App struct {
	configFilepath string
}

func (a *App) setup() r.V1router {
	// Config

	err := config.SetupFromJSON(a.configFilepath)
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

	// Services
	userService := s.NewUserUsecase(&userRepo)
	seriesService := lego.NewLegoSeriesService(&seriesRepo)
	setsService := lego.NewLegoSetUseCase(&setsRepo)

	// Router
	v1router := r.InitRouter(
		userService,
		seriesService,
		setsService)

	return v1router
}

func (a *App) Run(port string) {
	router := a.setup()
	err := router.Run(port)
	if err != nil {
		panic(err)
	}
}
