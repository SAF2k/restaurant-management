package models

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	StoreID   string    `gorm:"unique;not null" json:"store_id"`
	Name      string    `gorm:"not null" json:"name" validate:"required"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
