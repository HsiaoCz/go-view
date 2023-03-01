package main

import (
	"flag"
	"go-cloud/mircoservice/grpc/proto"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpc client

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", ":8990", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// 连接server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("did not connect:", err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// 执行调用并打印响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: *name})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greeting:%s", r.GetReply())
}
