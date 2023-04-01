package app

import (
	"fmt"
	"gorm.io/gorm"
	"legocy-go/internal/config"
	"legocy-go/internal/db"
	postgres "legocy-go/internal/db/postgres"
	"legocy-go/pkg/storage/client"
	"log"
)

func (a *App) GetDatabase() db.DataBaseConnection {
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

func (a *App) GetImageStorageClient() client.ImageStorage {
	return client.NewImageStorage(config.GetAppConfig().S3Port)
}
