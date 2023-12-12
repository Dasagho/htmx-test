package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/dasagho/htmx-test/service"
)

type SearchResultHandler struct {
	SearchResultService *service.SearchResultSerice
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathComponents := filepath.Join("views", "index", "components", "list.html")
	tmpl, err := template.ParseFiles(pathComponents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	queryValues := r.URL.Query()
	valor := queryValues.Get("q")
	data, err := s.SearchResultService.GetResults(valor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println("Lista en handler:", *data)
	// fmt.Println(pathComponents)
	// fmt.Println(tmpl.Tree)
	tmpl.ExecuteTemplate(w, "list", *data)
	err = tmpl.Execute(w, *data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
