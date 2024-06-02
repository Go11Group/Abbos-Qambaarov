package packages

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connectdb() (*sql.DB, error){
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/education?sslmode=disable")
	if err != nil {
		return db,err
	}

	return db, nil
}
