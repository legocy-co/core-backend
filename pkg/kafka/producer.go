package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"legocy-go/internal/config"
	"time"
)

func NewKafkaProducer(topicName string) (*kafka.Conn, error) {
	ctx, cf := context.WithTimeout(context.Background(), time.Second*2)
	defer cf()
	return kafka.DialLeader(
		ctx,
		"tcp",
		config.AppConfigInstance.KafkaConf.URI,
		topicName, 1)
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

	_, err = kafkaProducer.WriteMessages(
		kafka.Message{Value: dataJson})
	return err
}
