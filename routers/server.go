package routers

import (
	"net/http"

	"github.com/dasagho/htmx-test/db"
	"github.com/dasagho/htmx-test/handlers"
	"github.com/dasagho/htmx-test/service"
)

func NewServer() *http.ServeMux {
	// db, err := sql.Open("postgres", "your-database-connection-string")
	// if err != nil {
	//     // Manejar error
	// }
	// defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/mouse_entered", handlers.Mouse)

	resultRepo := &db.ResultRepository{}
	resultService := &service.SearchResultSerice{SearchRepo: resultRepo}
	resultHandler := handlers.SearchResultHandler{SearchResultService: resultService}
	apiMux.Handle("/trigger_delay", resultHandler)

	mux.Handle("/api/", http.StripPrefix("/api", apiMux))

	return mux
}
