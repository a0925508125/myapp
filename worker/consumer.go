package worker

type Order struct {
	UserID int64 `json:"user_id"`
	ItemID int64 `json:"item_id"`
	Count  int   `json:"count"`
	Ts     int64 `json:"ts"`
}

func StartOrderWorker() {
	//mq.StartConsumer([]string{"localhost:9092"}, "orders-topic", func(value []byte) {
	//	var o Order
	//	if err := json.Unmarshal(value, &o); err != nil {
	//		log.Println("Unmarshal error:", err)
	//		return
	//	}
	//	_, err := db.DB.Exec(
	//		"INSERT INTO orders(user_id, item_id, count, created_at) VALUES (?, ?, ?, FROM_UNIXTIME(?))",
	//		o.UserID, o.ItemID, o.Count, o.Ts,
	//	)
	//	if err != nil && err != sql.ErrNoRows {
	//		log.Println("DB insert error:", err)
	//	} else {
	//		log.Println("Order saved:", o)
	//	}
	//})
}
