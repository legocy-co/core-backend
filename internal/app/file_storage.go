package app

import (
	"legocy-go/config"
	storage "legocy-go/pkg/storage/client"
)

func (a *App) GetImageStorageClient() storage.ImageStorage {
	return storage.NewImageStorage(config.GetAppConfig().S3Host, config.GetAppConfig().S3Port)
}
