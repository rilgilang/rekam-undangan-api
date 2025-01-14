package entities

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Sticker struct {
	gorm.Model
	ID        string           `gorm:"type:varchar(36);primary_key;unique" json:"id"`
	Url       *string          `gorm:"type:text" json:"url"` // Pointer to allow NULL
	Tag       *json.RawMessage `gorm:"type:json" json:"tag"` // Pointer to allow NULL
	CreatedAt time.Time        `gorm:"not null" json:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time        `gorm:"not null" json:"updatedAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time       `sql:"index" json:"deletedAt,omitempty"`
}
