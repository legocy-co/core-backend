package kafka

import (
	"context"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/legocy-co/legocy/internal/pkg/config"
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

type producerJobResult struct {
	Err      error
	Producer *kafka.Producer
}

func newProducer() producerJobResult {
	producer, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": config.GetAppConfig().KafkaConf.URI,
		},
	)

	return producerJobResult{Err: err, Producer: producer}
}
