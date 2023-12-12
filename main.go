package main

import (
	"net/http"

	"github.com/dasagho/htmx-test/routers"
)

func main() {
	mux := routers.NewServer()
	http.ListenAndServe(":8080", mux)
}
