package admin_routes

import (
	admin_handler "backend/internal/admin/handler"
	"backend/internal/middleware"
	order_handler "backend/internal/orders/handler"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine, h *admin_handler.UserAdminHandler, p *admin_handler.ProductAdminHandler, o *order_handler.OrderHandler) {

	admin := r.Group("/admin")
	admin.Use(middleware.JWTAuth(), middleware.AdminOnly())

	users := admin.Group("/users")
	users.PUT("/:id", h.UpdateUser)
	users.PUT("/:id/block", h.BlockUser)

	products := admin.Group("/products")
	products.POST("", p.CreateProduct)
	products.PUT("/:id", p.UpdateProduct)
	products.DELETE("/:id", p.DeleteProduct)

	order := admin.Group("/orders")
	order.PUT("/:order_id", o.UpdateStatus)
}
