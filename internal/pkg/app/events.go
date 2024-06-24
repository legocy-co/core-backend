package app

import (
	"github.com/legocy-co/legocy/internal/pkg/events"
	"github.com/legocy-co/legocy/internal/pkg/kafka"
)

func (a *App) GetEventsDispatcher() events.Dispatcher {
	return kafka.NewDispatcher(a.GetLogger())
}
