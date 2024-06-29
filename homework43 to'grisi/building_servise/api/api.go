package api

import (
	"database/sql"
	"mymod/api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(db *sql.DB) *http.Server {
	mux := gin.Default()

	c := handler.NewCardRepo(db)
	s := handler.NewStation(db)
	t := handler.NewTerminal(db)

	mux.POST("cards/create", c.CreateCard)
	mux.GET("cards/getbyid/:id", c.GetCardById)
	mux.GET("cards/get", c.GetCards)
	mux.PUT("cards/update/:id", c.UpdateCard)
	mux.DELETE("cards/delete/:id", c.DeleteCard)
	
	mux.POST("stations/create", s.CreateStation)
	mux.GET("stations/getbyid/:id", s.GetStationById)
	mux.GET("stations/get", s.GetStations)
	mux.PUT("stations/update/:id", s.UpdateStation)
	mux.DELETE("stations/delete/:id", s.DeleteStation)

	mux.POST("terminals/create", t.CreateTerminal)
	mux.GET("terminals/getbyid/:id", t.GetTerminalById)
	mux.GET("terminals/get", t.GetTerminals)
	mux.PUT("terminals/update/:id", t.UpdateTerminal)
	mux.DELETE("terminals/delete/:id", t.DeleteTerminal)

	return &http.Server{Addr: ":8080", Handler: mux}
}