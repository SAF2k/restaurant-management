package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID       primitive.ObjectID `bson:"_id"`
	Store_id string             `json:"store_id"`
	Name     *string            `json:"name" validate:"required"`

	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
