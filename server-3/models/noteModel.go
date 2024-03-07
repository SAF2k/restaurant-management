package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id"`
	Text      *string            `json:"text" validate:"required,min=2,max=20"`
	Title     *string            `json:"title"`
	Note_id   string             `json:"note_id"`
	Create_at time.Time          `json:"created_at"`
	Update_at time.Time          `json:"updated_at"`
}
