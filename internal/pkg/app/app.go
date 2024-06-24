package app

import (
	"github.com/joho/godotenv"
	d "github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/pkg/config"
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
		return nil, err
	}

	cfg := config.GetAppConfig()
	if cfg == nil {
		return nil, config.ErrConfigNotFound
	}

	//Database
	dbCfg := config.GetDBConfig()
	if err := app.setDatabase(dbCfg); err != nil {
		return nil, err
	}

	return app, nil
}
