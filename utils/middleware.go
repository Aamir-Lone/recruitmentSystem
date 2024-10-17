package utils

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Claims struct to hold JWT claims
type Claims struct {
	Email   string `json:"email"`
	UserID  string `json:"userID"`
	IsAdmin bool   `json:"isAdmin"` // Added a field to check if the user is an admin
	jwt.StandardClaims
}

// AuthMiddleware validates the JWT and extracts user information.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Split the token from "Bearer <token>"
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the JWT and get claims
		claims, valid := ValidateJWT(token)
		if !valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Create a new context with userID and email
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "email", claims.Email)

		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// AdminAuthMiddleware ensures the user is an admin by checking the claims.
func AdminAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Split the token from "Bearer <token>"
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the JWT and get claims
		claims, valid := ValidateJWT(token)
		if !valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the user is an admin
		if !claims.IsAdmin {
			http.Error(w, "You do not have permission to access this resource", http.StatusForbidden)
			return
		}

		// Call the next handler if the user is an admin
		next.ServeHTTP(w, r)
	}
}

// ValidateJWT is a placeholder function. You should implement your JWT validation logic here.
// func ValidateJWT(token string) (Claims, bool) {
// 	// Implement JWT validation logic and extract claims
// 	// This is a placeholder for demonstration
// 	// Replace with actual JWT validation
// 	return Claims{Email: "example@example.com", UserID: "12345", IsAdmin: true}, true
// }
