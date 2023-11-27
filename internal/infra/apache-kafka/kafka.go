package apachekafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consumer(topics []string, servers string, msgChannel chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "go-mensageria",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Println("ERROR: failed to initialize Apache Kafka")
		panic(err)
	}

	kafkaConsumer.SubscribeTopics(topics, nil)

	for {
		message, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChannel <- message
		}
	}
}
