package cart_repository

import cart_entity "backend/internal/cart/entity"

type CartRepository interface {
	Add(userID, productID uint) error
	Remove(userID, productID uint) error
	UpdateQty(userID, productId uint, qty uint) error
	Clear(userID uint) error
	Exists(userID, productID uint) (bool, error)
	GetByUser(userID uint) ([]cart_entity.Cart, error)
}
