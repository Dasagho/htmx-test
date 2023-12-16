package handlers

import (
	"fmt"
	"net/http"

	logging "github.com/dasagho/htmx-test/log"
	"github.com/dasagho/htmx-test/views"
)

func SketchHandler(w http.ResponseWriter, r *http.Request) {
	err := views.Render(w, "sketch", nil)
	if err != nil {
		logging.Error(fmt.Sprintf("Fallo al intentar renderizar el template sketch: %s", err.Error()))
		http.Error(w, "Fail render page", http.StatusInternalServerError)
		return
	}
}
