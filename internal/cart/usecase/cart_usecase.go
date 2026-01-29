package cart_usecase

import (
	cart_entity "backend/internal/cart/entity"
	cart_repository "backend/internal/cart/repository"
	"errors"
)

type CartUsecase struct {
	repo cart_repository.CartRepository
}

func NewCartUsecase(repo cart_repository.CartRepository) *CartUsecase {
	return &CartUsecase{repo}
}

func (u *CartUsecase) Add(userID, productID uint) error {

	if productID == 0 {
		return errors.New("Invalid Product ID")
	}

	return u.repo.Add(userID, productID)

}

func (u *CartUsecase) Remove(userID, productID uint) error {
	if productID == 0 {
		return errors.New("Invalid Product ID")
	}

	return u.repo.Remove(userID, productID)

}

func (u *CartUsecase) UpdateQty(userID, productID uint, qty int) error {

	if qty == 0 {
		return u.repo.Remove(userID, productID)
	}

	return u.repo.UpdateQty(userID, productID, uint(qty))

}

func (u *CartUsecase) Clear(userID uint) error {

	return u.repo.Clear(userID)
}

func (u *CartUsecase) GetMyUser(userID uint) ([]cart_entity.Cart, error) {

	return u.repo.GetByUser(userID)

}

func (u *CartUsecase) ValidateCart(userID uint) error {

	items, err := u.repo.GetByUser(userID)
	if err != nil {
		return err
	}

	if len(items) == 0 {
		return errors.New("Cart is Empty")
	}

	return nil

}
