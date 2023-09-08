package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"legocy-go/internal/config"
	"time"
)

func NewKafkaProducer(topicName string) (kafka.Writer, error) {
	_, cf := context.WithTimeout(context.Background(), time.Second*2)
	defer cf()
	return newKafkaProducer(topicName, config.AppConfigInstance.KafkaConf.URI), nil
}

func newKafkaProducer(topicName, uri string) kafka.Writer {
	return kafka.Writer{
		Addr:     kafka.TCP(uri),
		Topic:    topicName,
		Balancer: &kafka.LeastBytes{},
	}
}

func ProduceJSONEvent(topicName string, data any) error {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return ErrUnjsonableData
	}

	kafkaProducer, err := NewKafkaProducer(topicName)
	if err != nil {
		return err
	}

	logrus.Info("Sending Kafka Message...")

	return kafkaProducer.WriteMessages(
		context.Background(),
		kafka.Message{Value: dataJson})
}
