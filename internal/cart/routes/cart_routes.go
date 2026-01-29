package cart_routes

import (
	cart_handler "backend/internal/cart/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine, h *cart_handler.CartHandler) {

	cart := r.Group("/cart")
	cart.Use(middleware.JWTAuth())

	cart.POST("/:productID", h.Add)
	cart.DELETE("/:productID", h.Remove)
	cart.DELETE("", h.Clear)
	cart.PUT("/:productID", h.UpdateQty)
	cart.GET("", h.GetMyCart)

}
