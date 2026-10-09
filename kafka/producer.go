package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

var Writer = &kafka.Writer{
	Addr:  kafka.TCP("localhost:9092"),
	Topic: "user-created",
}

func PublishUserCreated(message []byte) error {
	return Writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: message,
		},
	)
}
