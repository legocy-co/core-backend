package kafka

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log/slog"
)

type Dispatcher struct {
	log *slog.Logger
}

func NewDispatcher(log *slog.Logger) *Dispatcher {
	return &Dispatcher{
		log: log,
	}
}

func (d *Dispatcher) ProduceJSONEvent(topicName string, data any) error {

	producer, err := NewKafkaProducer()
	if err != nil {
		return err
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					d.log.Error(
						"Failed to deliver kafka message",
						slog.String("topic", *ev.TopicPartition.Topic),
					)
				} else {
					d.log.Info(
						"Kafka message delivered",
						slog.String("topic", *ev.TopicPartition.Topic),
						slog.Any("message", ev.Value),
					)
				}
			}
		}
	}()

	err = produceJSONEvent(producer, topicName, data)

	producer.Flush(1000)
	producer.Close()

	return err
}

func produceJSONEvent(producer *kafka.Producer, topicName string, data any) error {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return ErrUnjsonableData
	}

	return producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topicName,
				Partition: kafka.PartitionAny,
			},
			Value: dataJson,
		},
		nil,
	)
}
