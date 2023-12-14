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
		List: models.Result{List: []string{}},
		ContactList: []models.Contact{
			models.NewContact(1, 123456789, "Manolo de los palotes"),
			models.NewContact(2, 987654321, "El señor de la noche"),
			models.NewContact(3, 147258369, "Don Omar"),
		},
	}

	err := views.GetTemplates().ExecuteTemplate(w, "index", data)
	if err != nil {
		fmt.Fprintf(w, "error al ejecutar el template index %s", err)
		return
	}
	logging.Info("Index correct attended")
}
