package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1、建立连接
	// 使用rpc.Dail()内部进行了编码
	// 直接net.Dail()调用建立了一个连接
	conn, err := net.Dial("tcp", ":9091")
	if err != nil {
		log.Fatal("连接失败:", err)
	}

	// reply := new(string)
	var reply string //要保证它有内存分配
	// 使用json编码包装一下
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "bob", &reply)
	if err != nil {
		log.Fatal("调用失败...")
	}
	fmt.Println(reply)
}
