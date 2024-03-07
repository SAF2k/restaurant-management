package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     *string            `json:"name" binding:"required" validate:"required,min=2,max=100"`
	Category *string            `json:"category" binding:"required" validate:"required,min=2,max=100"`
	Menu_id  string             `json:"menu_id"`
	Store_id *string            `json:"store_id"`

	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type MenuResponse struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}
