package main

import (
	"net/http"
	"os"

	"github.com/dasagho/htmx-test/db"
	logging "github.com/dasagho/htmx-test/log"
	"github.com/dasagho/htmx-test/routers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Puerto por defecto
	}

	logging.CreateLogs()
	db.ConnectDB()
	mux := routers.NewServer()
	http.ListenAndServe(":"+port, mux)
}
