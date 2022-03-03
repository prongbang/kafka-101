package event

import "reflect"

var Topics = []string{
	reflect.TypeOf(PostEvent{}).Name(),
	reflect.TypeOf(ReplyEvent{}).Name(),
}

type Event interface {
}

type PostEvent struct {
	Message string `json:"message"`
}

type ReplyEvent struct {
	Message string `json:"message"`
}
