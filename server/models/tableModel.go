package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID        primitive.ObjectID `bson:"_id"`
	Table_id  *string            `json:"table_id" binding:"required" validate:"required,min=2,max=100"`
	Table_no  *string            `json:"table_no" binding:"required" validate:"required,min=2,max=100"`
	Table_max *int               `json:"table_max" validate:"required,min=1,max=100"`
	Create_at time.Time          `json:"create_at"`
	Update_at time.Time          `json:"update_at"`
}
