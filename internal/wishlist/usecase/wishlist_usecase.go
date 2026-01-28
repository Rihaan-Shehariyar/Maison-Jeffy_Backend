package wishlist_usecase

import (
	wishlist_entity "backend/internal/wishlist/entity"
	wishlist_repository "backend/internal/wishlist/repository"
	"errors"
)

type WishlistUsecase struct {
	repo wishlist_repository.WishlistRepository
}

func NewWishlistUsecasePg(repo wishlist_repository.WishlistRepository) *WishlistUsecase {
	return &WishlistUsecase{repo}
}

func (u *WishlistUsecase) Add(userID uint, productID uint) error {

	if productID == 0 {
		return errors.New("invalid product id")
	}

	exists, err := u.repo.Exists(userID, productID)

	if err != nil {
		return err
	}
	if exists {
		return errors.New("Product already in wishlist")
	}

	return u.repo.Add(userID, productID)

}

func (u *WishlistUsecase) Remove(userId, productID uint) error {
	return u.repo.Remove(userId, productID)
}

func (u *WishlistUsecase) GetMyWishlist(userID uint) ([]wishlist_entity.Wishlist, error) {

	return u.repo.GetMyWishlist(userID)

}
