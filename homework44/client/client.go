
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	for {
		client, err := rpc.Dial("tcp", "localhost:8080")
		if err != nil {
			log.Fatal("Error client Dial:", err)
		}
		
		fmt.Println("So'zni kiriting: ")
		var translate string
		var word string
		fmt.Scan(&word)
		
		err = client.Call("Translate.Tarjimon", word, &translate)
		if err != nil {
			log.Fatal("Call error:", err)
		}

		fmt.Println(word,":  Tarjimasi-->:", translate)
	}
}