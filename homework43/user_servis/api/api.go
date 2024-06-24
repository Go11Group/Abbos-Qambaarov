package api

import (
	"database/sql"
	"mymod/api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(db *sql.DB) *http.Server {
	mux := gin.Default()

	h := handler.NewHandler(db)

	mux.POST("users/create", h.CreateUser)
	mux.GET("users/getbyid/:id", h.GetUserById)
	mux.GET("users/get", h.GetUsers)
	mux.PUT("users/update/:id", h.UpdateUser)
	mux.DELETE("users/delete/:id", h.DeleteUser)

	

	return &http.Server{Addr: ":8080", Handler: mux}
}
