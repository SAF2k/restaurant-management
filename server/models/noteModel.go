package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `json:"name" binding:"required" validate:"required,min=2,max=100"`
	Note_id   string             `json:"note_id"`
	Create_at time.Time          `json:"create_at"`
	Update_at time.Time          `json:"update_at"`
}
