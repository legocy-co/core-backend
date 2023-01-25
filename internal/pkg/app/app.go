package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	r "legocy-go/api/v1/router"
	s "legocy-go/api/v1/usecase"
	"legocy-go/config"
	p "legocy-go/infrastructure/db/postgres"
	repo "legocy-go/infrastructure/db/postgres/repository"
	"log"
)

type Application interface {
	setup() *gin.Engine
	Run(port string)
}

func New(configFilepath string) Application {
	return &App{configFilepath: configFilepath}
}

type App struct {
	configFilepath string
}

func (a *App) setup() *gin.Engine {
	// Config

	err := config.SetupFromJSON(a.configFilepath)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	cfg := config.GetAppConfig()
	if cfg == nil {
		log.Fatalln("Config Empty")
		return nil
	}

	// Database
	dbCfg := cfg.DbConf
	var db *gorm.DB
	conn, err := p.CreateConnection(&dbCfg, db)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	conn.Init()

	// Repositories
	userRepo := repo.NewUserPostgresRepository(conn)
	seriesRepo := repo.NewLegoSeriesPostgresRepository(conn)

	// Services
	userService := s.NewUserUsecase(&userRepo)
	seriesService := s.NewLegoSeriesService(&seriesRepo)

	// Router
	router := r.InitRouter(userService, seriesService)

	return router
}

func (a *App) Run(port string) {
	router := a.setup()
	router.Run(":" + port)
}
