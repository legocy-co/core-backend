package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func IsKafkaConnected() bool {
	producer := NewKafkaProducer(HEALTHCHECK_TOPIC)

	logrus.Debug("Checking Kafka Connection...")

	err := producer.WriteMessages(
		context.Background(), kafka.Message{Value: []byte("OK"), Partition: 0})
	return err != nil
}
