package server

import (
	"mymod/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Servers() *http.Server {
	router := gin.Default()

	h := handlers.NewHandler()

	router.POST("/api/user", h.UserHandler)
	router.GET("/api/user/:id", h.UserHandler)
	router.GET("/api/user", h.UserHandler)
	router.PUT("/api/user/:id", h.UserHandler)
	router.DELETE("/api/user/:id", h.UserHandler)


	return &http.Server{Addr: ":8081", Handler: router}
}