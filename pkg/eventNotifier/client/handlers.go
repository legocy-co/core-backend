package client

import (
	"context"
	"legocy-go/pkg/eventNotifier/models"
	"legocy-go/pkg/eventNotifier/proto"
	"legocy-go/pkg/eventNotifier/proto/mapper"
)

func (client EventNotifierClient) NotifyEvent(eventData models.NotifyEventData) error {
	conn, err := client.getConnection()
	if err != nil {
		return err
	}

	defer conn.Close()

	grpcClient := proto.NewNotifyEventServiceClient(conn)
	grpcRequest := mapper.GetNotifyEventRequest(eventData)
	_, err = grpcClient.NotifyEvent(context.Background(), grpcRequest)
	return err
}
