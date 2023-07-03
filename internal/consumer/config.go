package consumer

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func Kafka() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	return c
}
