package entities

import (
	"gorm.io/gorm"
	"time"
)

type PaymentHistory struct {
	gorm.Model
	ID             string    `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	Total          int       `gorm:"type:integer" json:"total"`
	PaymentReceipt string    `gorm:"type:varchar(255)" json:"payment_receipt"`
	ForMonth       string    `gorm:"type:varchar(15)" json:"for_month"`
	RoomId         string    `gorm:"type:varchar(36);not null" json:"room_id"`
	CreatedAt      time.Time `gorm:"not null" json:"created_at," sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"not null" json:"updated_at," sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

type CreatePaymentHistoryPayload struct {
	RoomID         string `json:"room_id"`
	Total          int    `json:"total"`
	PaymentReceipt string `json:"payment_receipt"`
}
