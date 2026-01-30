package order_usecase

import (
	cart_repository "backend/internal/cart/repository"
	order_entity "backend/internal/orders/entity"
	order_repository "backend/internal/orders/repository"
	"backend/internal/product/repositorys"
	"errors"

	"gorm.io/gorm"
)

type OrderUseCase struct {
	db          *gorm.DB
	orderRepo   order_repository.OrderRepository
	cartRepo    cart_repository.CartRepository
	productRepo repositorys.ProductRepository
}

func NewOrderUsecase(
	db *gorm.DB,
	orderRepo order_repository.OrderRepository,
	cartRepo cart_repository.CartRepository,
	productRepo repositorys.ProductRepository,
) *OrderUseCase {
	return &OrderUseCase{
		db:          db,
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (u *OrderUseCase) PlaceOrder(userID uint) error {

	return u.db.Transaction(func(tx *gorm.DB) error {
		cartItems, err := u.cartRepo.GetByUser(userID)
		if err != nil {
			return err
		}

		if len(cartItems) == 0 {
			return errors.New("Cart Is empty")
		}

		var total float64
		var orderItems []order_entity.OrderItem

		for _, item := range cartItems {

			product, err := u.productRepo.FindByID(item.ProductID)
			if err != nil {
				return errors.New("Invalid Product ID")
			}

			total += product.Price * float64(item.Quantity)

			orderItems = append(orderItems, order_entity.OrderItem{
				ProductID: product.ID,
				Price:     product.Price,
				Quantity:  item.Quantity,
			})

		}

		order := order_entity.Order{

			UserID:      userID,
			TotalAmount: total,
			Status:      "placed",
			OrderItems:  orderItems,
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		if err := u.cartRepo.Clear(userID); err != nil {
			return err
		}

		return nil

	})

}

func (u *OrderUseCase) GetMyOrders(userId uint) ([]order_entity.Order, error) {

	return u.orderRepo.GetByUser(userId)

}

func (u *OrderUseCase) GetByOrderId(orderId uint) (*order_entity.Order, error) {

	return u.orderRepo.GetByOrderId(orderId)
}

func (u *OrderUseCase) UpdateStatus(orderId uint, status string) error {

	validStatus := map[string]bool{

		"placed":    true,
		"paid":      true,
		"shipped":   true,
		"delivered": true,
		"cancelled": true,
	}

	if !validStatus[status] {
		return errors.New("Invalid order status")
	}

	return u.orderRepo.UpdateStatus(orderId, status)

}
