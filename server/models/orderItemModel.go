package models

import (
	"time"
)

type OrderItem struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	Quantity      *int      `gorm:"column:quantity" json:"quantity" validate:"required"`
	Unit_price    *float64  `gorm:"column:unit_price" json:"unit_price" validate:"required"`
	Food_id       string    `gorm:"column:food_id" json:"food_id"`
	Order_item_id string    `gorm:"column:order_item_id" json:"order_item_id"`
	Order_id      string    `gorm:"column:order_id" json:"order_id"`
	Created_at    time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at    time.Time `gorm:"column:updated_at" json:"updated_at"`
}
