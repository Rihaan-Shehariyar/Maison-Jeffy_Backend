package cart_repository

import (
	cart_entity "backend/internal/cart/entity"
	"errors"

	"gorm.io/gorm"
)

type CartRepositoryPg struct {
	db *gorm.DB
}

func NewCartRepositoryPg(db *gorm.DB) CartRepository {
	return &CartRepositoryPg{db}
}

func (r *CartRepositoryPg) Add(userID, productID uint) error {

	var cart cart_entity.Cart

	err := r.db.Where("user_id = ? AND  product_id = ?", userID, productID).First(&cart).Error

	if err == nil {
		return r.db.Model(&cart).Update("quantity", cart.Quantity+1).Error
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.db.Create(&cart_entity.Cart{
			UserID:    userID,
			ProductID: productID,
			Quantity:  1,
		}).Error
	}

	return err

}

func (r *CartRepositoryPg) Exists(userID, productID uint) (bool, error) {
	var count int64
	err := r.db.Model(&cart_entity.Cart{}).
		Where("user_id = ? AND product_id = ?", userID, productID).
		Count(&count).Error
	return count > 0, err
}

func (r *CartRepositoryPg) GetByUser(userID uint) ([]cart_entity.Cart, error) {

	var items []cart_entity.Cart

	err := r.db.Where("user_id = ?", userID).Find(&items).Error
	return items, err

}

func (r *CartRepositoryPg) Remove(userID, productID uint) error {

	return r.db.Where("user_id = ? AND product_id = ? ", userID, productID).Delete(&cart_entity.Cart{}).Error

}

func (r *CartRepositoryPg) Clear(userID uint) error {

	return r.db.Where("user_id = ?", userID).Delete(&cart_entity.Cart{}).Error

}

func (r *CartRepositoryPg) UpdateQty(userID, productId uint, qty uint) error {

	return r.db.Model(&cart_entity.Cart{}).
		Where("user_id = ? AND product_id = ? ", userID, productId).
		Update("quantity", qty).Error
}
