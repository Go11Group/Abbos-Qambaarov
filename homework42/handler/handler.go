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

// postgresdagi lesson repositoryni olish
type Handler2 struct {
	PostLesson packages.RepoNewLesson
}

// postgresdagi enrollmentni repositoryni olish
type Handler3 struct {
	PostEnrollment packages.RepoNewEnrollment
}
func NewHandler(RepoUser packages.RepoNewUser, RepoCourse packages.RepoNewCourse, RepoLesson packages.RepoNewLesson, RepoEnrollment packages.RepoNewEnrollment) *http.Server {

	// posgresdagi repositorylarni va funksiyalarni olish
	handler := Handler{PostUser: RepoUser}
	handler1 := Handler1{PostCourse: RepoCourse}
	handler2 := Handler2{PostLesson: RepoLesson}
	handler3 := Handler3{PostEnrollment: RepoEnrollment}


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

	// lessonga taluqli  handlarni chaqirish
	router.POST("/postlesson", handler2.CreateLesson)
	router.GET("/getlessons", handler2.GetAllLesson)
	router.GET("/getlessonbyid/:id", handler2.GetLessonByID)
	router.PUT("/putlesson/:id", handler2.UpdateLesson)
	router.DELETE("/deletelesson/:id", handler2.DeleteLesson)

	// enrollmentga taluqli  handlarni chaqirish
	router.GET("/getenrollments", handler3.GetAllEnrollment)
	router.GET("/getenrollmentbyid/:id", handler3.GetEnrollmentByID)
	router.POST("/postenrollment", handler3.CreateEnrollment)
	router.PUT("/putenrollment/:id", handler3.UpdateEnrollment)
	router.DELETE("/deleteenrollment/:id", handler3.DeleteEnrollment)

	// Qo'shimcha API lar
	router.GET("/users/:user_id/courses", handler.GetUserCourses)
	router.GET("/courses/:course_id/lessons", handler1.CourseLessons)
	router.GET("/courses/:course_id/enrollments", handler1.CourseUsers)
	router.GET("/users/search", handler.SearchUsers)
	

	router.Run(":8080")
	return &http.Server{Addr: ":8080", Handler: router}


}