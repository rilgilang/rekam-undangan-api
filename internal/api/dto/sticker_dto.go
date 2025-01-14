package dto

import (
	"time"
)

type OneSticker struct {
	ID        string      `json:"id"`
	Url       *string     `json:"url"`
	Tag       interface{} `json:"tag"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
