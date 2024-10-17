package controllers

import (
	"RecruitmentManagementSystem/models"
	"RecruitmentManagementSystem/services"
	"RecruitmentManagementSystem/utils"
	"encoding/json"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser, err := services.CreateUser(user)
	if err != nil {
		http.Error(w, "Unable to create user", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	json.NewDecoder(r.Body).Decode(&credentials)

	user, err := services.AuthenticateUser(credentials.Email, credentials.PasswordHash)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
