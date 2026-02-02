package order_repository

import (
	order_entity "backend/internal/orders/entity"
)

type OrderRepository interface {
	Create(order *order_entity.Order) error
	GetByUser(userID uint) ([]order_entity.Order, error)
	GetByOrderId(orderID uint) (*order_entity.Order, error)
	UpdateStatus(orderId uint, status string) error
}

 
