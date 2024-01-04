package kafka

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/legocy-co/legocy/config"
	log "github.com/sirupsen/logrus"
)

func NewKafkaProducer() (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.GetAppConfig().KafkaConf.URI,
	})

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Produced event to topic %s: key = %-10s value = %-100s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	return producer, err
}

func ProduceJSONEvent(topicName string, data any) error {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return ErrUnjsonableData
	}

	kafkaProducer, err := NewKafkaProducer()
	if err != nil {
		return err
	}

	err = kafkaProducer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topicName,
				Partition: kafka.PartitionAny,
			},
			Value: dataJson,
		},
		nil,
	)

	kafkaProducer.Flush(1000)
	kafkaProducer.Close()
	return err
}
