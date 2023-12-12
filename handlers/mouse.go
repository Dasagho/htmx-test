package handlers

import (
	"fmt"
	"net/http"
)

func Mouse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Por el culo")
}
