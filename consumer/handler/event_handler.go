package handler

type EventHandler interface {
	Handle(topic string, data []byte)
}
