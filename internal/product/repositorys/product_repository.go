package repositorys

import "backend/internal/product/entity"

type ProductRepository interface {
	FindAll(category string,maxPrice *float64,sort string,search string) ([]entitys.Product,error)
    FindByID(id uint) (*entitys.Product,error)
    Create(product *entitys.Product) error
    Update(product *entitys.Product)error
    Delete(id uint)error
}

