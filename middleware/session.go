package middleware

import (
	"crypto/rand"
	"fmt"
	"net/http"

	logging "github.com/dasagho/htmx-test/log"
)

var (
	sessionStore = map[string]map[string]interface{}{}
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.Debug("Session Middlware")
		_, ok := getSession(r)
		if !ok {
			sessionData := map[string]interface{}{"": ""}
			setSession(w, sessionData)
			http.Redirect(w, r, "/", http.StatusFound)
			logging.Debug("Created and redirect")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func generateSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func setSession(w http.ResponseWriter, sessionData map[string]interface{}) {
	sessionID := generateSessionID()
	sessionStore[sessionID] = sessionData
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: sessionID,
		Path:  "/",
	})
}

func getSession(r *http.Request) (map[string]interface{}, bool) {
	c, err := r.Cookie("session_id")
	if err != nil {
		return nil, false
	}
	sessionData, ok := sessionStore[c.Value]
	return sessionData, ok
}
