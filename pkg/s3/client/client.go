package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ImageStorage struct {
	host string
	port string
}

func NewImageStorage(host string, port string) ImageStorage {
	return ImageStorage{host: host, port: port}
}

func (s ImageStorage) getConnection() (*grpc.ClientConn, error) {
	return grpc.Dial(s.host+s.port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}
