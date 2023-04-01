package app

import (
	"fmt"
	"legocy-go/internal/config"
	d "legocy-go/internal/db"
	"legocy-go/internal/fixtures"
	"log"
)

type App struct {
	database d.DataBaseConnection
}

func (a *App) isReady() bool {
	return a.database.IsReady()
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

	// Fixtures
	go func(load bool) {
		if !load {
			return
		}
		fixtures.LoadLegoSeries(app.GetLegoSeriesRepo())
		fixtures.LoadLegoSets(app.GetLegoSetRepo(), app.GetLegoSeriesRepo())
	}(dbCfg.LoadFixtures)

	// Check all deps
	if !app.isReady() {
		panic("Some dependencies failed to inject")
	}

	return &app
}
