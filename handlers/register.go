package handlers

import (
	"NGC_AVENGER/config"
	"NGC_AVENGER/models"

	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Validasi input
	err := validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Simpan ke database
	_, err = config.DB.Exec("INSERT INTO users (email, password, full_name, age, occupation, role) VALUES (?, ?, ?, ?, ?, ?)",
		user.Email, user.Password, user.FullName, user.Age, user.Occupation, user.Role)
	if err != nil {
		http.Error(w, "Error saving to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
