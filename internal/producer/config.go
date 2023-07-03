package producer

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

func Kafka() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}
	return p
}
