package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Id          string
	ProductName string
	Category    string
	Description string
	BasePrice   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
