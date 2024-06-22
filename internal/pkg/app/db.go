package app

import (
	"github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/data/postgres"
	"github.com/legocy-co/legocy/internal/pkg/config"
)

func (a *App) GetDatabase() data.Storage {
	return a.database
}

func (a *App) setDatabase(dbCfg *config.DatabaseConfig) error {
	conn, err := postgres.New(dbCfg, a.GetLogger())
	if err != nil {
		return err
	}

	a.database = conn
	return nil
}
