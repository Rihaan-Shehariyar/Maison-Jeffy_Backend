package wishlist_repository

import wishlist_entity "backend/internal/wishlist/entity"

type WishlistRepository interface {
	Add(userID, productID uint) error
	Remove(userID, productID uint) error
	GetMyWishlist(userID uint) ([]wishlist_entity.Wishlist,error)
    Exists(userID,productID uint)(bool,error)
}

