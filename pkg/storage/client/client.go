package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ImageStorage struct {
	port string
}

func NewImageStorage(port string) ImageStorage {
	return ImageStorage{port: port}
}

func (s ImageStorage) getConnection() (*grpc.ClientConn, error) {
	return grpc.Dial("localhost"+s.port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}
