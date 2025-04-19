package main

import (
	"fmt"
	"log"
	"tabnews-go/pkg/db"
)

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", server.Home)
	// mux.HandleFunc("/api/v1/status", server.Status)
	// mux.HandleFunc("/api/v1/migrations", server.Migrations)

	// log.Println("Server listing :8080....")
	// log.Fatal(http.ListenAndServe(":8080", mux))

	db,err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
}
