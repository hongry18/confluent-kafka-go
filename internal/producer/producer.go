package producer

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func Producer() {
	p := Kafka()
	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivered failed: %v\n", ev.TopicPartition)
					continue
				}

				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}()

	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		_ = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	topic2 := "topic-1"
	for i := 1; i < 21; i++ {
		_ = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic2, Partition: kafka.PartitionAny},
			Value:          []byte(fmt.Sprintf("data %d on topic-1", i)),
		}, nil)
	}

	p.Flush(15 * 1000)
}
