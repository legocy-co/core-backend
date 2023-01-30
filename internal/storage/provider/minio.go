package provider

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"legocy-go/internal/storage"
	"log"
)

// MinioProvider - Наш провайдер для хранилища
type MinioProvider struct {
	minioAuthData
	client *minio.Client
}

type minioAuthData struct {
	url      string
	user     string
	password string
	token    string
	ssl      bool
}

func (m *MinioProvider) Connect() error {
	var err error

	// if already connected - return
	if m.client != nil {
		return nil
	}

	m.client, err = minio.New(m.url, &minio.Options{
		Creds:  credentials.NewStaticV4(m.user, m.password, ""),
		Secure: m.ssl,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func NewMinioProvider(minioURL string, minioUser string, minioPassword string, ssl bool) (storage.ImageStorage, error) {
	return &MinioProvider{
		minioAuthData: minioAuthData{
			password: minioPassword,
			url:      minioURL,
			user:     minioUser,
			ssl:      ssl,
		}}, nil
}
