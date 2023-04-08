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
	if err != nil {
		log.Fatalf("did not connect %v", err)
		return "", storage.ErrConnectionRefused
	}

	defer conn.Close()

	client := proto.NewS3ServiceClient(conn)

	request := mapper.GetImageUploadRequest(image, bucketName)
	response, err := client.UploadImage(context.Background(), request)
	if err != nil {
		return "", err
	}

	return response.ImageURL, nil
}

func (s ImageStorage) DownloadImage(bucketName, imageName string) (*models.ImageUnit, error) {
	conn, err := s.getConnection()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := proto.NewS3ServiceClient(conn)

	request := mapper.GetImageDownloadRequest(bucketName, imageName)
	imageResponse, err := client.DownloadImage(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return mapper.DownloadImageResponseToImageUnit(imageResponse), nil

}
