package kafka

import (
	"context"
	"fmt"

	kafkago "github.com/segmentio/kafka-go"
)

func StartUserCreatedConsumer() {
	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "user-created",
		GroupID: "user-created-consumer-group",
	})

	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Kafka consumer error:", err)
			continue
		}

		fmt.Println("Received UserCreated event:")
		fmt.Println(string(message.Value))
	}
}
