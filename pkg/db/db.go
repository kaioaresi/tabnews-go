package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Ping() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://local_user:pwddb123@localhost:5432/tabnews?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	
	return db, err
}
