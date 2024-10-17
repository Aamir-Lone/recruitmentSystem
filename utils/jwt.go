package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_very_secret_key_that_is_long_and_random")

// type Claims struct {
// 	Email string `json:"email"`
// 	jwt.StandardClaims
// }

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, bool) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, false
	}
	return claims, true
}
