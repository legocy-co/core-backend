package mapper

import (
	"github.com/legocy-co/legocy/pkg/eventNotifier/models"
	"github.com/legocy-co/legocy/pkg/eventNotifier/proto"
	"strconv"
)

func GetNotifyEventRequest(data models.NotifyEventData) *proto.NotifyEventRequest {
	return &proto.NotifyEventRequest{
		ChatID:  strconv.Itoa(data.ChatID),
		Message: data.Message,
	}
}
