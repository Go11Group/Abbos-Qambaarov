package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ismingizni kiriting:  ")
	var name string
	fmt.Scan(&name)
	for {
		fmt.Print(name, ":  ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		input = name+":  "+input
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		fmt.Print("Server response: ", response)
	}
}
