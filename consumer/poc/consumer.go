package poc

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func Consume() {
	servers := viper.GetStringSlice("kafka.servers")
	groups := viper.GetString("kafka.groups")

	fmt.Println(servers)

	// With Group
	consumerGroup, err := sarama.NewConsumerGroup(servers, groups, nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = consumerGroup.Close() }()

	// Without Group
	consumer, err := sarama.NewConsumer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = consumer.Close() }()

	topic := "test"

	// $ kafka-topics --bootstrap-server=192.168.10.107:9092 --topic=test --describe
	//  Topic: test     TopicId: xxx    PartitionCount: 1    ReplicationFactor: 1    Configs: segment.bytes=1073741824
	//  Topic: test     Partition: 0    Leader: 1003    	 Replicas: 1003  		 Isr: 1003
	partition := int32(0)
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer func() { _ = partitionConsumer.Close() }()

	fmt.Println("Consumer Listening...")
	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Println(err)
		case msg := <-partitionConsumer.Messages():
			fmt.Println("Message:", string(msg.Value))
		}
	}
}
