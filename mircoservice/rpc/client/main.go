package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 1、建立连接
	client, err := rpc.Dial("tcp", ":9091")
	if err != nil {
		log.Fatal("连接失败:", err)
	}

	// reply := new(string)
	var reply string //要保证它有内存分配
	err = client.Call("HelloService.Hello", "bob", &reply)
	if err != nil {
		log.Fatal("调用失败...")
	}
	fmt.Println(reply)
}
