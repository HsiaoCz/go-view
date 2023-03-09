package main

import (
	"go-cloud/go-kit/api"
	"net/http"

	httpss "github.com/go-kit/kit/transport/http"
)

func main() {
	user := api.UserService{}
	endp := api.GenUserEndpoint(user)

	serverHandler := httpss.NewServer(endp, api.DecodeUserRequest, api.EncodeUserResponse)
	http.ListenAndServe(":9091", serverHandler)
}
