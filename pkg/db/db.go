package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Ping() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://local_user:pwddb123@localhost:5432/tabnews?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connection ok!")

	return db, err
}
