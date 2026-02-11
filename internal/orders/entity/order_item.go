package order_entity

import entitys "backend/internal/product/entity"

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `gorm:"not null;index" json:"order_id"`
	ProductID uint    `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  uint    `json:"quantity"`

	Product entitys.Product `gorm:"foreignKey:ProductID" json:"product"`
}
