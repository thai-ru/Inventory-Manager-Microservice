package db

import (
	"Inventory_Microservice/config"
	"Inventory_Microservice/internal/product_service/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewDB initializes a new database connection based on the config settings
func NewDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	var dialector gorm.Dialector

	// Determine which driver to use based on cfg.Driver
	switch cfg.Driver {
	case "postgres":
		dialector = postgres.Open(cfg.Source)
	case "sqlite":
		dialector = sqlite.Open(cfg.Source)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}

	// Open the connection with GORM
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Perform auto-migrations
	if err := db.AutoMigrate(&model.Product{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	return db, nil
}
