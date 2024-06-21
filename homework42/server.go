package main

import (
	"fmt"
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	reqs, err := http.NewRequest(r.Method, "http://localhost:8080"+r.RequestURI, r.Body)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reqs.Header = r.Header

	respon, err := client.Do(reqs)
	if err != nil {
		fmt.Printf("Error client %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer respon.Body.Close()

	body, err := io.ReadAll(respon.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("response from 8080:", string(body))

	for k, v := range respon.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(respon.StatusCode)
	w.Write(body)
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("listening to server 8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println(err)
	}
}
