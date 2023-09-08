package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func IsKafkaConnected() bool {
	producer, err := NewKafkaProducer(HEALTHCHECK_TOPIC)
	if err != nil {
		logrus.Error("Error establishing Kafka Connection")
		return false
	}

	logrus.Info("Checking Kafka Connection...")

	err = producer.WriteMessages(
		context.Background(),
		kafka.Message{Value: []byte("OK"), Partition: 1})
	return err != nil
}
