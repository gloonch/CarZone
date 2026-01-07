package login

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gloonch/CarZone/models"
	"github.com/golang-jwt/jwt/v4"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)

		return
	}

	valid := (credentials.UserName == "admin" && credentials.Password == "admin123")
	if !valid {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
	}

	tokenString, err := GenerateToken(credentials.UserName)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		log.Println("Error generating token")

		return
	}

	response := map[string]string{"token": tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.Header().Set("Authorization", "Bearer "+tokenString)
}

func GenerateToken(username string) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expiration.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("seycreyt"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
