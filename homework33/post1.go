package main

import (
	"fmt"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/Abbos", qara)

	err := http.ListenAndServe(":8088",mux)
	if err != nil {
		fmt.Println(err)
	}

}

func qara(w http.ResponseWriter, r *http.Request)  {

	if r.Method != http.MethodPost {
		fmt.Println("Error methon Post")
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
        return
	}

	fmt.Println("URL:", r.URL)
    fmt.Println("Host:", r.Host)
    fmt.Println("Method:", r.Method)

	_, err = w.Write([]byte("working this program"))
	if err != nil{
		fmt.Println("Error writing respons: ",err)
	}


}