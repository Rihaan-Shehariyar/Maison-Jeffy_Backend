package order_entity

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	UserID    uint
	ProductID uint
	Price     float64
	Quantity  uint
}
