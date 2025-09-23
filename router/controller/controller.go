package controller

import (
	"log"
	"myapp/kafka"
	"myapp/router/base_controller"
	"net/http"

	"myapp/errcode"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	base_controller.BaseController
	ES *elasticsearch.Client
}

func NewController() *Controller {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"}, // docker-compose service 名稱
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return &Controller{
		ES: es,
	}
}

func (ctrl *Controller) GetHello(c *gin.Context) {
	err := kafka.ProduceMessage("hello task payload")
	if err != nil {
		log.Println("Kafka publish error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to publish task"})
		return
	} else {
		log.Println("Kafka publish success")
	}
	ctrl.JsonResponse(c, errcode.Success, nil)
}
