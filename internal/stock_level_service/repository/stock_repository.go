package repository

import (
	"Inventory_Microservice/internal/stock_level_service/model"
	"gorm.io/gorm"
)

type StockRepository interface {
	AddStock(stock *model.StockLevel) error
}

type DatabaseStockStore struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *DatabaseStockStore {
	return &DatabaseStockStore{db: db}
}

func (db *DatabaseStockStore) AddStock(stock *model.StockLevel) error {
	return db.db.Create(stock).Error
}
