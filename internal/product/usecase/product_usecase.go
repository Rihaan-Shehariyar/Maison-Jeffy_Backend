package usecases

import (
	"backend/internal/product/entity"
	"backend/internal/product/repositorys"
	"errors"
)

type ProductUseCase struct {
	repo repositorys.ProductRepository
}

func NewProductRepositoryUseCase (repo repositorys.ProductRepository)*ProductUseCase{
  return &ProductUseCase{repo:repo}
}

func(u *ProductUseCase) GetAllProducts(
 category string,maxPrice *float64,sort string,search string,
)([]entitys.Product, error){
  
 if sort != ""{
  validSorts := map[string]bool{
	 "price_asc" : true,
     "price_desc" : true,
     "latest" : true,
}

  if !validSorts[sort]{
  return nil,errors.New("Invalid Sort Option")
}
 
 
}

return u.repo.FindAll(category,maxPrice,sort,search)
  
}

func (u *ProductUseCase)GetProductByID(id uint)(*entitys.Product,error){
   return u.repo.FindByID(id)
}



