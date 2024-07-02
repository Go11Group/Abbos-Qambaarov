package server

import (
	"mymod/api"
	"github.com/gin-gonic/gin"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Start() error {
	router := gin.Default()

	routes.SetupRoutes(router)

	return router.Run(s.addr)
}