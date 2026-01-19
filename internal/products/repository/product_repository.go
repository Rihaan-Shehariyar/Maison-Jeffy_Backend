package repository

import "backend/internal/products/entity"

type ProductRepository interface {
	FindAll() ([]entity.Product,error)
    FindById(id uint) (*entity.Product,error)
    Create(product *entity.Product) error
}

