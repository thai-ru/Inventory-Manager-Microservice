package db

import (
	"Inventory_Microservice/config"
	"Inventory_Microservice/internal/stock_level_service/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewStockDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dialector := sqlite.Open(cfg.Source)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to stock database: %w", err)
	}

	if err := db.AutoMigrate(&model.StockLevel{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate stock: %w", err)
	}

	return db, nil
}
