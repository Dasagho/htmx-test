package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/dasagho/htmx-test/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pathViews := filepath.Join("views", "index", "*.html")
	tmpl, err := template.New("base").ParseGlob(pathViews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pathComponents := filepath.Join("views", "index", "components", "*.html")
	tmpla, err := tmpl.ParseGlob(pathComponents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := models.Body{
		MouseDiv: "Body",
		Search:   "Search",
		List:     models.Result{List: []string{}},
	}

	tmpla.ExecuteTemplate(w, "index.html", data)
}
