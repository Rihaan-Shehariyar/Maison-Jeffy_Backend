package order_entity

import "time"

type Order struct {

	ID          uint `gorm:"primaryKey"`
	UserID      uint `gorm:"index"`
	TotalAmount float64
	Status      string
	CreatedAt   time.Time
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID"`

}
