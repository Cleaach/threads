package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}




var JwtSecret = []byte("ENTER_SECRET_HERE")

func ExtractUserIDFromJWT(r *http.Request) (int, error) {
	// Get the "Authorization" header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("missing Authorization header")
	}

	// Split the header to extract the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, fmt.Errorf("invalid Authorization header format")
	}

	tokenStr := parts[1]

	// Parse the token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return JwtSecret, nil
	})
	
	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	// Extract user ID from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["userID"].(float64)
	fmt.Println(userID)
	if !ok {
		return 0, fmt.Errorf("userID not found in token claims")
	}

	return int(userID), nil
}
