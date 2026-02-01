package order_entity

type OrderItem struct {

	ID        uint `gorm:"primaryKey"`
	OrderID   uint `gorm:"not null;index"`
	ProductID uint
	Price     float64
	Quantity  uint

}
