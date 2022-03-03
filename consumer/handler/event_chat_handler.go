package handler

import (
	"encoding/json"
	"fmt"
	"github.com/prongbang/kafka/event"
	"reflect"
)

type eventChatHandler struct {
	// Uc UseCase
}

func (p *eventChatHandler) Handle(topic string, data []byte) {
	switch topic {
	case reflect.TypeOf(event.PostEvent{}).Name():
		value := &event.PostEvent{}
		err := json.Unmarshal(data, value)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Post:", value.Message)
		break
	case reflect.TypeOf(event.ReplyEvent{}).Name():
		value := &event.ReplyEvent{}
		err := json.Unmarshal(data, value)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Reply:", value.Message)
		break
	}
}

func NewEventChatHandler() EventHandler {
	return &eventChatHandler{}
}
