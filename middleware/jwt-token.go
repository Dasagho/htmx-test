package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dasagho/htmx-test/models"
	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const (
	userContextKey contextKey = "userData"
)

var (
	jwtToken string
)

func init() {
	jwtToken = os.Getenv("JWT_TOKEN")
	if jwtToken == "" {
		log.Fatal("Missing JWT Token")
	}
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the cookie
		cookie, err := r.Cookie("auth_token") // Assuming the cookie is named "auth_token"
		if err != nil {
			http.Error(w, "Unauthorized - No token", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value

		// Verify and parse the token
		claims := &models.Claims{} // Replace 'models.Claims' with your actual Claims struct
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtToken), nil // Replace with your secret key
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized - Invalid token", http.StatusUnauthorized)
			return
		}

		// Pass the user data to the handler
		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Function to generate a JWT token
func GenerateJWTToken(userID string) (string, error) {
	// Set up JWT claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token (replace "yourSigningKey" with your actual key)
	tokenString, err := token.SignedString([]byte(jwtToken))

	return tokenString, err
}
