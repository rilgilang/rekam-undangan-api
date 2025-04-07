package entities

import (
	"gorm.io/gorm"
	"time"
)

type Room struct {
	gorm.Model
	ID                   string     `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	IDCard               string     `gorm:"type:varchar(255);unique" json:"-"`
	RoomImage            *string    `gorm:"type:varchar(255);unique" json:"room_image"`
	RoomNumber           int        `gorm:"type:integer;not null" json:"room_number"`
	Renter               string     `gorm:"type:varchar(255)" json:"renter"`
	Price                int        `gorm:"type:integer" json:"price"`
	AlreadyPaidThisMonth bool       `gorm:"type:boolean" json:"already_paid_this_month"`
	Available            bool       `gorm:"type:boolean" json:"available"`
	FirstCheckIn         time.Time  `gorm:"not null" json:"first_check_in"`
	CheckIn              time.Time  `gorm:"not null" json:"check_in"`
	CheckOut             time.Time  `gorm:"not null" json:"check_out"`
	CreatedAt            time.Time  `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time  `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt            *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

type DetailedRoom struct {
	ID                   string           `json:"id"`
	IDCard               string           `json:"-"`
	RoomImage            *string          `json:"room_image"`
	RoomNumber           int              `json:"room_number"`
	Renter               string           `json:"renter"`
	Price                int              `json:"price"`
	AlreadyPaidThisMonth bool             `json:"already_paid_this_month"`
	Available            bool             `json:"available"`
	FirstCheckIn         time.Time        `json:"first_check_in"`
	CheckIn              time.Time        `json:"check_in"`
	CheckOut             time.Time        `json:"check_out"`
	PaymentHistory       []PaymentHistory `json:"payment_history"`
	CreatedAt            time.Time        `json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time        `json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt            *time.Time       `json:"deleted_at,omitempty"`
}

type UpdateRenterPayload struct {
	ID     string `json:"id,omitempty"`
	IDCard string `json:"id_card,omitempty"`
	Renter string `json:"renter,omitempty"`
}
