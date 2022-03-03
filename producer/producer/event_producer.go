package producer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/prongbang/kafka/event"
	"reflect"
)

type EventProducer interface {
	Produce(e event.Event) error
}

type eventProducer struct {
	producer sarama.SyncProducer
}

func (e2 *eventProducer) Produce(e event.Event) error {
	value, err := json.Marshal(e)
	if err != nil {
		return err
	}

	topic := reflect.TypeOf(e).Name()

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := e2.producer.SendMessage(&msg)
	if err != nil {
		return err
	}
	fmt.Printf("Partition=%v, Offset=%v", partition, offset)
	return nil
}

func NewEventProducer(producer sarama.SyncProducer) EventProducer {
	return &eventProducer{
		producer,
	}
}
