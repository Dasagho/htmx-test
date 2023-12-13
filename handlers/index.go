package handlers

import (
	"fmt"
	"net/http"

	logging "github.com/dasagho/htmx-test/log"
	"github.com/dasagho/htmx-test/models"
	"github.com/dasagho/htmx-test/views"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	logging.Info("Index Requested")
	data := models.Body{
		MouseDiv: "Mamala",
		Search:   "Search",
		List:     models.Result{List: []string{}},
	}

	err := views.GetTemplates().ExecuteTemplate(w, "index", data)
	if err != nil {
		fmt.Fprintf(w, "error al ejecutar el template index %s", err)
		return
	}
	logging.Info("Index correct attended")
}
