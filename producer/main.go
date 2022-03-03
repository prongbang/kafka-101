package main

import (
	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/kafka/producer/command"
	"github.com/prongbang/kafka/producer/configuration"
	"github.com/prongbang/kafka/producer/handler"
	producer2 "github.com/prongbang/kafka/producer/producer"
	"github.com/spf13/viper"
	"log"
)

func main() {
	_ = configuration.Load()

	servers := viper.GetStringSlice("kafka.servers")
	producer, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = producer.Close() }()

	app := fiber.New()

	eventProducer := producer2.NewEventProducer(producer)
	chatCommander := command.NewChatCommander(eventProducer)
	chatHandler := handler.NewChatHandler(chatCommander)

	v1 := app.Group("/v1")
	{
		v1.Post("/post", chatHandler.Post)
		v1.Post("/reply", chatHandler.Reply)
	}

	log.Fatal(app.Listen(":3000"))
}
