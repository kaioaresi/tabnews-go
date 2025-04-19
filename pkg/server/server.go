package server

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Home Page")
}
