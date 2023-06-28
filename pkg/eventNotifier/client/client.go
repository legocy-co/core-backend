package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EventNotifierClient struct {
	port string
}

func NewEventNotifierClient(port string) EventNotifierClient {
	return EventNotifierClient{port: port}
}

func (client EventNotifierClient) getConnection() (*grpc.ClientConn, error) {
	return grpc.Dial("localhost"+client.port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}
