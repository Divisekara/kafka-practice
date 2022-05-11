package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	fmt.Println("kafka producer")
	brokers := []string{"localhost:9092"} // we can give multiple brokers here in an array
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	prod, err := sarama.NewSyncProducer(brokers, cfg)
	handleErr(err)

	msg := sarama.ProducerMessage{
		Topic: "user-test",
		Key:   sarama.StringEncoder("2222"),
		Value: sarama.StringEncoder("foo-user"),
		// Headers:   nil,
		// Metadata:  nil,
		// Offset:    0,
		// Partition: 0,
		// Timestamp: time.Time{},
	}

	part, offset, err := prod.SendMessage(&msg)
	handleErr(err)
	fmt.Printf("partition: [%v], offset: [%v] \n", part, offset)
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}
