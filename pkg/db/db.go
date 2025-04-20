package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Client *sql.DB
}

func NewDBClient() *DBConfig {
	db, err := sql.Open("postgres", "postgres://local_user:pwddb123@localhost:5432/tabnews?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &DBConfig{
		Client: db,
	}
}

func (c *DBConfig) Ping() error {

	err := c.Client.Ping()
	if err != nil {
		return err
	}

	log.Println("Database connection ok!")

	return nil
}
