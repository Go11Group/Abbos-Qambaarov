package handler

import (
	"mymod/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

// postgresdagi user repositoryni olish
type Handler struct {
	PostUser packages.RepoNewUser
}

// postgresdagi course repositoryni olish
type Handler1 struct {
	PostCourse packages.RepoNewCourse
}


func NewHandler(RepoUser packages.RepoNewUser, RepoCourse packages.RepoNewCourse) *http.Server {

	// posgresdagi repositorylarni va funksiyalarni olish
	handler := Handler{PostUser: RepoUser}
	handler1 := Handler1{PostCourse: RepoCourse}


	// userga taluqli  handlarni chaqirish 
	router := gin.Default()
	router.POST("/postuser", handler.CreateUser)
	router.GET("/getusers/", handler.GetUsers)
	router.GET("/getuserbyid/:id", handler.GetUserByID)
	router.PUT("/putuser/:id", handler.UpdateUser)
	router.DELETE("/deleteuser/:id", handler.DeleteUser)	

	// coursega taluqli  handlarni chaqirish 
	router.POST("/postcourse", handler1.CreateCourse)
	router.GET("/getcourses", handler1.GetAllCourse)
	router.GET("/getcoursebyid/:id", handler1.GetCourseByID)
	router.PUT("/putcourse/:id", handler1.UpdateCourse)
	router.DELETE("/deletecourse/:id", handler1.DeleteCourse)

	


	router.Run(":8080")
	return &http.Server{Addr: ":8080", Handler: router}


}