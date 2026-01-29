package cart_entity

import "time"

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null;index"`
	ProductID uint `gorm:"not null;index"`
	Quantity  uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
