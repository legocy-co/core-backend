package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/legocy-co/legocy/internal/pkg/config"
)

func NewKafkaConsumer(topicNames []string) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": config.GetAppConfig().KafkaConf.URI,
			"group.id":          config.GetAppConfig().KafkaConf.ConsumerGroupId,
			"auto.offset.reset": "earliest",
		},
	)

	if err != nil {
		return nil, err
	}

	err = consumer.SubscribeTopics(topicNames, nil)
	return consumer, err
}
