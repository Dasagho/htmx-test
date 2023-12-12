package handlers

import (
	"fmt"
	"net/http"

	"github.com/dasagho/htmx-test/models"
	"github.com/dasagho/htmx-test/views"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := models.Body{
		MouseDiv: "Mamala",
		Search:   "Search",
		List:     models.Result{List: []string{}},
	}

	views.InitializeTemplate()
	err := views.GetTemplates().ExecuteTemplate(w, "index", data)
	if err != nil {
		fmt.Fprintf(w, "error al ejecutar el template index %s", err)
	}

}
