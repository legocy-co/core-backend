package mapper

import (
	"legocy-go/pkg/eventNotifier/models"
	"legocy-go/pkg/eventNotifier/proto"
	"strconv"
)

func GetNotifyEventRequest(data models.NotifyEventData) *proto.NotifyEventRequest {
	return &proto.NotifyEventRequest{
		ChatID:  strconv.Itoa(data.ChatID),
		Message: data.Message,
	}
}
