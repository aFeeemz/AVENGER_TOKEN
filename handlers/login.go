package handlers

import (
	"NGC_AVENGER/config"
	"NGC_AVENGER/models"
	"database/sql"
	"time"

	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	// "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var input models.User //Containing the login email

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User //Containing the database email
	err := config.DB.QueryRow("SELECT id, email, password, full_name, age, occupation, role FROM users WHERE email = ?", input.Email).
		Scan(&user.ID, &user.Email, &user.Password, &user.FullName, &user.Age, &user.Occupation, &user.Role)

	if err == sql.ErrNoRows || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	expirationTimes := time.Now().Add(2 * time.Minute)

	claims := &models.Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTimes),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
