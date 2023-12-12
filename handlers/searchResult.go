package handlers

import (
	"net/http"

	"github.com/dasagho/htmx-test/db"
	"github.com/dasagho/htmx-test/service"
	"github.com/dasagho/htmx-test/views"
)

type SearchResultHandler struct {
	SearchResultService *service.SearchResultSerice
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	valor := queryValues.Get("q")
	data, err := s.SearchResultService.GetResults(valor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = views.GetTemplates().ExecuteTemplate(w, "list", *data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewSearchResultHandler() SearchResultHandler {
	resultRepo := &db.ResultRepository{}
	resultService := &service.SearchResultSerice{SearchRepo: resultRepo}
	resultHandler := SearchResultHandler{SearchResultService: resultService}
	return resultHandler
}
