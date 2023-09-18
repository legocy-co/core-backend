package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func IsKafkaConnected(ctx context.Context) bool {
	producer, err := NewKafkaProducer(HEALTHCHECK_TOPIC)
	if err != nil {
		logrus.Error("Error establishing Kafka Connection")
		return false
	}

	defer producer.Close()

	logrus.Info("Checking Kafka Connection...")

	_, err = producer.WriteMessages(
		kafka.Message{Value: []byte("OK")})

	return err == nil
}
