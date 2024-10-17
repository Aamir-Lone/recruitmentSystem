package utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your_secret_key")

// JWTAuthMiddleware validates the token and extracts user info
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract email from token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["email"] == nil {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Add email to the request context
		email := claims["email"].(string)
		ctx := context.WithValue(r.Context(), "userID", email) // or "email" if that's what you want

		// Pass the request to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
