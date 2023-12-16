package routers

import (
	"net/http"

	"github.com/dasagho/htmx-test/db"
	handlers "github.com/dasagho/htmx-test/handlers/api"
	"github.com/dasagho/htmx-test/service"
)

func NewApiMux() *http.ServeMux {
	apiMux := http.NewServeMux()
	resultHandler := handlers.NewSearchResultHandler()
	contactData := db.NewMyContactData()
	contactService := service.NewMyContactService(contactData)
	contactHandler := handlers.NewContactHandler(contactService)

	apiMux.Handle("/contact", contactHandler)
	apiMux.Handle("/list-input", resultHandler)

	return apiMux
}
