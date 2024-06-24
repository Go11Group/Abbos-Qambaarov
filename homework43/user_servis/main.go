package main

import (
	"mymod/api"
	"mymod/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDb()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	server := api.Routes(db)
	err = server.ListenAndServe() // 80
	if err != nil {
		panic(err.Error())
	}
}
