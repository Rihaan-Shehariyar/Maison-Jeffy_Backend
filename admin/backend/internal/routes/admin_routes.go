package routes

import (
	"backend/admin/backend/internal/handlers"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {

	admin := r.Group("/admin")

	admin.POST("/login", handlers.AdminLogin)

	admin.Use(middleware.JWTAuth(), middleware.AdminOnly())

	admin.POST("/products", handlers.CreateProduct)
	admin.GET("/products", handlers.GetAllProducts)
	admin.DELETE("/products/:id", handlers.DeleteProduct)
	admin.PUT("/products/:id", handlers.UpdateProduct)

	admin.GET("/users", handlers.GetAllUsers)
	admin.PUT("/users/:id/block", handlers.BlockUser)

	admin.GET("/orders", handlers.GetAllOrders)
	admin.PUT("/orders/:id/status", handlers.UpdateOrderStatus)

	admin.GET("/dashboard", handlers.GetDashboardStats)

}
