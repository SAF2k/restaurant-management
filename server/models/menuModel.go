package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" binding:"required" validate:"required,min=2,max=100"`
	Category   string             `json:"category" binding:"required" validate:"required,min=2,max=100"`
	Start_Date *time.Time         `json:"start_date"`
	End_Date   *time.Time         `json:"end_date"`
	Menu_id    string             `json:"menu_id"`
	Created_at time.Time          `json:"create_at"`
	Updated_at time.Time          `json:"update_at"`
}
