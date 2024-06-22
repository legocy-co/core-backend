package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/legocy-co/legocy/internal/pkg/config"
	log "github.com/sirupsen/logrus"
	"time"
)

func NewKafkaProducer() (*kafka.Producer, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	out := make(chan producerJobResult)

	go func() {
		out <- newProducer()
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("kafka producer creation timeout")
	case producer := <-out:
		return producer.Producer, producer.Err
	}
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

type producerJobResult struct {
	Err      error
	Producer *kafka.Producer
}

func newProducer() producerJobResult {
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

	return producerJobResult{Err: err, Producer: producer}
}
