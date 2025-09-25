package main

import (
	"myapp/kafka"
	"myapp/router"
	"myapp/router/middleware"

	"github.com/elastic/go-elasticsearch/v8"
)

func init() {
	kafka.InitKafka()
}

//docker-compose -f docker-compose-tool.yml -p tool up -d
//docker-compose -f docker-compose.yml up -d

func main() {
	// 初始化 ES client
	esClient, _ := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
	})

	r := router.CreateRouter()
	r.Use(middleware.ESLogger(esClient))
	router.SetupRouter(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
