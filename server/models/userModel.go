package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  *string            `json:"username" binding:"required" validate:"required,min=2,max=100"`
	Password  *string            `json:"password" binding:"required" validate:"required,min=2,max=100"`
	Role      *string            `json:"role" validate:"eq=ADMIN|eq=USER|eq="`
	Create_at time.Time          `json:"create_at"`
	Update_at time.Time          `json:"update_at"`
}
