package wishlist_repository

import (
	wishlist_entity "backend/internal/wishlist/entity"

	"gorm.io/gorm"
)

type wishlistRepositoryPg struct {
	db *gorm.DB
}

func NewWishlistRepositoryPg(db *gorm.DB) WishlistRepository{
  return &wishlistRepositoryPg{db}
}

func (r *wishlistRepositoryPg) Add(userID,productID uint)error{
  return r.db.Create(&wishlist_entity.Wishlist{
   UserID: userID,
   ProductID: productID,
}).Error
}

func (r *wishlistRepositoryPg) Remove(userID,productID uint) error{
  return r.db.Where("user_id = ? AND product_id = ?",userID,productID).
   Delete(&wishlist_entity.Wishlist{}).Error
}

func(r *wishlistRepositoryPg) GetMyWishlist(userID uint)([]wishlist_entity.Wishlist,error){
  var items []wishlist_entity.Wishlist

 err:= r.db.Where("user_id = ?",userID).Find(&items).Error
   return items,err 

}

func (r *wishlistRepositoryPg)Exists(userID,productID uint)(bool,error){

  var count int64

 err:= r.db.Model(&wishlist_entity.Wishlist{}).
        Where("user_id = ? AND product_id = ? ",userID,productID).
        Count(&count).Error
 
 return count>0,err

}
