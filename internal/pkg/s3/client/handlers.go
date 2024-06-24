package client

import (
	"context"
	proto2 "github.com/legocy-co/legocy/internal/pkg/s3/proto"

	"google.golang.org/grpc"
)

// UploadImageFromFile uploads an image from a file
// to the S3 bucket
// It returns the Slug of the uploaded image
func (s ImageStorage) UploadImageFromFile(ctx context.Context, in *proto2.UploadImageFileRequest, opts ...grpc.CallOption) (string, error) {
	conn, err := s.getConnection()
	if err != nil {
		return "", err
	}

	defer conn.Close()

	client := proto2.NewS3ServiceClient(conn)

	response, err := client.UploadImageFromFile(ctx, in, opts...)
	if err != nil {
		return "", err
	}

	return response.GetImageURL(), nil
}

// UploadImageFromURL uploads an image from a URL
// to the S3 bucket
// It returns the Slug of the uploaded image
func (s ImageStorage) UploadImageFromURL(ctx context.Context, in *proto2.UploadImageURLRequest, opts ...grpc.CallOption) (string, error) {
	conn, err := s.getConnection()
	if err != nil {
		return "", err
	}

	defer conn.Close()

	client := proto2.NewS3ServiceClient(conn)

	response, err := client.UploadImageFromURL(ctx, in, opts...)
	if err != nil {
		return "", err
	}

	return response.GetImageURL(), nil
}
