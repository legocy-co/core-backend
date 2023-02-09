package storage

import (
	"context"
	models "legocy-go/internal/storage/models"
)

type ImageStorage interface {
	Connect() error // Инициализатор подключения
	IsReady() bool
	UploadFile(context.Context, models.ImageUnit) (string, error)   // Загрузка файлов
	DownloadFile(context.Context, string) (models.ImageUnit, error) // Скачивание файлов
}
