package order_entity

import (
	"backend/internal/auth/entity"
	"time"
)

type Order struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	UserID      uint        `gorm:"index" json:"user_id"`
    User        entity.User `gorm:"foreignKey:UserID" json:"user"`
	TotalAmount float64     `json:"total_amount"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}
