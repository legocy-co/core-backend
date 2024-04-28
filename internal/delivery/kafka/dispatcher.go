package kafka

import "github.com/legocy-co/legocy/pkg/kafka"

type Dispatcher struct{}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) ProduceJSONEvent(topicName string, data any) error {
	return kafka.ProduceJSONEvent(topicName, data)
}
