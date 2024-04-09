package app

import (
	"github.com/legocy-co/legocy/internal/pkg/config"
	storage "github.com/legocy-co/legocy/pkg/s3/client"
)

func (a *App) GetImageStorageClient() storage.ImageStorage {
	return storage.NewImageStorage(config.GetAppConfig().S3Host, config.GetAppConfig().S3Port)
}
