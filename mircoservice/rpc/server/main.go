package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	X int
	Y int
}

type ServiceA struct{}

func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	service := new(ServiceA)
	rpc.Register(service) //注册rpc服务
	rpc.HandleHTTP()      //基于http协议
	l, e := net.Listen("tcp", "9091")
	if e != nil {
		log.Fatal(e)
	}
	http.Serve(l, nil)
}
