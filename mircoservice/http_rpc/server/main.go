package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	_ = rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	// 2、注册处理逻辑
	// 3、启动服务
	http.ListenAndServe(":9091", nil)
}

// 有一个问题，这里的registerName里面是硬编码的，可能会导致客户端和服务端不一致
// 这里可以将RegisterName放在一个单独的包里面
