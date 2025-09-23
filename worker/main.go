package main

//func main() {
//	brokers := []string{"kafka:9092"}
//	topic := "hello-tasks"
//	groupID := "hello-worker"
//
//	// 等待 Kafka 啟動
//	kafka.WaitForKafka(brokers)
//
//	// 建立 Consumer
//	r := kafka.NewConsumer(brokers, topic, groupID)
//
//	log.Println("Worker started, waiting for messages...")
//
//	// 循環讀取訊息
//	for {
//		m, err := r.ReadMessage(context.Background())
//		if err != nil {
//			log.Println("Error reading message:", err)
//			time.Sleep(time.Second)
//			continue
//		}
//
//		// 處理訊息
//		log.Printf("Received task: %s\n", string(m.Value))
//		// TODO: 可以放資料庫寫入或其他任務
//	}
//}
