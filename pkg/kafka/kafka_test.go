package kafka

import (
	"context"
	k "github.com/segmentio/kafka-go"
	"testing"
	"time"
)

// TestKafkaConnection calls new Apache Kafka Producer
func TestKafkaConnection(t *testing.T) {

	uri := "localhost:29092"
	_, cf := context.WithTimeout(context.Background(), time.Second*3)

	defer cf()

	producer, err := newKafkaProducer(HEALTHCHECK_TOPIC, uri)
	if err != nil {
		t.Error(err)
	}

	defer producer.Close()

	_, err = producer.WriteMessages(
		k.Message{Value: []byte("OK")},
	)
	if err != nil {
		t.Error(err)
	}
}
