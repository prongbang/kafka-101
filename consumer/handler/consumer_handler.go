package handler

import "github.com/Shopify/sarama"

type consumerHandler struct {
	eventChatHandler EventHandler
}

func (c *consumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// Handle
		c.eventChatHandler.Handle(msg.Topic, msg.Value)

		// Mark message to received
		session.MarkMessage(msg, "")
	}
	return nil
}

func NewConsumerHandler(eventChatHandler EventHandler) sarama.ConsumerGroupHandler {
	return &consumerHandler{
		eventChatHandler,
	}
}
