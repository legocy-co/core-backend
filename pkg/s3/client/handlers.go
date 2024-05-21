package client

import (
	"context"
	"github.com/legocy-co/legocy/pkg/s3"
	"github.com/legocy-co/legocy/pkg/s3/proto"
	"github.com/legocy-co/legocy/pkg/s3/proto/mapper"
	log "github.com/sirupsen/logrus"
	"os"
)

func (s ImageStorage) UploadImage(file *os.File, filename, bucketName string) (string, error) {

	conn, err := s.getConnection()
	if err != nil {
		log.Fatalf("did not connect %v", err)
		return "", s3.ErrConnectionRefused
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
