package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"log"

	"GO_Auth/db"
	"GO_Auth/models"
	"GO_Auth/services"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email        string     `json:"email"`
		Password     string     `json:"password"`
		FirstName    string     `json:"first_name"`
		LastName     string     `json:"last_name"`
		DateOfBirth  *time.Time `json:"date_of_birth"`
		Phone        string     `json:"phone"`
		AddressLine1 string     `json:"address_line_1"`
		AddressLine2 string     `json:"address_line_2"`
		City         string     `json:"city"`
		State        string     `json:"state"`
		Country      string     `json:"country"`
		PostalCode   string     `json:"postal_code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password required", http.StatusBadRequest)
		return
	}

	if len(req.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters", http.StatusBadRequest)
		return
	}

	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to secure password", 500)
		return
	}

	var userID string
	err = db.DB.QueryRow(
		context.Background(),
		`INSERT INTO users (
			email, password_hash,
			first_name, last_name,
			date_of_birth, phone,
			address_line1, address_line2,
			city, state, country, postal_code,
			is_email_verified, is_active
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,false,true
		)
		RETURNING id`,
		req.Email,
		hashedPassword,
		req.FirstName,
		req.LastName,
		req.DateOfBirth,
		req.Phone,
		req.AddressLine1,
		req.AddressLine2,
		req.City,
		req.State,
		req.Country,
		req.PostalCode,
	).Scan(&userID)

	if err != nil {
	log.Println("Signup error:", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}


	accessToken, _ := services.GenerateAccessToken(userID)
	refreshToken, _ := services.GenerateRefreshToken(userID)

	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}


func Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	var user models.User
	err := db.DB.QueryRow(
		context.Background(),
		`SELECT id, password_hash
		 FROM users
		 WHERE email=$1 AND is_active=true`,
		req.Email,
	).Scan(&user.ID, &user.PasswordHash)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !services.CheckPassword(req.Password, user.PasswordHash) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, _ := services.GenerateAccessToken(user.ID)
	refreshToken, _ := services.GenerateRefreshToken(user.ID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
