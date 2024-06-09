package handler

import (
	"mymod/storage/postgres"
	"net/http"
)

type Handler struct {
	
	PostUser packages.RepoNewUser
}

type Handler2 struct {
	
	PostProduct packages.RepoNewProducts
}



func NewHandler(RepoUser packages.RepoNewUser, RepoProduct packages.RepoNewProducts) *http.Server {

	handler := Handler{PostUser: RepoUser}
	handler2 := Handler2{PostProduct: RepoProduct}

	mux := http.NewServeMux()


	mux.HandleFunc("GET /getuser/", handler.GetUser)
	mux.HandleFunc("GET /getuserbyid/", handler.GetUserByID)
	mux.HandleFunc("POST /postuser/", handler.CreateUser)
	mux.HandleFunc("PUT /updateuser/", handler.UpdateUser)
	mux.HandleFunc("DELETE /deleteuser/", handler.DeleteUser)

	mux.HandleFunc("GET /getproduct/", handler2.GetProduct)
	mux.HandleFunc("GET /getproductbyid/", handler2.GetProductByID)
	mux.HandleFunc("POST /postproduct/", handler2.CreateProduct)
	mux.HandleFunc("PUT /updateproduct/", handler2.UpdateProduct)
	mux.HandleFunc("DELETE /deleteproduct/", handler2.DeleteProduct)

	return &http.Server{Addr: ":8080" ,Handler: mux}

}