package repository

import (
	"Inventory_Microservice/internal/product_service/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product *model.Product) error
	FindByID(id string) (*model.Product, error)
}

type DatabaseProductStore struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *DatabaseProductStore {
	return &DatabaseProductStore{
		DB: db,
	}
}

func (p *DatabaseProductStore) Save(product *model.Product) error {
	return p.DB.Create(product).Error
}
