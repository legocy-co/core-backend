package kafka

type Dispatcher struct{}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) ProduceJSONEvent(topicName string, data any) error {
	return ProduceJSONEvent(topicName, data)
}
