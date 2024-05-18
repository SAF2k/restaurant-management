package models

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Name     *string   `gorm:"column:name" json:"name" binding:"required" validate:"required,min=2,max=100"`
	Category *string   `gorm:"column:category" json:"category" binding:"required" validate:"required,min=2,max=100"`
	Menu_id  string    `gorm:"column:menu_id" json:"menu_id"`
	Store_id *string   `gorm:"column:store_id" json:"store_id"`

	Created_at time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type MenuResponse struct {
	Name     string `gorm:"column:name" json:"name"`
	Category string `gorm:"column:category" json:"category"`
}
