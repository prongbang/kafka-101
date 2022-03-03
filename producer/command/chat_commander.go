package command

import (
	"fmt"
	"github.com/prongbang/kafka/event"
	"github.com/prongbang/kafka/producer/producer"
)

type ChatCommander interface {
	Post(cmd PostCommand) error
	Replay(cmd ReplyCommand) error
}

type chatCommander struct {
	eventProducer producer.EventProducer
}

func (c *chatCommander) Post(cmd PostCommand) error {
	e := event.PostEvent{
		Message: cmd.Message,
	}
	fmt.Printf("Event:%#v", e)
	return c.eventProducer.Produce(e)
}

func (c *chatCommander) Replay(cmd ReplyCommand) error {
	e := event.ReplyEvent{
		Message: cmd.Message,
	}
	fmt.Printf("Event:%#v", e)
	return c.eventProducer.Produce(e)
}

func NewChatCommander(eventProducer producer.EventProducer) ChatCommander {
	return &chatCommander{
		eventProducer,
	}
}
