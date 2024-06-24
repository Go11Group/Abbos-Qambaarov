package api

import (
	"mymod/api/handler"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes() *http.Server {
	mux := gin.Default()

	h := handler.NewHandler()

	mux.GET("/users/get", h.Client)
	mux.PUT("/users/update/:id", h.Client)


	return &http.Server{Handler: mux, Addr: ":8081"}
}