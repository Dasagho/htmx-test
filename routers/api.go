package routers

import (
	"net/http"

	"github.com/dasagho/htmx-test/handlers"
)

func NewApiMux() *http.ServeMux {
	apiMux := http.NewServeMux()
	resultHandler := handlers.NewSearchResultHandler()

	apiMux.HandleFunc("/mouse_entered", handlers.Mouse)
	apiMux.Handle("/trigger_delay", resultHandler)

	return apiMux
}
