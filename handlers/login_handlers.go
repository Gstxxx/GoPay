package handlers

import (
	"GoPay/models"
	"GoPay/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := authenticate(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{Token: token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func authenticate(email, password string) (*models.User, error) {
	for _, user := range models.Users {
		if user.Email == email && user.Password == password {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User not found or invalid password")
}
