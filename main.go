package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dasagho/htmx-test/db"
	logging "github.com/dasagho/htmx-test/log"
	"github.com/dasagho/htmx-test/routers"
)

func main() {
	port := "8080"

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	logging.CreateLogs()
	db.ConnectDB()
	server := routers.NewServer()
	log.Printf("Server listening on: http://%s:%s", host, port)
	http.ListenAndServe(":"+port, server)
}
