package controller

import (
	"context"
	"log"
	"myapp/kafka"
	"myapp/proto/pb"
	"myapp/router/base_controller"
	"net/http"
	"time"

	"myapp/errcode"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
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

	// 連線到本地 gRPC server
	conn, err := grpc.Dial("transaction:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("連線失敗: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 設定 context 與超時
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 呼叫 SayHello RPC
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "HELLO"})
	if err != nil {
		log.Fatalf("RPC 呼叫失敗: %v", err)
	}

	log.Printf("伺服器回覆: %s", resp.Message)

	ctrl.JsonResponse(c, errcode.Success, nil)
}
