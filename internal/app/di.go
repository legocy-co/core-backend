package app

import (
	"fmt"
	"gorm.io/gorm"
	"legocy-go/internal/config"
	"legocy-go/internal/db"
	p "legocy-go/internal/db/postgres"
	"legocy-go/internal/storage"
	"legocy-go/internal/storage/provider"
	"log"
)

func (a *App) GetDatabase() db.DataBaseConnection {
	return a.database
}

func (a *App) setDatabase(dbCfg *config.DatabaseConfig) {
	var dbConn *gorm.DB
	conn, err := p.CreateConnection(dbCfg, dbConn)
	if err != nil {
		log.Fatalln(fmt.Sprintf("[Database] %v", err.Error()))
		return
	}
	conn.Init()
	a.database = conn
}

func (a *App) GetStorage() storage.ImageStorage {
	return a.imageStorage
}

func (a *App) setStorage(minioCfg config.MinioConfig) {
	imgStorage, err := provider.NewMinioProvider(
		minioCfg.Url,
		minioCfg.User, minioCfg.Password, minioCfg.Token,
		minioCfg.Ssl)
	if err != nil {
		log.Fatalln(fmt.Sprintf("[Minio] %v", err.Error()))
		return
	}
	err = imgStorage.Connect()
	if err != nil {
		log.Fatalln(fmt.Sprintf("[Minio] %v", err.Error()))
	}

	a.imageStorage = imgStorage
}
