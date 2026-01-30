package order_handler

import (
	order_usecase "backend/internal/orders/usecase"
	"backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase order_usecase.OrderUseCase
}

func NewOrderHandler(usecase order_usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{usecase}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {

	userID := c.GetUint("user_id")

	if err := h.usecase.PlaceOrder(userID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Order Placed Successfully"})
}

func (h *OrderHandler) GetMyOrders(c *gin.Context) {

	userId := c.GetUint("user_id")

	orders, err := h.usecase.GetMyOrders(userId)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, orders)

}

func (h *OrderHandler) GetByOrderId(c *gin.Context) {

	idParam := c.Param("order_id")
	OrderID, err := strconv.Atoi(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid Order ID")
		return
	}

	orders, err := h.usecase.GetByOrderId(uint(OrderID))
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, orders)

}

func (h *OrderHandler) UpdateStatus(c *gin.Context) {

	idParam := c.Param("order_id")
	orderId, err := strconv.Atoi(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid Order ID")
		return
	}

	var req struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid Json")
		return
	}

	if err := h.usecase.UpdateStatus(uint(orderId), req.Status); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Order Status Updated"})

}
