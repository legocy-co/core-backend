package kafka

import (
	"github.com/segmentio/kafka-go"
	"legocy-go/internal/config"
)

func NewKafkaProducer(topicName string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config.AppConfigInstance.KafkaConf.URI},
		Topic:   topicName,
	})
}
