package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "book"
	password = "root"
)

type books struct {
	id   int
	name string
	authon_name string
	page int
}


func main() {

	con := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)

	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows,err := db.Query("select a.id, b.name, a.authon_name, b.page from author as a JOIN books as b ON  a.bookid = b.id")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	kitoblar := []books{}
	kitob:=books{}

	for rows.Next() {
		err = rows.Scan(&kitob.id, &kitob.name, &kitob.authon_name, &kitob.page)
		if err != nil {
			panic(err)
		}
		kitoblar = append(kitoblar, kitob)
	}

	for _,v := range kitoblar {
		fmt.Println(v)
	}
	

	

}
