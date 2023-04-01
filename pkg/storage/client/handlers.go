package client

import (
	"context"
	"legocy-go/pkg/storage"
	"legocy-go/pkg/storage/models"
	"legocy-go/pkg/storage/proto"
	"legocy-go/pkg/storage/proto/mapper"
	"log"
)

func (s ImageStorage) UploadImage(
	image *models.ImageUnit, bucketName string) (string, error) {

	conn, err := s.getConnection()
	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect %v", err)
		return "", storage.ErrConnectionRefused
	}

	client := proto.NewS3ServiceClient(conn)

	request := mapper.GetImageRequest(image, bucketName)
	response, err := client.UploadImage(context.Background(), request)
	if err != nil {
		return "", err
	}

	return response.ImageURL, nil
}
