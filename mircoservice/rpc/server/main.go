package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是reply
	*reply = "Hello:" + request
	return nil
}

func main() {
	// rpc三步走逻辑
	// 1. 实例化一个server
	listener, err := net.Listen("tcp", "9091")
	if err != nil {
		log.Fatal(err)
	}
	// 2、注册处理逻辑
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3、启动服务
    conn,err:=listener.Accept()  // 当一个新的连接进来的时候，就有一个socket套接字
    if err!=nil {
		log.Fatal(err)
	}
    rpc.ServeConn(conn)
}

// rpc调用过程中需要解决的问题
// call id  序列化和反序列化
// go语言的的rpc的序列化协议是什么(gob) 