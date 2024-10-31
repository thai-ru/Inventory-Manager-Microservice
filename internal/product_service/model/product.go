package model

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	ProductName string
	Category    string
	Description string
	BasePrice   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
