package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`

	PasswordHash string     `json:"-"` 

	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	Phone	 string   `json:"phone,omitempty"`

	AddressLine1 string `json:"address_line_1,omitempty"`
	AddressLine2 string `json:"address_line_2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Country      string `json:"country,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`

	IsVerified bool      `json:"is_email_verified"`

	CreatedAt  time.Time `json:"created_at"`
}