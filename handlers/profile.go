package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"GO_Auth/db"
	"GO_Auth/models"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	var user models.User
	err := db.DB.QueryRow(
		context.Background(),
		`SELECT id, email, first_name, last_name,
		        date_of_birth, phone,
		        address_line1, address_line2,
		        city, state, country, postal_code,
		        is_email_verified, created_at
		 FROM users WHERE id=$1`,
		userID,
	).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.Phone,
		&user.AddressLine1,
		&user.AddressLine2,
		&user.City,
		&user.State,
		&user.Country,
		&user.PostalCode,
		&user.IsVerified,
		&user.CreatedAt,
	)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
