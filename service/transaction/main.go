package main

import (
	"context"
	"fmt"
	"log"
	"myapp/proto/pb"
	"net"

	"google.golang.org/grpc"
)

// --- 實作 proto 裡的 Greeter service ---

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello 實作 RPC 方法
func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("收到請求: %s", req.Name)
	return &pb.HelloReply{Message: "Hello, " + req.Name + "!"}, nil
}

// --- 主程式啟動 gRPC Server ---
func main() {
	// 監聽 50051 port
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("無法監聽 port: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})

	fmt.Println("🚀 gRPC Server 已啟動於 port 8090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("無法啟動 server: %v", err)
	}
}
