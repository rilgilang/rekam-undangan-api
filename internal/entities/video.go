package entities

import (
	"gorm.io/gorm"
	"time"
)

type ProcessVideoPayload struct {
	UniqueId string `json:"unique_id"`
	URL      string `json:"url"`
}

type Video struct {
	gorm.Model
	ID          string    `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	UniqueId    string    `gorm:"type:varchar(255);unique" json:"unique_id"`
	OriginalUrl string    `gorm:"type:text" json:"original_url"`
	URL         string    `gorm:"type:text" json:"url"`
	UserID      string    `gorm:"type:varchar(36)" json:"user_id"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at," sql:"DEFAULT:CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at," sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
