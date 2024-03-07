package models

import (
	"time"
)

type Invoice struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	InvoiceID      string    `gorm:"column:invoice_id" json:"invoice_id"`
	OrderID        *string   `gorm:"column:order_id" json:"order_id" binding:"required" validate:"required,min=2,max=100"`
	PaymentMethod  *string   `gorm:"column:payment_method" json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	PaymentStatus  *string   `gorm:"column:payment_status" json:"payment_status" validate:"eq=PAID|eq=UNPAID|eq="`
	PaymentDueDate time.Time `gorm:"column:payment_due_date" json:"payment_due_date"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}
