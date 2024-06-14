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
	RepoCourse := packages.RepoNewCourse{Db: db}
	RepoLesson := packages.RepoNewLesson{Db: db}
	RepoEnrollment := packages.RepoNewEnrollment{Db: db}	
	server := handler.NewHandler(RepoUser, RepoCourse, RepoLesson, RepoEnrollment)

	server.ListenAndServe()

}