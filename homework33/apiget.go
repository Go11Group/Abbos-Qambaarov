package main

import (
	"fmt"
	"net/http"
)

func main()  {
    http.HandleFunc("POST /salom", salom)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }

}

func salom(w http.ResponseWriter, r *http.Request) {

    fmt.Println("URL: ", r.URL)
    fmt.Println("Host: ",r.Host)
    fmt.Println("Method: ",r.Method)

    _, err := w.Write([]byte("salomlar"))
    if err != nil {
        fmt.Println("Error writing respons",err)
    }

    
}