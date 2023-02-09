package app

import (
	"fmt"
	"legocy-go/internal/config"
	d "legocy-go/internal/db"
	"legocy-go/internal/storage"
	"log"
)

type App struct {
	database     d.DataBaseConnection
	imageStorage storage.ImageStorage
}

func (a *App) isReady() bool {
	return a.imageStorage.IsReady() && a.database.IsReady()
}

func New(configFilepath string) *App {

	app := App{}

	// Load config
	err := config.SetupFromJSON(configFilepath)
	if err != nil {
		log.Fatalln(fmt.Sprintf("[Config] %v", err.Error()))
	}

	cfg := config.GetAppConfig()
	if cfg == nil {
		log.Fatalln("Error getting app config")
	}

	//Database
	dbCfg := config.GetDBConfig()
	if cfg == nil {
		log.Fatalln("empty db config")
	}
	app.setDatabase(dbCfg)

	// Item Storage
	minioCfg := config.GetMinioConfig()
	if minioCfg == nil {
		log.Fatalln("empty minio config")
	}
	app.setStorage(*minioCfg)

	if !app.isReady() {
		panic("Some dependencies failed to inject")
	}

	return &app
}
