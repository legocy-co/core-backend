package app

import (
	"fmt"
	"github.com/joho/godotenv"
	d "github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/pkg/config"
	log "github.com/sirupsen/logrus"
)

type App struct {
	database d.Storage
}

func New() (*App, error) {

	godotenv.Load()

	app := &App{}

	// Load config
	err := config.SetupFromEnv()
	if err != nil {
		log.Fatalln(fmt.Sprintf("[Config] %v", err.Error()))
	}

	cfg := config.GetAppConfig()
	if cfg == nil {
		log.Fatalln("Error getting app config")
	}

	//Database
	dbCfg := config.GetDBConfig()
	if err := app.setDatabase(dbCfg); err != nil {
		return nil, err
	}

	return app, nil
}
