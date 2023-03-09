package server

import (
	"go-cloud/go_01/11api/storage"

	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(lisenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: lisenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	r := gin.Default()
	r.GET("/user/:id", s.handleGetUserByID)
	return r.Run(s.listenAddr)
}
