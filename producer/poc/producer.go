package poc

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func Producer() {
	server := viper.GetStringSlice("kafka.servers")

	producer, err := sarama.NewSyncProducer(server, nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = producer.Close() }()

	msg := sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("Hello, I'am producer"),
	}

	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Partition=%v, Offset=%v", partition, offset)
}
