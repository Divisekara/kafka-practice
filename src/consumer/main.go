package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {

	brokers := []string{"localhost:9092"}

	cfg := sarama.NewConfig()
	c, err := sarama.NewConsumer(brokers, cfg)
	handlErr(err)

	pc, err := c.ConsumePartition("user-test", 0,sarama.offsetOldest)

}

func handlErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}
