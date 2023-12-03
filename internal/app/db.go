package app

import (
	"fmt"
	"github.com/legocy-co/legocy/config"
	"github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/data/postgres"
	"gorm.io/gorm"
	"log"
)

func (a *App) GetDatabase() data.DataBaseConnection {
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
