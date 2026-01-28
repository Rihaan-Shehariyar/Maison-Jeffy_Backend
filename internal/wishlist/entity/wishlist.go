package wishlist_entity

import "time"

type Wishlist struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `gorm:"not null"`
	ProductID  uint `gorm:"uniqueIndex:idx_user_product"`
	Created_at time.Time
}


