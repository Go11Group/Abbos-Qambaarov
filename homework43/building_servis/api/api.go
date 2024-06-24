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
	c := handler.NewCardRepo(db)

	mux.POST("users/create", h.CreateUser)
	mux.GET("users/getbyid/:id", h.GetUserById)
	mux.GET("users/get", h.GetUsers)
	mux.PUT("users/update/:id", h.UpdateUser)
	mux.DELETE("users/delete/:id", h.DeleteUser)

	mux.POST("cards/create", c.CreateCard)
	mux.GET("cards/getbyid/:id", c.GetCardById)
	mux.GET("cards/get", c.GetCards)
	mux.PUT("cards/update/:id", c.UpdateCard)
	mux.DELETE("cards/delete/:id", c.DeleteCard)
	

	return &http.Server{Addr: ":8080", Handler: mux}
}
