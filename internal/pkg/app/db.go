package app

import (
	"fmt"
	"github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/data/postgres"
	"github.com/legocy-co/legocy/internal/pkg/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (a *App) GetDatabase() data.DBConn {
	return a.database
}

func (a *App) setDatabase(dbCfg *config.DatabaseConfig) {
	var dbConn *gorm.DB
	conn, err := postgres.CreateConnection(dbCfg, dbConn)
	if err != nil {
		log.Fatalln(fmt.Sprintf("[Database] %v", err.Error()))
		return
	}
	conn.Init()
	a.database = conn
}
