package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LoginDTO defined the /login payload
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"password"`
}

// Deadline implements context.Context.
func (*LoginDTO) Deadline() (deadline time.Time, ok bool) {
	panic("unimplemented")
}

// Done implements context.Context.
func (*LoginDTO) Done() <-chan struct{} {
	panic("unimplemented")
}

// Err implements context.Context.
func (*LoginDTO) Err() error {
	panic("unimplemented")
}

// Value implements context.Context.
func (*LoginDTO) Value(key any) any {
	panic("unimplemented")
}

// SignupDTO defined the /login payload
type SignupDTO struct {
	LoginDTO
	User_id    *string   `json:"user_id"`
	Name       string    `json:"name" validate:"required,min=3"`
	Store_id   string    `json:"store_id" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// UserResponse todo
type UserResponse struct {
	ID         primitive.ObjectID `bson:"-"`
	User_id    string             `json:"-"`
	Name       string             `json:"name"`
	Email      string             `json:"email"`
	Store_id   string             `json:"store_id"`
	Password   string             `json:"-"`
	Created_at time.Time          `json:"-"`
	Updated_at time.Time          `json:"-"`
}

// AccessResponse todo
type AccessResponse struct {
	Token string `json:"token"`
}

// AuthResponse todo
type AuthResponse struct {
	User *UserResponse   `json:"user"`
	Auth *AccessResponse `json:"auth"`
}
