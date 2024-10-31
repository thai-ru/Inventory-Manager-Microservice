package model

import (
	"gorm.io/gorm"
	"time"
)

type StockLevel struct {
	gorm.Model
	ProductID    string
	Quantity     int64
	MinThreshold int64
	MaxThreshold int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
