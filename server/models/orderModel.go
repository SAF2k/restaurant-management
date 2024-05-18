package models

import (
	"time"
)

type Order struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	Order_id     string    `gorm:"column:order_id" json:"order_id"`
	Table_id     string    `gorm:"column:table_id" json:"table_id"`
	Order_status *string   `gorm:"column:order_status" json:"order_status" validate:"eq=OPEN|eq=CLOSE|eq="`
	Order_date   time.Time `gorm:"column:order_date" json:"order_date"`
	Created_at   time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
