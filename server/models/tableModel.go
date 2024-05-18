package models

import (
	"time"

	"github.com/google/uuid"
)

type Table struct {
	ID             uuid.UUID `gorm:"primaryKey" json:"id"`
	TableID        string    `gorm:"unique;not null" json:"table_id"`
	StoreID        string    `gorm:"not null" json:"store_id"`
	NumberOfGuests int       `gorm:"not null" json:"number_of_guests" validate:"required"`
	TableNumber    int       `gorm:"not null" json:"table_number" validate:"required"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}
