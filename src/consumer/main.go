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

	pc, err := c.ConsumePartition("user-test", 0, sarama.OffsetOldest)
	handlErr(err)

	for true {
		select {
		case msg := <-pc.Messages():
			log.Printf("key: [%v], value: [%v], offset: [%v]", msg.Key, msg.Value, msg.Offset)
		}
	}

}

func handlErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}
