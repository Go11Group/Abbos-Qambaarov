package handler

import (
	packages "mymod/storage/postgres"
	"net/http"

	"github.com/gorilla/mux"
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

	m := mux.NewRouter()

	m.HandleFunc("/getuser/", handler.GetUser).Methods("GET")
	m.HandleFunc("/getuserbyid/", handler.GetUserByID).Methods("GET")
	m.HandleFunc("/postuser/", handler.CreateUser).Methods("POST")
	m.HandleFunc("/updateuser/", handler.UpdateUser).Methods("PUT")
	m.HandleFunc("/deleteuser/", handler.DeleteUser).Methods("DELETE")

	m.HandleFunc("/getproduct/", handler2.GetProduct).Methods("GET")
	m.HandleFunc("/getproductbyid/", handler2.GetProductByID).Methods("GET")
	m.HandleFunc("/postproduct/", handler2.CreateProduct).Methods("POST")
	m.HandleFunc("/updateproduct/", handler2.UpdateProduct).Methods("PUT")
	m.HandleFunc("/deleteproduct/", handler2.DeleteProduct).Methods("DELETE")

	m.HandleFunc("/getstudent/", handler3.GetStudent).Methods("GET")
	m.HandleFunc("/getstudentbyid/", handler3.GetStudentByID).Methods("GET")
	m.HandleFunc("/poststudent/", handler3.CreateStudent).Methods("POST")
	m.HandleFunc("/updatestudent/", handler3.UpdateStudent).Methods("PUT")
	m.HandleFunc("/deletestudent/", handler3.DeleteStudent).Methods("DELETE")

	return &http.Server{Addr: ":8080", Handler: m}

}
