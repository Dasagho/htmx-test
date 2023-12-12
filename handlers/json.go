package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dasagho/htmx-test/models"
)

func JsonResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	list := models.Result{
		List: []string{"Hey", "hey2"},
	}
	jsonData, err := json.Marshal(list)
	if err != nil {
		http.Error(w, "Error al generar el JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
