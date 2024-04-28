package di

import (
	"github.com/legocy-co/legocy/internal/delivery/kafka"
	"github.com/legocy-co/legocy/internal/pkg/events"
)

func ProvideDispatcher() events.Dispatcher {
	return kafka.NewDispatcher()
}
