package provider

import (
	"context"
	"github.com/minio/minio-go/v7"
	"legocy-go/internal/storage/models"
	"log"
)

// UploadFile - Отправляет файл в minio
func (m *MinioProvider) UploadFile(ctx context.Context, object models.ImageUnit) (string, error) {
	// Получаем "уникальное" имя объекта для загружаемого фото
	imageName := object.GenerateObjectName()

	_, err := m.client.PutObject(
		ctx,
		UserObjectsBucketName, // Константа с именем бакета
		imageName,
		object.Payload,
		object.PayloadSize,
		minio.PutObjectOptions{ContentType: "image/png"},
	)

	return imageName, err
}

// DownloadFile - Возвращает файл из minio
func (m *MinioProvider) DownloadFile(ctx context.Context, image string) (models.ImageUnit, error) {
	reader, err := m.client.GetObject(
		ctx,
		UserObjectsBucketName,
		image,
		minio.GetObjectOptions{},
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	return models.ImageUnit{}, nil
}
