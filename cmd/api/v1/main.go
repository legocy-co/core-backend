package main

import (
	"fmt"
	r "legocy-go/api/v1/router"
	s "legocy-go/api/v1/usecase"
	config "legocy-go/config"
	p "legocy-go/infrastructure/db/postgres"
	repo "legocy-go/infrastructure/db/postgres/repository"
	"log"

	"gorm.io/gorm"
)

func main() {

	// Config

	err := config.SetupFromJSON("/Users/wjojf/Documents/dev/legocy-go-clean/config/json/config.json")
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
	fmt.Println(conn)

	// Repositories
	userRepo := repo.NewUserPostgresRepository(conn)
	seriesRepo := repo.NewLegoSeriesPostgresRepository(conn)

	// Services
	userService := s.NewUserUsecase(&userRepo)
	seriesService := s.NewLegoSeriesService(&seriesRepo)

	// Router
	router := r.InitRouter(userService, seriesService)
	router.Run(":" + "8080")

}
