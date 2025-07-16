package entities

import (
	"time"
)

type User struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	VerifiedAt time.Time `json:"verified_at"`
	UpdatedAt  time.Time `gorm:"not null" json:"updated_at," sql:"DEFAULT:CURRENT_TIMESTAMP"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at," sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
