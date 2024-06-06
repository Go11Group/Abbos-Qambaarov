package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", userHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	//id va name ni urldan kiritamiz 
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")

	fmt.Println("ID: ", id)
	fmt.Println("Name: ", name)

	response := fmt.Sprintf("User ID: %s, \nUser Name: %s", id, name)
	_, err := w.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response", err)
	}
}
