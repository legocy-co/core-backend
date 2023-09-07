package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"legocy-go/internal/config"
)

func NewKafkaProducer(topicName string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config.AppConfigInstance.KafkaConf.URI},
		Topic:   topicName,
	})
}

func ProduceJSONEvent(topicName string, data any) error {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return ErrUnjsonableData
	}

	kafkaProducer := NewKafkaProducer(topicName)

	logrus.Debug("Sending Kafka Message...")

	return kafkaProducer.WriteMessages(
		context.Background(),
		kafka.Message{Value: dataJson, Partition: 0})
}
