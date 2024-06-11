package main

import (
	"fmt"
	"mymod/handler"
	packages "mymod/storage/postgres"
	
)

type Users struct {
	Id       string
	UserName string
	Email    string
	Password string
}


func main() {

	db, err := packages.ConnectDB()
	if err != nil {
		fmt.Println(err)
        return
	}
	RepoUser := packages.RepoNewUser{Db: db}
	RepoProduct := packages.RepoNewProducts{Db: db}
	RepoStudent := packages.RepoNewStudent{Db: db}
	server := handler.NewHandler(RepoUser, RepoProduct, RepoStudent)

	server.ListenAndServe()

}
