package order_repository

import (
	order_entity "backend/internal/orders/entity"

	"gorm.io/gorm"
)

type OrderRepositoryPg struct {
	db *gorm.DB
}

func NewOrderRepositoryPg(db *gorm.DB) OrderRepository {
	return &OrderRepositoryPg{db: db}
}

func (r *OrderRepositoryPg) Create(order *order_entity.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepositoryPg) GetByUser(userID uint) ([]order_entity.Order, error) {

	var orders []order_entity.Order

	err := r.db.Preload("OrderItems").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error

	return orders, err

}

func (r *OrderRepositoryPg) GetByOrderId(orderID uint) (*order_entity.Order, error) {

	var order order_entity.Order

	err := r.db.Preload("OrderItems").
		First(&order, orderID).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepositoryPg) UpdateStatus(orderId uint, status string) error {

	return r.db.Model(&order_entity.Order{}).
		Where("order_id = ?", orderId).
		Update("status", status).Error
}
