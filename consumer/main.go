package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/prongbang/kafka/consumer/configuration"
	"github.com/prongbang/kafka/consumer/handler"
	"github.com/prongbang/kafka/event"
	"github.com/spf13/viper"
)

func main() {
	_ = configuration.Load()

	servers := viper.GetStringSlice("kafka.servers")
	groups := viper.GetString("kafka.groups")

	// New Consumer Group
	consumerGroup, err := sarama.NewConsumerGroup(servers, groups, nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = consumerGroup.Close() }()

	eventChatHandler := handler.NewEventChatHandler()
	consumerHandler := handler.NewConsumerHandler(eventChatHandler)
	fmt.Println("Chat consumer listened...")
	for {
		err = consumerGroup.Consume(context.Background(), event.Topics, consumerHandler)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
