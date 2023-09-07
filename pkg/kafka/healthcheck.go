package kafka

import (
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

	_, err = producer.WriteMessages(
		kafka.Message{Value: []byte("OK"), Partition: 1})
	return err != nil
}
