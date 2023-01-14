package main

import (
	h "legocy-go/api/v1/handlers"
	r "legocy-go/api/v1/router"
	s "legocy-go/api/v1/usecase"
	config "legocy-go/config"
	p "legocy-go/infrastructure/db/postgres"
	repo "legocy-go/infrastructure/db/postgres/repository"
	"log"

	"github.com/jinzhu/gorm"
)

func main() {

	// Config

	err := config.SetupFromJSON("config/json/config.json")
	if err != nil {
		log.Fatalln(err)
		return
	}

	cfg := config.GetAppConfig()
	if cfg == nil {
		log.Fatalln("Config Empty")
		return
	}

	// Database
	dbCfg := cfg.DbConf
	var db *gorm.DB
	conn, err := p.CreateConnection(&dbCfg, db)
	if err != nil {
		log.Fatalln(err)
		return
	}
	conn.Init()

	// Repositories
	userRepo := repo.NewUserPostgresRepository(conn)
	seriesRepo := repo.NewLegoSeriesPostgresRepository(conn)

	// Services

	userService := s.NewUserUsecase(&userRepo)
	seriesService := s.NewLegoSeriesService(&seriesRepo)

	// Handlers
	tokenHandler := h.NewTokenHandler(userService)
	seriesHandler := h.NewLegoSeriesHandler(seriesService)

	// Router
	router := r.InitRouter(tokenHandler, seriesHandler)
	router.Run(":" + "8080")

}
