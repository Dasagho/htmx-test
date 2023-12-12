package main

import (
	"net/http"
	"os"

	"github.com/dasagho/htmx-test/routers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Puerto por defecto
	}
	mux := routers.NewServer()
	http.ListenAndServe(":"+port, mux)
}
