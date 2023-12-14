package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	logging "github.com/dasagho/htmx-test/log"
	"github.com/dasagho/htmx-test/service"
)

type ContactHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type myContactHandler struct {
	service service.ContactService
}

// ServeHTTP implements ContactHandler.
func (h *myContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetContact(h, w, r)
	case "POST":
		handlePostContact(h, w, r)
	case "PUT":
		handlePutContact(h, w, r)
	case "DELETE":
		handleDeleteContact(h, w, r)
	default:
		http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
	}
}

func handleGetContact(h *myContactHandler, w http.ResponseWriter, r *http.Request) {
	// Lógica para manejar GET
	fmt.Fprintf(w, "GET request")
}

func handlePostContact(h *myContactHandler, w http.ResponseWriter, r *http.Request) {
	// Lógica para manejar POST
	fmt.Fprintf(w, "POST request")
}

func handlePutContact(h *myContactHandler, w http.ResponseWriter, r *http.Request) {
	// Lógica para manejar PUT
	fmt.Fprintf(w, "PUT request")
}

func handleDeleteContact(h *myContactHandler, w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	valor := queryValues.Get("id")
	id, err := strconv.Atoi(valor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logging.Debug(fmt.Sprintf("Try delete contact %d", id))

	err = h.service.DeleteContact(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logging.Debug("Contact deleted succesfully")
}

func NewContactHandler(service service.ContactService) ContactHandler {
	return &myContactHandler{
		service: service,
	}
}
