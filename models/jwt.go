package models

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	jwt.StandardClaims // This embeds common standard claims like `exp`
	// Add any custom claims here, if needed
}
