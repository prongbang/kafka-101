kafka_topic_list:
	kafka-topics --bootstrap-server=192.168.10.107:9092 --list

kafka_create_topic:
	kafka-topics --bootstrap-server=192.168.10.107:9092 --create --topic=test

kafka_create_topic_describe:
	kafka-topics --bootstrap-server=192.168.10.107:9092 --topic=test --describe

kafka_console_consumer:
	kafka-console-consumer --bootstrap-server=192.168.10.107:9092 --topic=test --from-beginning

kafka_console_producer:
	kafka-console-producer --bootstrap-server=192.168.10.107:9092 --topic=test

kafka_console_consumer_group:
	kafka-console-consumer --bootstrap-server=192.168.10.107:9092 --topic=test --group=my-consumer-group --from-beginning

kafka_console_producer_group:
	kafka-console-producer --bootstrap-server=192.168.10.107:9092 --topic=test

# make kafka_console_producer topic=PostEvent
# > {"message": "Hi, Bob"}
# make kafka_console_producer topic=ReplyEvent
# > {"message": "Hi, Alice"}
kafka_console_producer:
	kafka-console-producer --bootstrap-server=192.168.10.107:9092 --topic=$(topic)

# make kafka_consumer_group group=my-consumer-group
kafka_consumer_group:
	kafka-console-consumer --bootstrap-server=192.168.10.107:9092 --include="PostEvent|ReplyEvent" --group=$(group)

kafka_consumer_group_list:
	kafka-consumer-groups --bootstrap-server=192.168.10.107:9092 --list

