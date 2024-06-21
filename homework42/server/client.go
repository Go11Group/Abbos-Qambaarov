package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}

	reqs, err := http.NewRequest("GET", "http://localhost:8082/users", nil)
	if err != nil {
		fmt.Printf("Error creating NewRequest %s\n", err)
		return
	}

	respon, err := client.Do(reqs)
	if err != nil {
		fmt.Println("Error making client:", err)
		return
	}
	defer respon.Body.Close()

	body, err := io.ReadAll(respon.Body)
	if err != nil {
		fmt.Println("Error reading server:", err)
		return
	}

	fmt.Println("Response from server 8082:", string(body))
}