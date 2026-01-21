package usecases

import (
	"backend/internal/product/entity"
	"backend/internal/product/repositorys"
)

type ProductUseCase struct {
	repo repositorys.ProductRepository
}

func NewProductRepositoryUseCase (repo repositorys.ProductRepository)*ProductUseCase{
  return &ProductUseCase{repo:repo}
}

func(u *ProductUseCase) GetAllProducts()([]entitys.Product, error){
   return u.repo.FindAll()
}

func (u *ProductUseCase)GetProductByID(id uint)(*entitys.Product,error){
   return u.repo.FindByID(id)
}



