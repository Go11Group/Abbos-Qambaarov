package main

import (
	"log"

	"mymod/server"
)

func main() {
	ser := server.NewServer(":8083")
	err := ser.Servers()
	if err != nil {
		log.Println(err)
	}
}