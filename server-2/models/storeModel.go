package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID       primitive.ObjectID `bson:"_id"`
	Store_id string             `json:"store_id" validate:"required"`
	Name     *string            `json:"name" validate:"required"`

	Created_at time.Time `json:"create_at"`
	Updated_at time.Time `json:"update_at"`
}
