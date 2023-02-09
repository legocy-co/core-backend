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

func (m *MinioProvider) IsReady() bool {
	return m.client != nil
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
		Creds:  credentials.NewStaticV4(m.user, m.password, m.token),
		Secure: m.ssl,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func NewMinioProvider(minioURL string, minioUser, minioPassword, token string, ssl bool) (storage.ImageStorage, error) {
	//Client will be initialized by Connect() method
	return &MinioProvider{
		minioAuthData: minioAuthData{
			password: minioPassword,
			url:      minioURL,
			user:     minioUser,
			ssl:      ssl,
			token:    token,
		},
		client: nil,
	}, nil
}
