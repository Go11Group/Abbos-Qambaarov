package handler

import (
	packages "mymod/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	PostUser packages.RepoNewUser
}

type Handler2 struct {
	PostProduct packages.RepoNewProducts
}

type Handler3 struct {
	PostStudent packages.RepoNewStudent
}

func NewHandler(RepoUser packages.RepoNewUser, RepoProduct packages.RepoNewProducts, RepoStudent packages.RepoNewStudent) *http.Server {

	handler := Handler{PostUser: RepoUser}
	handler2 := Handler2{PostProduct: RepoProduct}
	handler3 := Handler3{PostStudent: RepoStudent}


	router := gin.Default()
	router.GET("/getuser/", handler.GetUser)
	router.GET("/getuserbyid/:id", handler.GetUserByID)
	router.POST("/postuser", handler.CreateUser)
	router.PUT("/putuser/:id", handler.UpdateUser)
	router.DELETE("/deleteuser/:id", handler.DeleteUser)	

	router.GET("/getproduct/", handler2.GetProduct)
	router.GET("/getproductbyid/:id", handler2.GetProductByID)
	router.POST("/postproduct", handler2.CreateProduct)
	router.PUT("/putproduct/:id", handler2.UpdateProduct)
	router.DELETE("/deleteproduct/:id", handler2.DeleteProduct)

    router.GET("/getstudent/", handler3.GetStudent)
	router.GET("/getstudentbyid/:id", handler3.GetStudentByID)
	router.POST("poststudent", handler3.CreateStudent)
	router.PUT("/putstudent/:id", handler3.UpdateStudent)
	router.DELETE("/deletestudent/:id", handler3.DeleteStudent)

	router.Run(":8080")
	return &http.Server{Addr: ":8080", Handler: router}


}
