package handlers

import (
	"fmt"
	"net/http"
	"time"

	logging "github.com/dasagho/htmx-test/log"
	"github.com/dasagho/htmx-test/middleware"
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

	userID := "exampleUserID" // After successful authentication

	// Generate JWT token
	tokenString, err := middleware.GenerateJWTToken(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new HTTP cookie
	cookie := http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Expires:  time.Now().Add(72 * time.Hour),
		HttpOnly: true,                 // Important: Make the cookie inaccessible to JavaScript
		Path:     "/",                  // The cookie is available for all paths
		Secure:   true,                 // Set to true if you are using HTTPS
		SameSite: http.SameSiteLaxMode, // CSRF protection
	}

	// Set the cookie in the response header
	http.SetCookie(w, &cookie)

	err = views.GetTemplates().ExecuteTemplate(w, "index", data)
	if err != nil {
		fmt.Fprintf(w, "error al ejecutar el template index %s", err)
		return
	}
	logging.Info("Index correct attended")
}
