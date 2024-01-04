package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

func IsKafkaConnected(ctx context.Context) bool {
	producer, err := NewKafkaProducer()
	if err != nil {
		logrus.Error("Error establishing Kafka Connection")
		return false
	}

	defer producer.Close()

	logrus.Info("Checking Kafka Connection...")

	t := HEALTHCHECK_TOPIC

	err = producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &t,
				Partition: kafka.PartitionAny,
			},
		}, nil)

	return err == nil
}
