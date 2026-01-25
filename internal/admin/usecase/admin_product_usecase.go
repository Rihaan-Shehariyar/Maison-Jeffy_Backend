package admin_usecase

import (
	entitys "backend/internal/product/entity"
	"backend/internal/product/repositorys"
	"errors"
)

type ProductAdminUsecase struct {
	repo repositorys.ProductRepository
}

func NewProductAdminUsecase(repo repositorys.ProductRepository) *ProductAdminUsecase {
	return &ProductAdminUsecase{repo}
}

func (u *ProductAdminUsecase) CreateProduct(product *entitys.Product) error {
	return u.repo.Create(product)
}

func (u *ProductAdminUsecase) UpdateProduct(id uint, name, description string, stock int, price float64) error {
	product, err := u.repo.FindByID(id)
	if err != nil {
		return errors.New("Product Not Found")
	}

	if name != "" {
		product.Name = name
	}
	if description != "" {
		product.Description = description
	}
	if price > 0 {
		product.Price = price
	}

	if stock >= 0 {
		product.Stock = stock

	}

	return u.repo.Update(product)
}

func (u *ProductAdminUsecase) DeleteProduct(id uint) error {

	_, err := u.repo.FindByID(id)

	if err != nil {
		return errors.New("Product Not Found")
	}

	return u.repo.Delete(id)

}
