package models

import (
	"time"
)

// LoginDTO represents the payload for login
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// SignupDTO represents the payload for signup
type SignupDTO struct {
	LoginDTO
	UserID    *string   `json:"user_id"`
	Name      string    `json:"name" validate:"required,min=3"`
	StoreID   string    `json:"store_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User represents user data stored in the database
type User struct {
	UserID    int       `json:"user_id"`
	Name      string    `json:"name" validate:"required,min=3"`
	Email     string    `json:"email" validate:"required,email"`
	StoreID   string    `json:"store_id" validate:"required"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// AccessResponse represents the response for successful authentication
type AccessResponse struct {
	Token string `json:"token"`
}

// AuthResponse represents the response for authentication
type AuthResponse struct {
	User *User           `json:"user"`
	Auth *AccessResponse `json:"auth"`
}
