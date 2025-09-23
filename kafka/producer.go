package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer

func NewKafkaWriter(brokers []string, topic string) {
	Writer = &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// ProduceMessage 封裝一個發送訊息的方法
func ProduceMessage(msg string) error {
	if Writer == nil {
		return ErrWriterNotInitialized
	}

	err := Writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(msg),
		},
	)
	if err != nil {
		log.Printf("failed to write message: %v", err)
		return err
	}
	log.Printf("sent message: %s", msg)
	return nil
}

// 自定義錯誤
var ErrWriterNotInitialized = fmt.Errorf("kafka writer not initialized")

//進入容器
//docker exec -it kafka /bin/bash
//查看topic
// /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list
//查看groups
// /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list
//查看broker
// /opt/kafka/bin/kafka-broker-api-versions.sh --bootstrap-server kafka:9092
