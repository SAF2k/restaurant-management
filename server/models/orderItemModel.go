package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID        primitive.ObjectID `bson:"_id"`
	Order_id  *string            `json:"order_id" binding:"required" validate:"required,min=2,max=100"`
	Menu_id   *string            `json:"menu_id" validate:"required,min=2,max=100"`
	Food_id   *string            `json:"food_id" validate:"required,min=2,max=100"`
	Quantity  *int               `json:"quantity" validate:"required,min=1,max=100"`
	Price     *float64           `json:"price" validate:"required,min=0,max=1000000"`
	Create_at time.Time          `json:"create_at"`
	Update_at time.Time          `json:"update_at"`
}
