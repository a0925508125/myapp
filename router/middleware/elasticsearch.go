package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

type HTTPLog struct {
	Path    string `json:"path"`
	Method  string `json:"method"`
	Status  int    `json:"status"`
	Latency int64  `json:"latency_ms"`
	UserID  int64  `json:"user_id,omitempty"`
	Ts      int64  `json:"ts"`
}

func ESLogger(esClient *es.Client) gin.HandlerFunc {
	log.Printf("ESLogger")
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start).Milliseconds()
		status := c.Writer.Status()

		var userID int64 = 0
		if v, exists := c.Get("user_id"); exists {
			if id, ok := v.(int64); ok {
				userID = id
			}
		}

		log := map[string]interface{}{
			"path":       c.FullPath(),
			"method":     c.Request.Method,
			"status":     status,
			"latency":    latency,
			"user_id":    userID,
			"ts":         time.Now().Unix(),
			"body":       c.Request.Body,
			"@timestamp": time.Now().UTC().Format(time.RFC3339),
		}
		b, _ := json.Marshal(log)
		// fire-and-forget
		go esClient.Index("http-logs", bytes.NewReader(b))

		// 寫入 Elasticsearch
		go func() {
			res, err := esClient.Index(
				"http-logs",
				bytes.NewReader(b),
				esClient.Index.WithRefresh("true"),
			)
			if err != nil {
				fmt.Println("ES write error:", err)
				return
			}
			defer res.Body.Close()
			if res.IsError() {
				fmt.Println("ES response error:", res.String())
			} else {
				fmt.Println("ES write success:", string(b))
			}
		}()
	}
}
