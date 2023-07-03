package app

import (
	"legocy-go/internal/config"
	"legocy-go/pkg/eventNotifier/client"
)

func (a *App) GetNotifyEventClient() client.EventNotifierClient {
	return client.NewEventNotifierClient(config.GetAppConfig().EventNotifierPort)
}
