package models

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name" binding:"required" validate:"required,min=2,max=100"`
	Price     float64   `gorm:"not null" json:"price" binding:"required" validate:"required,min=0,max=1000000"`
	FoodImage *string   `json:"food_image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FoodID    string    `gorm:"index" json:"food_id"`
	MenuName  *string   `json:"menu_name"`
	MenuID    *uint     `gorm:"index" json:"menu_id"`
	StoreID   *uint     `gorm:"index" json:"store_id"`
}

// Define foreign key constraints
func (f *Food) BeforeSave(tx *gorm.DB) error {
	if f.MenuID != nil {
		tx = tx.Model(&Menu{}).Where("id = ?", *f.MenuID).First(&Menu{})
		if tx.Error != nil {
			return tx.Error
		}
	}
	if f.StoreID != nil {
		tx = tx.Model(&Store{}).Where("id = ?", *f.StoreID).First(&Store{})
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}
