package app

import (
	"fmt"
	"gorm.io/gorm"
	"legocy-go/config"
	"legocy-go/internal/data"
	postgres "legocy-go/internal/data/postgres"
	storage "legocy-go/pkg/storage/client"
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

func (a *App) GetImageStorageClient() storage.ImageStorage {
	return storage.NewImageStorage(config.GetAppConfig().S3Port)
}
