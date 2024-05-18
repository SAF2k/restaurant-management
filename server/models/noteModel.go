package models

import (
	"time"
)

type Note struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Text      *string   `gorm:"column:text" json:"text" validate:"required,min=2,max=20"`
	Title     *string   `gorm:"column:title" json:"title"`
	Note_id   string    `gorm:"column:note_id" json:"note_id"`
	Create_at time.Time `gorm:"column:created_at" json:"created_at"`
	Update_at time.Time `gorm:"column:updated_at" json:"updated_at"`
}
