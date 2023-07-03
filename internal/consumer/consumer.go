package consumer

import "fmt"

func Consumer() {
	c := Kafka()

	c.SubscribeTopics([]string{"topic-1"}, nil)
	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			fmt.Printf("consumer error: %v (%v)\n", err, msg)
			continue
		}

		fmt.Printf("message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	}
}
