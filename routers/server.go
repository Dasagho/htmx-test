package routers

import (
	"net/http"

	handlers "github.com/dasagho/htmx-test/handlers/pages"
	"github.com/dasagho/htmx-test/middleware"
	"github.com/dasagho/htmx-test/views"
)

func NewServer() http.Handler {
	// Parse html templates
	views.InitializeTemplate()

	// Init main Mux
	mainMux := http.NewServeMux()
	mainMux.HandleFunc("/", handlers.IndexHandler)
	mainMux.HandleFunc("/sketch", handlers.SketchHandler)

	// Init Api Mux
	apiMux := NewApiMux()

	// Init static Mux
	fs := http.FileServer(http.Dir("static"))

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))
	mainMux.Handle("/static/", http.StripPrefix("/static/", fs))

	protectedMux := middleware.SessionMiddleware(mainMux)
	return protectedMux
}
