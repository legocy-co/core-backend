package events

type Dispatcher interface {
	ProduceJSONEvent(topicName string, data any) error
}
