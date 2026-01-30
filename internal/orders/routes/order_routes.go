package order_routes

import (
	"backend/internal/middleware"
	order_handler "backend/internal/orders/handler"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, h *order_handler.OrderHandler) {

	order := r.Group("/orders")
	order.Use(middleware.JWTAuth())

	order.POST("", h.PlaceOrder)
	order.GET("", h.GetMyOrders)
	order.GET("/:order_id", h.GetByOrderId)

}
