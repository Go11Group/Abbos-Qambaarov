package main

import (

	"database/sql"
	"fmt"
	"math/rand"

	"github.com/go-faker/faker/v3"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/master?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 1000000; i++ {
		_, err = db.Exec("INSERT INTO product (name, category, cost) VALUES ($1, $2, $3)",
			faker.FirstName(), faker.LastName(), rand.Intn(10000))
		if err != nil {
			fmt.Println(err)
		}
		if i%10000 == 0 {
			fmt.Println(i)
		}
	}
}
