package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" binding:"required" validate:"required,min=2,max=100"`
	Price      *float64           `json:"price" binding:"required" validate:"required,min=0,max=1000000"`
	Food_Image *string            `json:"food_image" binding:"required" validate:"required,min=2,max=100"`
	Created_at time.Time          `json:"create_at"`
	Updated_at time.Time          `json:"update_at"`
	Food_id    string             `json:"food_id"`
	Menu_id    *string            `json:"menu_id"`
}
