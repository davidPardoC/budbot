package models

import "time"

type Budget struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Amount    float64   `json:"amount" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int64     `json:"created_by"`
}
