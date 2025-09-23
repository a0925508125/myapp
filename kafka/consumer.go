package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var r *kafka.Reader

// NewConsumer 建立一個消費 Kafka Reader
func NewConsumer(brokers []string, topic, groupID string) {
	r = kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		//GroupID:     groupID,
		MinBytes:    1,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset, // <- 從最早的訊息開始
	})
}

// 初始化 Kafka Consumer
func ConsumeMessages() {
	if r == nil {
		log.Fatal("kafka reader not initialized")
		return
	}

	go func() {
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Println("Error reading message:", err)
				time.Sleep(1 * time.Second)
				continue
			}
			log.Printf("Received message: key=%s value=%s\n", string(m.Key), string(m.Value))
		}
	}()
}
