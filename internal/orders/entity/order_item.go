package order_entity

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
	Price     float64
	Quantity  uint
}
