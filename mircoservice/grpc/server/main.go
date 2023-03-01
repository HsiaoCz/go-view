package main

import (
	"context"
	"fmt"
	"go-cloud/mircoservice/grpc/proto"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Reply: "hello:" + request.Name,
	}, nil
}

func main() {
	// 监听本地的端口
	lis, err := net.Listen("tcp", ":8990")
	if err != nil {
		fmt.Printf("failed to listen:%v\n", err)
		return
	}
	// 创建grpc的服务器
	s := grpc.NewServer()
	// 在grpc服务端注册服务
	proto.RegisterGreeterServer(s, &Server{})
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve:%v", err)
		return
	}
}
