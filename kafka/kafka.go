package kafka

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

var brokers = []string{"kafka:29092"} // 可以加多個 broker
//var brokers = []string{"localhost:9092"} // 可以加多個 broker

const (
	Topic   = "hello-tasks"
	GroupID = "hello-worker"
)

func InitKafka() {
	// 建立 Topic
	for _, broker := range brokers {
		if err := createTopic(broker, Topic, 1, 1); err != nil {
			log.Fatalf("%s, failed to create Topic: %v", broker, err)
		}
	}

	NewConsumer(brokers, Topic, GroupID)
	NewKafkaWriter(brokers, Topic)

	ConsumeMessages()
}

// createTopic 建立 Topic
func createTopic(brokerAddress, topic string, partitions, replicationFactor int) error {
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	// 連到 controller 才能建立 Topic
	var controllerConn *kafka.Conn
	controllerAddr := fmt.Sprintf("%s:%d", controller.Host, controller.Port)
	controllerConn, err = kafka.Dial("tcp", controllerAddr)
	if err != nil {
		return fmt.Errorf("failed to dial controller: %w", err)
	}
	defer controllerConn.Close()

	// 建立 Topic
	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     partitions,
			ReplicationFactor: replicationFactor,
		},
	}

	return controllerConn.CreateTopics(topicConfigs...)
}
