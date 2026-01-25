package repositorys

import (
	entitys "backend/internal/product/entity"

	"gorm.io/gorm"
)

type productRepositoryPg struct {
	db *gorm.DB
}

func NewProductRepositoryPg(db *gorm.DB) ProductRepository {
	return &productRepositoryPg{db: db}
}

func (r *productRepositoryPg) FindAll() ([]entitys.Product, error) {
	var products []entitys.Product
	return products, r.db.Find(&products).Error

}

func (r *productRepositoryPg) FindByID(id uint) (*entitys.Product, error) {
	var product entitys.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepositoryPg) Create(product *entitys.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepositoryPg) Update(prduct *entitys.Product) error {
	return r.db.Save(prduct).Error
}
func (r *productRepositoryPg) Delete(id uint) error {
	return r.db.Delete(&entitys.Product{}, id).Error
}
