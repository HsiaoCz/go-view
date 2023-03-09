package main

import (
	"flag"
	"fmt"
	"go-cloud/go_01/11api/server"
	"go-cloud/go_01/11api/storage"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", "9091", "set http server listenAddr")
	flag.Parse()

	fmt.Println("the server is running on the port", *listenAddr)

	store := storage.NewMysqlStorage()

	srv := server.NewServer(*listenAddr, store)

	log.Fatal(srv.Start())
}
